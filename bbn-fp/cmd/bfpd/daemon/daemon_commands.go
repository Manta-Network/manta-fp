package daemon

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	fpcmd "github.com/Manta-Network/manta-fp/bbn-fp/cmd"
	fpcfg "github.com/Manta-Network/manta-fp/bbn-fp/config"
	"github.com/Manta-Network/manta-fp/bbn-fp/proto"
	dc "github.com/Manta-Network/manta-fp/bbn-fp/service/client"
	client2 "github.com/Manta-Network/manta-fp/eotsmanager/client"

	"cosmossdk.io/math"
	"github.com/babylonlabs-io/babylon/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkflags "github.com/cosmos/cosmos-sdk/client/flags"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	defaultFpdDaemonAddress = "127.0.0.1:" + strconv.Itoa(fpcfg.DefaultRPCPort)
	defaultEotsAddress      = "127.0.0.1:" + strconv.Itoa(12582)
	defaultAppHashStr       = "fd903d9baeb3ab1c734ee003de75f676c5a9a8d0574647e5385834d57d3e79ec"
)

// CommandGetDaemonInfo returns the get-info command by connecting to the bfpd daemon.
func CommandGetDaemonInfo() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "get-info",
		Aliases: []string{"gi"},
		Short:   "Get information of the running bfpd daemon.",
		Example: fmt.Sprintf(`bfpd get-info --daemon-address %s`, defaultFpdDaemonAddress),
		Args:    cobra.NoArgs,
		RunE:    runCommandGetDaemonInfo,
	}
	cmd.Flags().String(fpdDaemonAddressFlag, defaultFpdDaemonAddress, "The RPC server address of bfpd")
	return cmd
}

func runCommandGetDaemonInfo(cmd *cobra.Command, _ []string) error {
	daemonAddress, err := cmd.Flags().GetString(fpdDaemonAddressFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", fpdDaemonAddressFlag, err)
	}

	client, cleanUp, err := dc.NewFinalityProviderServiceGRpcClient(daemonAddress)
	if err != nil {
		return err
	}
	defer func() {
		if err := cleanUp(); err != nil {
			fmt.Printf("Failed to clean up grpc client: %v\n", err)
		}
	}()

	info, err := client.GetInfo(context.Background())
	if err != nil {
		return err
	}

	printRespJSON(info)
	return nil
}

// CommandCreateFP returns the create-bbn-fp command by connecting to the bfpd daemon.
func CommandCreateFP() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "create-bbn-fp",
		Aliases: []string{"cfp"},
		Short:   "Create a finality provider object and save it in database.",
		Long: "Create a new finality provider object and store it in the finality provider database. " +
			"It needs to have an operating EOTS manager available and running.",
		Example: strings.TrimSpace(
			fmt.Sprintf(`
Either by specifying all flags manually:

$bfpd create-bbn-fp --daemon-address %s ...

Or providing the path to bbn-fp.json:
$bfpd create-bbn-fp --daemon-address %s --from-file /path/to/bbn-fp.json

Where bbn-fp.json contains:

{
  "keyName": "The unique key name of the finality provider's Babylon account",
  "chainID": "The identifier of the consumer chain",
  "passphrase": "The pass phrase used to encrypt the keys",
  "commissionRate": "The commission rate for the finality provider, e.g., 0.05"",
  "moniker": ""A human-readable name for the finality provider",
  "identity": "A optional identity signature",
  "website": "Validator's (optional) website",
  "securityContract": "Validator's (optional) security contact email",
  "details": "Validator's (optional) details",
  "eotsPK": "The hex string of the finality provider's EOTS public key"
}
`, defaultFpdDaemonAddress, defaultFpdDaemonAddress)),
		Args: cobra.NoArgs,
		RunE: fpcmd.RunEWithClientCtx(runCommandCreateFP),
	}

	f := cmd.Flags()
	f.String(fpdDaemonAddressFlag, defaultFpdDaemonAddress, "The RPC server address of bfpd")
	f.String(keyNameFlag, "", "The unique key name of the finality provider's Babylon account")
	f.String(sdkflags.FlagHome, fpcfg.DefaultFpdDir, "The application home directory")
	f.String(chainIDFlag, "", "The identifier of the consumer chain")
	f.String(passphraseFlag, "", "The pass phrase used to encrypt the keys")
	f.String(commissionRateFlag, "", "The commission rate for the finality provider, e.g., 0.05")
	f.String(monikerFlag, "", "A human-readable name for the finality provider")
	f.String(identityFlag, "", "An optional identity signature (ex. UPort or Keybase)")
	f.String(websiteFlag, "", "An optional website link")
	f.String(securityContactFlag, "", "An email for security contact")
	f.String(detailsFlag, "", "Other optional details")
	f.String(fpEotsPkFlag, "", "The hex string of the finality provider's EOTS public key")
	f.String(fromFile, "", "Path to a json file containing finality provider data")

	cmd.PreRunE = func(cmd *cobra.Command, _ []string) error {
		fromFilePath, _ := cmd.Flags().GetString(fromFile)
		if fromFilePath == "" {
			// Mark flags as required only if --from-file is not provided
			if err := cmd.MarkFlagRequired(chainIDFlag); err != nil {
				return err
			}
			if err := cmd.MarkFlagRequired(keyNameFlag); err != nil {
				return err
			}
			if err := cmd.MarkFlagRequired(monikerFlag); err != nil {
				return err
			}
			if err := cmd.MarkFlagRequired(commissionRateFlag); err != nil {
				return err
			}
			if err := cmd.MarkFlagRequired(fpEotsPkFlag); err != nil {
				return err
			}
		}
		return nil
	}

	return cmd
}

func runCommandCreateFP(ctx client.Context, cmd *cobra.Command, _ []string) error {
	flags := cmd.Flags()

	fpJSONPath, err := flags.GetString(fromFile)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", fromFile, err)
	}

	var fp *parsedFinalityProvider
	if fpJSONPath != "" {
		fp, err = parseFinalityProviderJSON(fpJSONPath, ctx.HomeDir)
		if err != nil {
			panic(err)
		}
	} else {
		fp, err = parseFinalityProviderFlags(cmd, ctx.HomeDir)
		if err != nil {
			panic(err)
		}
	}

	daemonAddress, err := flags.GetString(fpdDaemonAddressFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", fpdDaemonAddressFlag, err)
	}

	client, cleanUp, err := dc.NewFinalityProviderServiceGRpcClient(daemonAddress)
	if err != nil {
		return err
	}
	defer func() {
		if err := cleanUp(); err != nil {
			fmt.Printf("Failed to clean up grpc client: %v\n", err)
		}
	}()

	res, err := client.CreateFinalityProvider(
		context.Background(),
		fp.keyName,
		fp.chainID,
		fp.eotsPK,
		fp.passphrase,
		fp.description,
		&fp.commissionRate,
	)
	if err != nil {
		return err
	}

	printRespJSON(res)
	return nil
}

// CommandUnjailFP returns the unjail-bbn-fp command by connecting to the bfpd daemon.
func CommandUnjailFP() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "unjail-bbn-fp",
		Aliases: []string{"ufp"},
		Short:   "Unjail the given finality provider.",
		Example: fmt.Sprintf(`bfpd unjail-bbn-fp [eots-pk] --daemon-address %s ...`, defaultFpdDaemonAddress),
		Args:    cobra.ExactArgs(1),
		RunE:    fpcmd.RunEWithClientCtx(runCommandUnjailFP),
	}

	f := cmd.Flags()
	f.String(fpdDaemonAddressFlag, defaultFpdDaemonAddress, "The RPC server address of bfpd")

	return cmd
}

func runCommandUnjailFP(_ client.Context, cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
	daemonAddress, err := flags.GetString(fpdDaemonAddressFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", fpdDaemonAddressFlag, err)
	}

	client, cleanUp, err := dc.NewFinalityProviderServiceGRpcClient(daemonAddress)
	if err != nil {
		return err
	}
	defer func() {
		if err := cleanUp(); err != nil {
			fmt.Printf("Failed to clean up grpc client: %v\n", err)
		}
	}()

	_, err = client.UnjailFinalityProvider(context.Background(), args[0])
	if err != nil {
		return err
	}

	return nil
}

func getDescriptionFromFlags(f *pflag.FlagSet) (stakingtypes.Description, error) {
	// get information for description
	var desc stakingtypes.Description
	monikerStr, err := f.GetString(monikerFlag)
	if err != nil {
		return desc, fmt.Errorf("failed to read flag %s: %w", monikerFlag, err)
	}
	identityStr, err := f.GetString(identityFlag)
	if err != nil {
		return desc, fmt.Errorf("failed to read flag %s: %w", identityFlag, err)
	}
	websiteStr, err := f.GetString(websiteFlag)
	if err != nil {
		return desc, fmt.Errorf("failed to read flag %s: %w", websiteFlag, err)
	}
	securityContactStr, err := f.GetString(securityContactFlag)
	if err != nil {
		return desc, fmt.Errorf("failed to read flag %s: %w", securityContactFlag, err)
	}
	detailsStr, err := f.GetString(detailsFlag)
	if err != nil {
		return desc, fmt.Errorf("failed to read flag %s: %w", detailsFlag, err)
	}

	description := stakingtypes.NewDescription(monikerStr, identityStr, websiteStr, securityContactStr, detailsStr)
	return description.EnsureLength()
}

// CommandLsFP returns the list-finality-providers command by connecting to the bfpd daemon.
func CommandLsFP() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list-finality-providers",
		Aliases: []string{"ls"},
		Short:   "List finality providers stored in the database.",
		Example: fmt.Sprintf(`bfpd list-finality-providers --daemon-address %s`, defaultFpdDaemonAddress),
		Args:    cobra.NoArgs,
		RunE:    runCommandLsFP,
	}
	cmd.Flags().String(fpdDaemonAddressFlag, defaultFpdDaemonAddress, "The RPC server address of bfpd")
	return cmd
}

func runCommandLsFP(cmd *cobra.Command, _ []string) error {
	daemonAddress, err := cmd.Flags().GetString(fpdDaemonAddressFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", fpdDaemonAddressFlag, err)
	}

	client, cleanUp, err := dc.NewFinalityProviderServiceGRpcClient(daemonAddress)
	if err != nil {
		return err
	}
	defer func() {
		if err := cleanUp(); err != nil {
			fmt.Printf("Failed to clean up grpc client: %v\n", err)
		}
	}()

	resp, err := client.QueryFinalityProviderList(context.Background())
	if err != nil {
		return err
	}
	printRespJSON(resp)

	return nil
}

// CommandInfoFP returns the bbn-fp-info command by connecting to the bfpd daemon.
func CommandInfoFP() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "bbn-fp-info [fp-eots-pk-hex]",
		Aliases: []string{"fpi"},
		Short:   "List finality providers stored in the database.",
		Example: fmt.Sprintf(`bfpd bbn-fp-info --daemon-address %s`, defaultFpdDaemonAddress),
		Args:    cobra.ExactArgs(1),
		RunE:    runCommandInfoFP,
	}
	cmd.Flags().String(fpdDaemonAddressFlag, defaultFpdDaemonAddress, "The RPC server address of bfpd")
	return cmd
}

func runCommandInfoFP(cmd *cobra.Command, args []string) error {
	fpPk, err := types.NewBIP340PubKeyFromHex(args[0])
	if err != nil {
		return err
	}

	daemonAddress, err := cmd.Flags().GetString(fpdDaemonAddressFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", fpdDaemonAddressFlag, err)
	}

	client, cleanUp, err := dc.NewFinalityProviderServiceGRpcClient(daemonAddress)
	if err != nil {
		return err
	}
	defer func() {
		if err := cleanUp(); err != nil {
			fmt.Printf("Failed to clean up grpc client: %v\n", err)
		}
	}()

	resp, err := client.QueryFinalityProviderInfo(context.Background(), fpPk)
	if err != nil {
		return err
	}
	printRespJSON(resp)

	return nil
}

// CommandAddFinalitySig returns the add-finality-sig command by connecting to the bfpd daemon.
func CommandAddFinalitySig() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "unsafe-add-finality-sig [fp-eots-pk-hex] [block-height]",
		Aliases: []string{"unsafe-afs"},
		Short:   "[UNSAFE] Send a finality signature to the consumer chain.",
		Long:    "[UNSAFE] Send a finality signature to the consumer chain. This command should only be used for presentation/testing purposes",
		Example: fmt.Sprintf(`bfpd unsafe-add-finality-sig [fp-eots-pk-hex] [block-height] --daemon-address %s`, defaultFpdDaemonAddress),
		Args:    cobra.ExactArgs(2),
		RunE:    runCommandAddFinalitySig,
	}
	cmd.Flags().String(fpdDaemonAddressFlag, defaultFpdDaemonAddress, "The RPC server address of bfpd")
	cmd.Flags().String(appHashFlag, defaultAppHashStr, "The last commit hash of the chain block")
	cmd.Flags().Bool(checkDoubleSignFlag, true, "If 'true', uses anti-slashing protection when doing EOTS sign")

	return cmd
}

func runCommandAddFinalitySig(cmd *cobra.Command, args []string) error {
	fpPk, err := types.NewBIP340PubKeyFromHex(args[0])
	if err != nil {
		return err
	}
	blkHeight, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		return err
	}

	flags := cmd.Flags()
	daemonAddress, err := flags.GetString(fpdDaemonAddressFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", fpdDaemonAddressFlag, err)
	}

	appHashHex, err := flags.GetString(appHashFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", appHashFlag, err)
	}

	checkDoubleSign, err := flags.GetBool(checkDoubleSignFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", checkDoubleSignFlag, err)
	}

	client, cleanUp, err := dc.NewFinalityProviderServiceGRpcClient(daemonAddress)
	if err != nil {
		return err
	}
	defer func() {
		if err := cleanUp(); err != nil {
			fmt.Printf("Failed to clean up grpc client: %v\n", err)
		}
	}()

	appHash, err := hex.DecodeString(appHashHex)
	if err != nil {
		return err
	}

	res, err := client.AddFinalitySignature(context.Background(), fpPk.MarshalHex(), blkHeight, appHash, checkDoubleSign)
	if err != nil {
		return err
	}
	printRespJSON(res)

	return nil
}

// CommandAddFinalitySig returns the add-finality-sig command by connecting to the bfpd daemon.
func CommandAddEotsKey() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "add-eots-key [key-name] [pass-phrase] [hd-path]",
		Aliases: []string{"add-ek"},
		Short:   "[eots] add a key to the eots.",
		Long:    "[add] add a key to the eots. This command should only be used for presentation/testing purposes",
		Example: fmt.Sprintf(`bfpd add-eots-key [key-name] [pass-phrase] [hd-path] --eots-address %s`, defaultFpdDaemonAddress),
		Args:    cobra.ExactArgs(2),
		RunE:    runCommandAddEotsKey,
	}
	cmd.Flags().String(eotsAddressFlag, defaultEotsAddress, "The RPC server address of eots")
	//cmd.Flags().String(keyNameFlag, "", "The key name of eots key")
	//cmd.Flags().String(passphraseFlag, "", "The pass phrase used to encrypt the keys")
	cmd.Flags().String(hdPathFlag, "", "The hd path to store eots key")

	return cmd
}

func runCommandAddEotsKey(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
	eotsAddress, err := flags.GetString(eotsAddressFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", eotsAddressFlag, err)
	}
	//keyName, err := flags.GetString(keyNameFlag)
	//if err != nil {
	//	return fmt.Errorf("failed to read flag %s: %w", keyNameFlag, err)
	//}
	//passphrase, err := flags.GetString(passphraseFlag)
	//if err != nil {
	//	return fmt.Errorf("failed to read flag %s: %w", passphraseFlag, err)
	//}
	hdPath, err := flags.GetString(hdPathFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", hdPathFlag, err)
	}

	em, err := client2.NewEOTSManagerGRpcClient(eotsAddress)
	if err != nil {
		return fmt.Errorf("failed to create EOTS manager client: %w", err)
	}

	var fpPk *types.BIP340PubKey
	fpPkBytes, err := em.CreateKey(args[0], args[1], hdPath)
	if err != nil {
		return err
	}
	fpPk, err = types.NewBIP340PubKey(fpPkBytes)
	if err != nil {
		return err
	}
	printRespJSON(fpPk)
	return nil
}

// CommandEditFinalityDescription edits description of finality provider
func CommandEditFinalityDescription() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "edit-bbn-fp [eots_pk]",
		Aliases: []string{"efp"},
		Short:   "Edit finality provider data without resetting unchanged fields",
		Long: "Edit the details of a finality provider using the specified EOTS public key. " +
			"\nThe provided [eots_pk] must correspond to the Babylon address controlled by the key specified in bfpd.conf. " +
			"\nIf one or more optional flags are passed (such as --moniker, --website, etc.), " +
			"the corresponding values are updated, while unchanged fields retain their current values from the Babylon Node.",
		Example: fmt.Sprintf(`bfpd edit-bbn-fp [eots_pk] --daemon-address %s --moniker "new-moniker"`, defaultFpdDaemonAddress),
		Args:    cobra.ExactArgs(1),
		RunE:    runCommandEditFinalityDescription,
	}
	cmd.Flags().String(fpdDaemonAddressFlag, defaultFpdDaemonAddress, "The RPC server address of bfpd")
	cmd.Flags().String(monikerFlag, "", "The finality provider's (optional) moniker")
	cmd.Flags().String(websiteFlag, "", "The finality provider's (optional) website")
	cmd.Flags().String(securityContactFlag, "", "The finality provider's (optional) security contact email")
	cmd.Flags().String(detailsFlag, "", "The finality provider's (optional) details")
	cmd.Flags().String(identityFlag, "", "The (optional) identity signature (ex. UPort or Keybase)")
	cmd.Flags().String(commissionRateFlag, "", "The (optional) commission rate percentage (ex. 0.2)")

	return cmd
}

func runCommandEditFinalityDescription(cmd *cobra.Command, args []string) error {
	fpPk, err := types.NewBIP340PubKeyFromHex(args[0])
	if err != nil {
		return err
	}

	flags := cmd.Flags()
	daemonAddress, err := flags.GetString(fpdDaemonAddressFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", fpdDaemonAddressFlag, err)
	}

	grpcClient, cleanUp, err := dc.NewFinalityProviderServiceGRpcClient(daemonAddress)
	if err != nil {
		return err
	}
	defer func() {
		if err := cleanUp(); err != nil {
			fmt.Printf("Failed to clean up grpc client: %v\n", err)
		}
	}()

	moniker, _ := cmd.Flags().GetString(monikerFlag)
	website, _ := cmd.Flags().GetString(websiteFlag)
	securityContact, _ := cmd.Flags().GetString(securityContactFlag)
	details, _ := cmd.Flags().GetString(detailsFlag)
	identity, _ := cmd.Flags().GetString(identityFlag)
	rate, _ := cmd.Flags().GetString(commissionRateFlag)

	desc := &proto.Description{
		Moniker:         moniker,
		Identity:        identity,
		Website:         website,
		SecurityContact: securityContact,
		Details:         details,
	}

	if err := grpcClient.EditFinalityProvider(cmd.Context(), fpPk, desc, rate); err != nil {
		return fmt.Errorf("failed to edit finality provider %v err %w", fpPk.MarshalHex(), err)
	}

	return nil
}

func printRespJSON(resp interface{}) {
	jsonBytes, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Println("unable to decode response: ", err)
		return
	}

	fmt.Printf("%s\n", jsonBytes)
}

func loadKeyName(homeDir string, cmd *cobra.Command) (string, error) {
	keyName, err := cmd.Flags().GetString(keyNameFlag)
	if err != nil {
		return "", fmt.Errorf("failed to read flag %s: %w", keyNameFlag, err)
	}
	// if key name is not specified, we use the key of the config
	if keyName != "" {
		return keyName, nil
	}

	// we add the following check to ensure that the chain key is created
	// beforehand
	cfg, err := fpcfg.LoadConfig(homeDir)
	if err != nil {
		return "", fmt.Errorf("failed to load config from %s: %w", fpcfg.CfgFile(homeDir), err)
	}

	keyName = cfg.BabylonConfig.Key
	if keyName == "" {
		return "", fmt.Errorf("the key in config is empty")
	}
	return keyName, nil
}

type parsedFinalityProvider struct {
	keyName        string
	chainID        string
	eotsPK         string
	passphrase     string
	description    stakingtypes.Description
	commissionRate math.LegacyDec
}

func parseFinalityProviderJSON(path string, homeDir string) (*parsedFinalityProvider, error) {
	type internalFpJSON struct {
		KeyName          string `json:"keyName"`
		ChainID          string `json:"chainID"`
		Passphrase       string `json:"passphrase"`
		CommissionRate   string `json:"commissionRate"`
		Moniker          string `json:"moniker"`
		Identity         string `json:"identity"`
		Website          string `json:"website"`
		SecurityContract string `json:"securityContract"`
		Details          string `json:"details"`
		EotsPK           string `json:"eotsPK"`
	}

	// #nosec G304 - The log file path is provided by the user and not externally
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var fp internalFpJSON
	if err := json.Unmarshal(contents, &fp); err != nil {
		return nil, err
	}

	if fp.ChainID == "" {
		return nil, fmt.Errorf("chainID is required")
	}

	if fp.KeyName == "" {
		cfg, err := fpcfg.LoadConfig(homeDir)
		if err != nil {
			return nil, fmt.Errorf("failed to load config from %s: %w", fpcfg.CfgFile(homeDir), err)
		}
		fp.KeyName = cfg.BabylonConfig.Key
		if fp.KeyName == "" {
			return nil, fmt.Errorf("the key is neither in config nor provided in the json file")
		}
	}

	if fp.Moniker == "" {
		return nil, fmt.Errorf("moniker is required")
	}

	if fp.CommissionRate == "" {
		return nil, fmt.Errorf("commissionRate is required")
	}

	if fp.EotsPK == "" {
		return nil, fmt.Errorf("eotsPK is required")
	}

	commissionRate, err := math.LegacyNewDecFromStr(fp.CommissionRate)
	if err != nil {
		return nil, fmt.Errorf("invalid commission rate: %w", err)
	}

	description, err := stakingtypes.NewDescription(fp.Moniker, fp.Identity, fp.Website, fp.SecurityContract, fp.Details).EnsureLength()
	if err != nil {
		return nil, err
	}

	return &parsedFinalityProvider{
		keyName:        fp.KeyName,
		chainID:        fp.ChainID,
		eotsPK:         fp.EotsPK,
		passphrase:     fp.Passphrase,
		description:    description,
		commissionRate: commissionRate,
	}, nil
}

func parseFinalityProviderFlags(cmd *cobra.Command, homeDir string) (*parsedFinalityProvider, error) {
	flags := cmd.Flags()

	commissionRateStr, err := flags.GetString(commissionRateFlag)
	if err != nil {
		return nil, fmt.Errorf("failed to read flag %s: %w", commissionRateFlag, err)
	}
	commissionRate, err := math.LegacyNewDecFromStr(commissionRateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid commission rate: %w", err)
	}

	description, err := getDescriptionFromFlags(flags)
	if err != nil {
		return nil, fmt.Errorf("invalid description: %w", err)
	}

	keyName, err := loadKeyName(homeDir, cmd)
	if err != nil {
		return nil, fmt.Errorf("not able to load key name: %w", err)
	}

	if keyName == "" {
		return nil, fmt.Errorf("keyname cannot be empty")
	}

	chainID, err := flags.GetString(chainIDFlag)
	if err != nil {
		return nil, fmt.Errorf("failed to read flag %s: %w", chainIDFlag, err)
	}

	if chainID == "" {
		return nil, fmt.Errorf("chain-id cannot be empty")
	}

	passphrase, err := flags.GetString(passphraseFlag)
	if err != nil {
		return nil, fmt.Errorf("failed to read flag %s: %w", passphraseFlag, err)
	}

	eotsPkHex, err := flags.GetString(fpEotsPkFlag)
	if err != nil {
		return nil, fmt.Errorf("failed to read flag %s: %w", fpEotsPkFlag, err)
	}

	if eotsPkHex == "" {
		return nil, fmt.Errorf("eots-pk cannot be empty")
	}

	return &parsedFinalityProvider{
		keyName:        keyName,
		chainID:        chainID,
		eotsPK:         eotsPkHex,
		passphrase:     passphrase,
		description:    description,
		commissionRate: commissionRate,
	}, nil
}
