// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DefaultVaultInitParams is an auto generated low-level Go binding around an user-defined struct.
type DefaultVaultInitParams struct {
	Version              uint64
	DelegatorIndex       uint64
	SlasherIndex         uint64
	VaultSlashBurner     common.Address
	VaultDefaultAdmin    common.Address
	VaultBeforeSlashHook common.Address
}

// OperatorSettings is an auto generated low-level Go binding around an user-defined struct.
type OperatorSettings struct {
	MinOperatorCommission *big.Int
	MaxOperatorCommission *big.Int
}

// SymbioticVaultSettings is an auto generated low-level Go binding around an user-defined struct.
type SymbioticVaultSettings struct {
	OperatorRegistry   common.Address
	VaultConfiguration common.Address
	EpochDuration      *big.Int
	StakeToken         common.Address
	DefaultVaultParams DefaultVaultInitParams
}

// MantaStakingMiddlewareMetaData contains all meta data concerning the MantaStakingMiddleware contract.
var MantaStakingMiddlewareMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_symbioticVaultSettings\",\"type\":\"tuple\",\"internalType\":\"structSymbioticVaultSettings\",\"components\":[{\"name\":\"operatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultConfiguration\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"stakeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultVaultParams\",\"type\":\"tuple\",\"internalType\":\"structDefaultVaultInitParams\",\"components\":[{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delegatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slasherIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"vaultSlashBurner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultDefaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeforeSlashHook\",\"type\":\"address\",\"internalType\":\"address\"}]}]},{\"name\":\"_operatorSettings\",\"type\":\"tuple\",\"internalType\":\"structOperatorSettings\",\"components\":[{\"name\":\"minOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"modifySupportedTokens\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_supported\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorNameExists\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSettings\",\"inputs\":[],\"outputs\":[{\"name\":\"minOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operators\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"operatorName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"commission\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseOperator\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"_operatorPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_operatorName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_commission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardAddressExists\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbioticVaultSettings\",\"inputs\":[],\"outputs\":[{\"name\":\"operatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultConfiguration\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"stakeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultVaultParams\",\"type\":\"tuple\",\"internalType\":\"structDefaultVaultInitParams\",\"components\":[{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delegatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slasherIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"vaultSlashBurner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultDefaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeforeSlashHook\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpauseOperator\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorSettings\",\"inputs\":[{\"name\":\"_operatorSettings\",\"type\":\"tuple\",\"internalType\":\"structOperatorSettings\",\"components\":[{\"name\":\"minOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSymbioticVaultSettings\",\"inputs\":[{\"name\":\"_symbioticVaultSettings\",\"type\":\"tuple\",\"internalType\":\"structSymbioticVaultSettings\",\"components\":[{\"name\":\"operatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultConfiguration\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"stakeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultVaultParams\",\"type\":\"tuple\",\"internalType\":\"structDefaultVaultInitParams\",\"components\":[{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delegatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slasherIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"vaultSlashBurner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultDefaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeforeSlashHook\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorPaused\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorPublicKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"operatorName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"commission\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnpaused\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardAddressSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SupportedTokensModified\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"supported\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b612076806100d96000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806368c4ac26116100ad57806385eb242d1161007157806385eb242d146103e457806391d14854146103f7578063a217fddf1461040a578063d547741f14610412578063ecfb6e8d1461042557600080fd5b806368c4ac261461033c5780637174d6691461035f57806372aef6bf1461038257806372f9adab146103be5780637a6efb69146103d157600080fd5b80631d6a65d7116100f45780631d6a65d7146102bf578063248a9ca3146102e25780632e5aaf33146103035780632f2ff15d1461031657806336568abe1461032957600080fd5b806301ffc9a71461013157806304ea59ea146101595780630a605c9b1461016e57806313e7c9d8146101815780631b72a9ff146101a5575b600080fd5b61014461013f36600461160d565b610438565b60405190151581526020015b60405180910390f35b61016c610167366004611671565b61046f565b005b61016c61017c36600461177d565b6104df565b61019461018f366004611799565b610526565b604051610150959493929190611806565b6000546001546002546040805160c0810182526003546001600160401b038082168352600160401b820481166020840152600160801b90910416918101919091526004546001600160a01b03908116606083015260055481166080830152600654811660a083015261022f9481169381811693600160a01b90910465ffffffffffff169291169085565b604080516001600160a01b03968716815294861660208087019190915265ffffffffffff9094168582015291851660608086019190915281516001600160401b0390811660808088019190915294830151811660a080880191909152938301511660c0860152810151851660e0850152918201518416610100840152015190911661012082015261014001610150565b6101446102cd366004611850565b60096020526000908152604090205460ff1681565b6102f56102f0366004611850565b610602565b604051908152602001610150565b61016c610311366004611799565b610624565b61016c610324366004611869565b6106f2565b61016c610337366004611869565b610714565b61014461034a366004611799565b600b6020526000908152604090205460ff1681565b61014461036d366004611799565b600a6020526000908152604090205460ff1681565b60075461039f9065ffffffffffff80821691600160301b90041682565b6040805165ffffffffffff938416815292909116602083015201610150565b61016c6103cc366004611799565b610747565b61016c6103df366004611998565b61080c565b61016c6103f23660046119b5565b610927565b610144610405366004611869565b610a5b565b6102f5600081565b61016c610420366004611869565b610a93565b61016c610433366004611a61565b610aaf565b60006001600160e01b03198216637965db0b60e01b148061046957506301ffc9a760e01b6001600160e01b03198316145b92915050565b600061047a81610fb4565b6001600160a01b0383166000818152600b6020908152604091829020805460ff19168615159081179091558251938452908301527f24171c14ca42bb2ec65eba4633a235ec1ec3a9cc14fdcec3adcf574875fd0b8691015b60405180910390a1505050565b60006104ea81610fb4565b5080516007805460209093015165ffffffffffff908116600160301b026bffffffffffffffffffffffff19909416921691909117919091179055565b600860205260009081526040902080546001820180546001600160a01b03831693600160a01b90930460ff1692919061055e90611b21565b80601f016020809104026020016040519081016040528092919081815260200182805461058a90611b21565b80156105d75780601f106105ac576101008083540402835291602001916105d7565b820191906000526020600020905b8154815290600101906020018083116105ba57829003601f168201915b505050600290930154919250506001600160a01b0381169065ffffffffffff600160a01b9091041685565b6000908152600080516020612021833981519152602052604090206001015490565b6001600160a01b038082166000908152600860205260409020548291166106665760405162461bcd60e51b815260040161065d90611b5b565b60405180910390fd5b600061067181610fb4565b6001600160a01b038316600090815260086020526040902054600160a01b900460ff16156106ed576001600160a01b038316600081815260086020908152604091829020805460ff60a01b1916905590519182527fae02c1bd695006b6d891af37fdeefea45a10ebcc17071e3471787db4f177288591016104d2565b505050565b6106fb82610602565b61070481610fb4565b61070e8383610fc1565b50505050565b6001600160a01b038116331461073d5760405163334bd91960e11b815260040160405180910390fd5b6106ed8282611066565b6001600160a01b038082166000908152600860205260409020548291166107805760405162461bcd60e51b815260040161065d90611b5b565b600061078b81610fb4565b6001600160a01b038316600090815260086020526040902054600160a01b900460ff166106ed576001600160a01b038316600081815260086020908152604091829020805460ff60a01b1916600160a01b17905590519182527fc5437eb8dd091f69800961953f2bb0bc16ae1ff2d3e52caa96796db65f8271da91016104d2565b600061081781610fb4565b508051600080546001600160a01b039283166001600160a01b0319918216179091556020808401516001805460408088015165ffffffffffff16600160a01b026001600160d01b0319909216938716939093171790556060808601516002805491871691861691909117905560809586015180516003805495830151948301516001600160401b03908116600160801b0267ffffffffffffffff60801b19968216600160401b026fffffffffffffffffffffffffffffffff1990981691909316179590951793909316929092179092559081015160048054918516918416919091179055928301516005805491841691831691909117905560a09092015160068054919092169216919091179055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b031660008115801561096c5750825b90506000826001600160401b031660011480156109885750303b155b905081158015610996575080155b156109b45760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156109de57845460ff60401b1916600160401b1785555b6109e66110e2565b6109f1600033610fc1565b506109fa6110ec565b610a038761080c565b610a0c866104df565b8315610a5257845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b6000918252600080516020612021833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b610a9c82610602565b610aa581610fb4565b61070e8383611066565b610ab76110fc565b845160208601206001600160a01b0381163314610b325760405162461bcd60e51b815260206004820152603360248201527f4d616e74615374616b696e674d6964646c65776172653a20696e76616c6964206044820152726f70657261746f72207075626c6963206b657960681b606482015260840161065d565b6000855111610b8f5760405162461bcd60e51b81526020600482015260356024820152600080516020612001833981519152604482015274206e616d652063616e6e6f7420626520656d70747960581b606482015260840161065d565b600085604051602001610ba29190611b98565b60408051601f1981840301815291815281516020928301206000818152600990935291205490915060ff1615610c255760405162461bcd60e51b81526020600482015260346024820152600080516020612001833981519152604482015273206e616d6520616c72656164792065786973747360601b606482015260840161065d565b6000818152600960209081526040808320805460ff1916600117905533835260089091529020546001600160a01b031615610cac5760405162461bcd60e51b8152602060048201526033602482015260008051602061200183398151915260448201527208185b1c9958591e481c9959da5cdd195c9959606a1b606482015260840161065d565b60075465ffffffffffff90811690851610801590610cdf575060075465ffffffffffff600160301b909104811690851611155b610d3e5760405162461bcd60e51b815260206004820152602a60248201527f4d616e74615374616b696e674d6964646c65776172653a20696e76616c69642060448201526931b7b6b6b4b9b9b4b7b760b11b606482015260840161065d565b6001600160a01b0385166000908152600a602052604090205460ff1615610dcd5760405162461bcd60e51b815260206004820152603960248201527f4d616e74615374616b696e674d6964646c65776172653a20726577617264206160448201527f64647265737320616c7265616479207265676973746572656400000000000000606482015260840161065d565b6001600160a01b038086166000908152600a60209081526040808320805460ff191660011790559286168252600b9052205460ff16610e625760405162461bcd60e51b815260206004820152602b60248201527f4d616e74615374616b696e674d6964646c65776172653a20746f6b656e206e6f60448201526a1d081cdd5c1c1bdc9d195960aa1b606482015260840161065d565b610e6b33611146565b6000610e773385611214565b50506040805160a0810182526001600160a01b038084168252600060208084018281528486018e81528d8516606087015265ffffffffffff8d1660808701523384526008909252949091208351815495511515600160a01b026001600160a81b03199096169316929092179390931781559151929350916001820190610efd9082611c04565b5060608201516002909101805460809093015165ffffffffffff16600160a01b026001600160d01b03199093166001600160a01b03909216919091179190911790556040517f82ef2e4bdc58c22c96126c5d61bc476199209d94dc13d6cca98424e7264ee16f90610f799033908b908b908b908b908890611cc3565b60405180910390a1505050610fad60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050505050565b610fbe813361157f565b50565b6000600080516020612021833981519152610fdc8484610a5b565b61105c576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556110123390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610469565b6000915050610469565b60006000805160206120218339815191526110818484610a5b565b1561105c576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610469565b6110ea6115bc565b565b6110f46115bc565b6110ea611605565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080546001190161114057604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6000546040516302910f8b60e31b81526001600160a01b038381166004830152909116906314887c5890602401602060405180830381865afa158015611190573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111b49190611d20565b610fbe5760405162461bcd60e51b815260206004820152603c602482015260008051602061200183398151915260448201527f206e6f74207265676973746572656420746f2073796d62696f74696300000000606482015260840161065d565b60408051610160810182526001600160a01b0383811682526004548116602083015260015465ffffffffffff600160a01b9091041682840152600060608084018290526080840182905260a0840182905260055490921660c084015260e0830181905261010083018190526101208301819052610140830181905283516002808252928101909452928392839290918391816020016020820280368337505060055482519293506001600160a01b0316918391506000906112d7576112d7611d3d565b60200260200101906001600160a01b031690816001600160a01b031681525050308160018151811061130b5761130b611d3d565b6001600160a01b039283166020918202929092018101919091526040805160c0810182526005548416606082018181526006548616608084015260a083018290528252818401869052938b168183015281518083018352600181850190815281528251610100810184526003546001600160401b03168152808501959095528251919490936000939192918301916113a591899101611d53565b60408051808303601f19018152918152908252600354600160401b90046001600160401b0316602080840191909152815192909101916113e791879101611e3b565b60408051808303601f1901815291815290825260016020808401829052600354600160801b90046001600160401b031684840152825187515115158183015283518082039092018252830183526060909301929092529054905163312249f960e21b81529192506001600160a01b03169063c48927e49061146c908490600401611ed5565b6060604051808303816000875af115801561148b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114af9190611fb3565b919950975095506001600160a01b038816158015906114d657506001600160a01b03871615155b80156114ea57506001600160a01b03861615155b61154d5760405162461bcd60e51b815260206004820152602e60248201527f4d616e74615374616b696e674d6964646c65776172653a206661696c6564207460448201526d1bc818dc99585d19481d985d5b1d60921b606482015260840161065d565b50505050509250925092565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6115898282610a5b565b6115b85760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440161065d565b5050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166110ea57604051631afcd79f60e31b815260040160405180910390fd5b6115596115bc565b60006020828403121561161f57600080fd5b81356001600160e01b03198116811461163757600080fd5b9392505050565b6001600160a01b0381168114610fbe57600080fd5b803561165e8161163e565b919050565b8015158114610fbe57600080fd5b6000806040838503121561168457600080fd5b823561168f8161163e565b9150602083013561169f81611663565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b03811182821017156116e2576116e26116aa565b60405290565b60405160c081016001600160401b03811182821017156116e2576116e26116aa565b803565ffffffffffff8116811461165e57600080fd5b60006040828403121561173257600080fd5b604051604081018181106001600160401b0382111715611754576117546116aa565b6040529050806117638361170a565b81526117716020840161170a565b60208201525092915050565b60006040828403121561178f57600080fd5b6116378383611720565b6000602082840312156117ab57600080fd5b81356116378161163e565b60005b838110156117d15781810151838201526020016117b9565b50506000910152565b600081518084526117f28160208601602086016117b6565b601f01601f19169290920160200192915050565b600060018060a01b038088168352861515602084015260a0604084015261183060a08401876117da565b941660608301525065ffffffffffff919091166080909101529392505050565b60006020828403121561186257600080fd5b5035919050565b6000806040838503121561187c57600080fd5b82359150602083013561169f8161163e565b80356001600160401b038116811461165e57600080fd5b60008183036101408112156118b957600080fd5b6118c16116c0565b915082356118ce8161163e565b825260208301356118de8161163e565b60208301526118ef6040840161170a565b604083015260608301356119028161163e565b606083015260c0607f198201121561191957600080fd5b506119226116e8565b61192e6080840161188e565b815261193c60a0840161188e565b602082015261194d60c0840161188e565b604082015260e08301356119608161163e565b60608201526101008301356119748161163e565b60808201526101208301356119888161163e565b60a0820152608082015292915050565b600061014082840312156119ab57600080fd5b61163783836118a5565b60008061018083850312156119c957600080fd5b6119d384846118a5565b91506119e3846101408501611720565b90509250929050565b60006001600160401b0380841115611a0657611a066116aa565b604051601f8501601f19908116603f01168101908282118183101715611a2e57611a2e6116aa565b81604052809350858152868686011115611a4757600080fd5b858560208301376000602087830101525050509392505050565b600080600080600060a08688031215611a7957600080fd5b85356001600160401b0380821115611a9057600080fd5b818801915088601f830112611aa457600080fd5b611ab3898335602085016119ec565b96506020880135915080821115611ac957600080fd5b508601601f81018813611adb57600080fd5b611aea888235602084016119ec565b945050611af960408701611653565b9250611b076060870161170a565b9150611b1560808701611653565b90509295509295909350565b600181811c90821680611b3557607f821691505b602082108103611b5557634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602f9082015260008051602061200183398151915260408201526e081b9bdd081c9959da5cdd195c9959608a1b606082015260800190565b60008251611baa8184602087016117b6565b9190910192915050565b601f8211156106ed576000816000526020600020601f850160051c81016020861015611bdd5750805b601f850160051c820191505b81811015611bfc57828155600101611be9565b505050505050565b81516001600160401b03811115611c1d57611c1d6116aa565b611c3181611c2b8454611b21565b84611bb4565b602080601f831160018114611c665760008415611c4e5750858301515b600019600386901b1c1916600185901b178555611bfc565b600085815260208120601f198616915b82811015611c9557888601518255948401946001909101908401611c76565b5085821015611cb35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600060018060a01b03808916835260c06020840152611ce560c08401896117da565b8381036040850152611cf781896117da565b96821660608501525065ffffffffffff94909416608083015250911660a0909101529392505050565b600060208284031215611d3257600080fd5b815161163781611663565b634e487b7160e01b600052603260045260246000fd5b81516001600160a01b0316815261016081016020830151611d7f60208401826001600160a01b03169052565b506040830151611d99604084018265ffffffffffff169052565b506060830151611dad606084018215159052565b506080830151611dc1608084018215159052565b5060a083015160a083015260c0830151611de660c08401826001600160a01b03169052565b5060e0830151611e0160e08401826001600160a01b03169052565b50610100838101516001600160a01b0390811691840191909152610120808501518216908401526101409384015116929091019190915290565b6020808252825180516001600160a01b039081168484015281830151811660408086019190915290910151811660608401528382015160a06080850152805160c085018190526000939291830191849160e08701905b80841015611eb357845183168252938501936001939093019290850190611e91565b5060408801516001600160a01b03811660a08901529450979650505050505050565b60208152611eef6020820183516001600160401b03169052565b60006020830151611f0b60408401826001600160a01b03169052565b506040830151610100806060850152611f286101208501836117da565b91506060850151611f4460808601826001600160401b03169052565b506080850151601f19808685030160a0870152611f6184836117da565b935060a08701519150611f7860c087018315159052565b60c08701516001600160401b03811660e0880152915060e0870151915080868503018387015250611fa983826117da565b9695505050505050565b600080600060608486031215611fc857600080fd5b8351611fd38161163e565b6020850151909350611fe48161163e565b6040850151909250611ff58161163e565b80915050925092509256fe4d616e74615374616b696e674d6964646c65776172653a206f70657261746f7202dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a264697066735822122064e31378a8649a00580a7a16eff2001f7ba4e7cf35405c73f42b422b20b5717864736f6c63430008190033",
}

// MantaStakingMiddlewareABI is the input ABI used to generate the binding from.
// Deprecated: Use MantaStakingMiddlewareMetaData.ABI instead.
var MantaStakingMiddlewareABI = MantaStakingMiddlewareMetaData.ABI

// MantaStakingMiddlewareBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MantaStakingMiddlewareMetaData.Bin instead.
var MantaStakingMiddlewareBin = MantaStakingMiddlewareMetaData.Bin

// DeployMantaStakingMiddleware deploys a new Ethereum contract, binding an instance of MantaStakingMiddleware to it.
func DeployMantaStakingMiddleware(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MantaStakingMiddleware, error) {
	parsed, err := MantaStakingMiddlewareMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MantaStakingMiddlewareBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MantaStakingMiddleware{MantaStakingMiddlewareCaller: MantaStakingMiddlewareCaller{contract: contract}, MantaStakingMiddlewareTransactor: MantaStakingMiddlewareTransactor{contract: contract}, MantaStakingMiddlewareFilterer: MantaStakingMiddlewareFilterer{contract: contract}}, nil
}

// MantaStakingMiddleware is an auto generated Go binding around an Ethereum contract.
type MantaStakingMiddleware struct {
	MantaStakingMiddlewareCaller     // Read-only binding to the contract
	MantaStakingMiddlewareTransactor // Write-only binding to the contract
	MantaStakingMiddlewareFilterer   // Log filterer for contract events
}

// MantaStakingMiddlewareCaller is an auto generated read-only Go binding around an Ethereum contract.
type MantaStakingMiddlewareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MantaStakingMiddlewareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MantaStakingMiddlewareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MantaStakingMiddlewareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MantaStakingMiddlewareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MantaStakingMiddlewareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MantaStakingMiddlewareSession struct {
	Contract     *MantaStakingMiddleware // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MantaStakingMiddlewareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MantaStakingMiddlewareCallerSession struct {
	Contract *MantaStakingMiddlewareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MantaStakingMiddlewareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MantaStakingMiddlewareTransactorSession struct {
	Contract     *MantaStakingMiddlewareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MantaStakingMiddlewareRaw is an auto generated low-level Go binding around an Ethereum contract.
type MantaStakingMiddlewareRaw struct {
	Contract *MantaStakingMiddleware // Generic contract binding to access the raw methods on
}

// MantaStakingMiddlewareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MantaStakingMiddlewareCallerRaw struct {
	Contract *MantaStakingMiddlewareCaller // Generic read-only contract binding to access the raw methods on
}

// MantaStakingMiddlewareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MantaStakingMiddlewareTransactorRaw struct {
	Contract *MantaStakingMiddlewareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMantaStakingMiddleware creates a new instance of MantaStakingMiddleware, bound to a specific deployed contract.
func NewMantaStakingMiddleware(address common.Address, backend bind.ContractBackend) (*MantaStakingMiddleware, error) {
	contract, err := bindMantaStakingMiddleware(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddleware{MantaStakingMiddlewareCaller: MantaStakingMiddlewareCaller{contract: contract}, MantaStakingMiddlewareTransactor: MantaStakingMiddlewareTransactor{contract: contract}, MantaStakingMiddlewareFilterer: MantaStakingMiddlewareFilterer{contract: contract}}, nil
}

// NewMantaStakingMiddlewareCaller creates a new read-only instance of MantaStakingMiddleware, bound to a specific deployed contract.
func NewMantaStakingMiddlewareCaller(address common.Address, caller bind.ContractCaller) (*MantaStakingMiddlewareCaller, error) {
	contract, err := bindMantaStakingMiddleware(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareCaller{contract: contract}, nil
}

// NewMantaStakingMiddlewareTransactor creates a new write-only instance of MantaStakingMiddleware, bound to a specific deployed contract.
func NewMantaStakingMiddlewareTransactor(address common.Address, transactor bind.ContractTransactor) (*MantaStakingMiddlewareTransactor, error) {
	contract, err := bindMantaStakingMiddleware(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareTransactor{contract: contract}, nil
}

// NewMantaStakingMiddlewareFilterer creates a new log filterer instance of MantaStakingMiddleware, bound to a specific deployed contract.
func NewMantaStakingMiddlewareFilterer(address common.Address, filterer bind.ContractFilterer) (*MantaStakingMiddlewareFilterer, error) {
	contract, err := bindMantaStakingMiddleware(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareFilterer{contract: contract}, nil
}

// bindMantaStakingMiddleware binds a generic wrapper to an already deployed contract.
func bindMantaStakingMiddleware(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MantaStakingMiddlewareMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MantaStakingMiddleware *MantaStakingMiddlewareRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MantaStakingMiddleware.Contract.MantaStakingMiddlewareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MantaStakingMiddleware *MantaStakingMiddlewareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.MantaStakingMiddlewareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MantaStakingMiddleware *MantaStakingMiddlewareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.MantaStakingMiddlewareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MantaStakingMiddleware.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MantaStakingMiddleware.Contract.DEFAULTADMINROLE(&_MantaStakingMiddleware.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MantaStakingMiddleware.Contract.DEFAULTADMINROLE(&_MantaStakingMiddleware.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MantaStakingMiddleware.Contract.GetRoleAdmin(&_MantaStakingMiddleware.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MantaStakingMiddleware.Contract.GetRoleAdmin(&_MantaStakingMiddleware.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MantaStakingMiddleware.Contract.HasRole(&_MantaStakingMiddleware.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MantaStakingMiddleware.Contract.HasRole(&_MantaStakingMiddleware.CallOpts, role, account)
}

// OperatorNameExists is a free data retrieval call binding the contract method 0x1d6a65d7.
//
// Solidity: function operatorNameExists(bytes32 ) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) OperatorNameExists(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "operatorNameExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorNameExists is a free data retrieval call binding the contract method 0x1d6a65d7.
//
// Solidity: function operatorNameExists(bytes32 ) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) OperatorNameExists(arg0 [32]byte) (bool, error) {
	return _MantaStakingMiddleware.Contract.OperatorNameExists(&_MantaStakingMiddleware.CallOpts, arg0)
}

// OperatorNameExists is a free data retrieval call binding the contract method 0x1d6a65d7.
//
// Solidity: function operatorNameExists(bytes32 ) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) OperatorNameExists(arg0 [32]byte) (bool, error) {
	return _MantaStakingMiddleware.Contract.OperatorNameExists(&_MantaStakingMiddleware.CallOpts, arg0)
}

// OperatorSettings is a free data retrieval call binding the contract method 0x72aef6bf.
//
// Solidity: function operatorSettings() view returns(uint48 minOperatorCommission, uint48 maxOperatorCommission)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) OperatorSettings(opts *bind.CallOpts) (struct {
	MinOperatorCommission *big.Int
	MaxOperatorCommission *big.Int
}, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "operatorSettings")

	outstruct := new(struct {
		MinOperatorCommission *big.Int
		MaxOperatorCommission *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinOperatorCommission = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxOperatorCommission = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OperatorSettings is a free data retrieval call binding the contract method 0x72aef6bf.
//
// Solidity: function operatorSettings() view returns(uint48 minOperatorCommission, uint48 maxOperatorCommission)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) OperatorSettings() (struct {
	MinOperatorCommission *big.Int
	MaxOperatorCommission *big.Int
}, error) {
	return _MantaStakingMiddleware.Contract.OperatorSettings(&_MantaStakingMiddleware.CallOpts)
}

// OperatorSettings is a free data retrieval call binding the contract method 0x72aef6bf.
//
// Solidity: function operatorSettings() view returns(uint48 minOperatorCommission, uint48 maxOperatorCommission)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) OperatorSettings() (struct {
	MinOperatorCommission *big.Int
	MaxOperatorCommission *big.Int
}, error) {
	return _MantaStakingMiddleware.Contract.OperatorSettings(&_MantaStakingMiddleware.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(address vault, bool paused, string operatorName, address rewardAddress, uint48 commission)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Vault         common.Address
	Paused        bool
	OperatorName  string
	RewardAddress common.Address
	Commission    *big.Int
}, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "operators", arg0)

	outstruct := new(struct {
		Vault         common.Address
		Paused        bool
		OperatorName  string
		RewardAddress common.Address
		Commission    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Vault = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Paused = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.OperatorName = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.RewardAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Commission = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(address vault, bool paused, string operatorName, address rewardAddress, uint48 commission)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) Operators(arg0 common.Address) (struct {
	Vault         common.Address
	Paused        bool
	OperatorName  string
	RewardAddress common.Address
	Commission    *big.Int
}, error) {
	return _MantaStakingMiddleware.Contract.Operators(&_MantaStakingMiddleware.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(address vault, bool paused, string operatorName, address rewardAddress, uint48 commission)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) Operators(arg0 common.Address) (struct {
	Vault         common.Address
	Paused        bool
	OperatorName  string
	RewardAddress common.Address
	Commission    *big.Int
}, error) {
	return _MantaStakingMiddleware.Contract.Operators(&_MantaStakingMiddleware.CallOpts, arg0)
}

// RewardAddressExists is a free data retrieval call binding the contract method 0x7174d669.
//
// Solidity: function rewardAddressExists(address ) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) RewardAddressExists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "rewardAddressExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardAddressExists is a free data retrieval call binding the contract method 0x7174d669.
//
// Solidity: function rewardAddressExists(address ) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) RewardAddressExists(arg0 common.Address) (bool, error) {
	return _MantaStakingMiddleware.Contract.RewardAddressExists(&_MantaStakingMiddleware.CallOpts, arg0)
}

// RewardAddressExists is a free data retrieval call binding the contract method 0x7174d669.
//
// Solidity: function rewardAddressExists(address ) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) RewardAddressExists(arg0 common.Address) (bool, error) {
	return _MantaStakingMiddleware.Contract.RewardAddressExists(&_MantaStakingMiddleware.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) SupportedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "supportedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) SupportedTokens(arg0 common.Address) (bool, error) {
	return _MantaStakingMiddleware.Contract.SupportedTokens(&_MantaStakingMiddleware.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) SupportedTokens(arg0 common.Address) (bool, error) {
	return _MantaStakingMiddleware.Contract.SupportedTokens(&_MantaStakingMiddleware.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MantaStakingMiddleware.Contract.SupportsInterface(&_MantaStakingMiddleware.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MantaStakingMiddleware.Contract.SupportsInterface(&_MantaStakingMiddleware.CallOpts, interfaceId)
}

// SymbioticVaultSettings is a free data retrieval call binding the contract method 0x1b72a9ff.
//
// Solidity: function symbioticVaultSettings() view returns(address operatorRegistry, address vaultConfiguration, uint48 epochDuration, address stakeToken, (uint64,uint64,uint64,address,address,address) defaultVaultParams)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) SymbioticVaultSettings(opts *bind.CallOpts) (struct {
	OperatorRegistry   common.Address
	VaultConfiguration common.Address
	EpochDuration      *big.Int
	StakeToken         common.Address
	DefaultVaultParams DefaultVaultInitParams
}, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "symbioticVaultSettings")

	outstruct := new(struct {
		OperatorRegistry   common.Address
		VaultConfiguration common.Address
		EpochDuration      *big.Int
		StakeToken         common.Address
		DefaultVaultParams DefaultVaultInitParams
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorRegistry = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.VaultConfiguration = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.EpochDuration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StakeToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.DefaultVaultParams = *abi.ConvertType(out[4], new(DefaultVaultInitParams)).(*DefaultVaultInitParams)

	return *outstruct, err

}

// SymbioticVaultSettings is a free data retrieval call binding the contract method 0x1b72a9ff.
//
// Solidity: function symbioticVaultSettings() view returns(address operatorRegistry, address vaultConfiguration, uint48 epochDuration, address stakeToken, (uint64,uint64,uint64,address,address,address) defaultVaultParams)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) SymbioticVaultSettings() (struct {
	OperatorRegistry   common.Address
	VaultConfiguration common.Address
	EpochDuration      *big.Int
	StakeToken         common.Address
	DefaultVaultParams DefaultVaultInitParams
}, error) {
	return _MantaStakingMiddleware.Contract.SymbioticVaultSettings(&_MantaStakingMiddleware.CallOpts)
}

// SymbioticVaultSettings is a free data retrieval call binding the contract method 0x1b72a9ff.
//
// Solidity: function symbioticVaultSettings() view returns(address operatorRegistry, address vaultConfiguration, uint48 epochDuration, address stakeToken, (uint64,uint64,uint64,address,address,address) defaultVaultParams)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) SymbioticVaultSettings() (struct {
	OperatorRegistry   common.Address
	VaultConfiguration common.Address
	EpochDuration      *big.Int
	StakeToken         common.Address
	DefaultVaultParams DefaultVaultInitParams
}, error) {
	return _MantaStakingMiddleware.Contract.SymbioticVaultSettings(&_MantaStakingMiddleware.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.GrantRole(&_MantaStakingMiddleware.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.GrantRole(&_MantaStakingMiddleware.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x85eb242d.
//
// Solidity: function initialize((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) _symbioticVaultSettings, (uint48,uint48) _operatorSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) Initialize(opts *bind.TransactOpts, _symbioticVaultSettings SymbioticVaultSettings, _operatorSettings OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "initialize", _symbioticVaultSettings, _operatorSettings)
}

// Initialize is a paid mutator transaction binding the contract method 0x85eb242d.
//
// Solidity: function initialize((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) _symbioticVaultSettings, (uint48,uint48) _operatorSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) Initialize(_symbioticVaultSettings SymbioticVaultSettings, _operatorSettings OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Initialize(&_MantaStakingMiddleware.TransactOpts, _symbioticVaultSettings, _operatorSettings)
}

// Initialize is a paid mutator transaction binding the contract method 0x85eb242d.
//
// Solidity: function initialize((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) _symbioticVaultSettings, (uint48,uint48) _operatorSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) Initialize(_symbioticVaultSettings SymbioticVaultSettings, _operatorSettings OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Initialize(&_MantaStakingMiddleware.TransactOpts, _symbioticVaultSettings, _operatorSettings)
}

// ModifySupportedTokens is a paid mutator transaction binding the contract method 0x04ea59ea.
//
// Solidity: function modifySupportedTokens(address _token, bool _supported) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) ModifySupportedTokens(opts *bind.TransactOpts, _token common.Address, _supported bool) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "modifySupportedTokens", _token, _supported)
}

// ModifySupportedTokens is a paid mutator transaction binding the contract method 0x04ea59ea.
//
// Solidity: function modifySupportedTokens(address _token, bool _supported) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) ModifySupportedTokens(_token common.Address, _supported bool) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.ModifySupportedTokens(&_MantaStakingMiddleware.TransactOpts, _token, _supported)
}

// ModifySupportedTokens is a paid mutator transaction binding the contract method 0x04ea59ea.
//
// Solidity: function modifySupportedTokens(address _token, bool _supported) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) ModifySupportedTokens(_token common.Address, _supported bool) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.ModifySupportedTokens(&_MantaStakingMiddleware.TransactOpts, _token, _supported)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address _operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) PauseOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "pauseOperator", _operator)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address _operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) PauseOperator(_operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.PauseOperator(&_MantaStakingMiddleware.TransactOpts, _operator)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address _operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) PauseOperator(_operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.PauseOperator(&_MantaStakingMiddleware.TransactOpts, _operator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xecfb6e8d.
//
// Solidity: function registerOperator(bytes _operatorPublicKey, string _operatorName, address _rewardAddress, uint48 _commission, address _token) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) RegisterOperator(opts *bind.TransactOpts, _operatorPublicKey []byte, _operatorName string, _rewardAddress common.Address, _commission *big.Int, _token common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "registerOperator", _operatorPublicKey, _operatorName, _rewardAddress, _commission, _token)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xecfb6e8d.
//
// Solidity: function registerOperator(bytes _operatorPublicKey, string _operatorName, address _rewardAddress, uint48 _commission, address _token) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) RegisterOperator(_operatorPublicKey []byte, _operatorName string, _rewardAddress common.Address, _commission *big.Int, _token common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RegisterOperator(&_MantaStakingMiddleware.TransactOpts, _operatorPublicKey, _operatorName, _rewardAddress, _commission, _token)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xecfb6e8d.
//
// Solidity: function registerOperator(bytes _operatorPublicKey, string _operatorName, address _rewardAddress, uint48 _commission, address _token) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) RegisterOperator(_operatorPublicKey []byte, _operatorName string, _rewardAddress common.Address, _commission *big.Int, _token common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RegisterOperator(&_MantaStakingMiddleware.TransactOpts, _operatorPublicKey, _operatorName, _rewardAddress, _commission, _token)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RenounceRole(&_MantaStakingMiddleware.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RenounceRole(&_MantaStakingMiddleware.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RevokeRole(&_MantaStakingMiddleware.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RevokeRole(&_MantaStakingMiddleware.TransactOpts, role, account)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address _operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) UnpauseOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "unpauseOperator", _operator)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address _operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) UnpauseOperator(_operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UnpauseOperator(&_MantaStakingMiddleware.TransactOpts, _operator)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address _operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) UnpauseOperator(_operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UnpauseOperator(&_MantaStakingMiddleware.TransactOpts, _operator)
}

// UpdateOperatorSettings is a paid mutator transaction binding the contract method 0x0a605c9b.
//
// Solidity: function updateOperatorSettings((uint48,uint48) _operatorSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) UpdateOperatorSettings(opts *bind.TransactOpts, _operatorSettings OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "updateOperatorSettings", _operatorSettings)
}

// UpdateOperatorSettings is a paid mutator transaction binding the contract method 0x0a605c9b.
//
// Solidity: function updateOperatorSettings((uint48,uint48) _operatorSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) UpdateOperatorSettings(_operatorSettings OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UpdateOperatorSettings(&_MantaStakingMiddleware.TransactOpts, _operatorSettings)
}

// UpdateOperatorSettings is a paid mutator transaction binding the contract method 0x0a605c9b.
//
// Solidity: function updateOperatorSettings((uint48,uint48) _operatorSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) UpdateOperatorSettings(_operatorSettings OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UpdateOperatorSettings(&_MantaStakingMiddleware.TransactOpts, _operatorSettings)
}

// UpdateSymbioticVaultSettings is a paid mutator transaction binding the contract method 0x7a6efb69.
//
// Solidity: function updateSymbioticVaultSettings((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) _symbioticVaultSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) UpdateSymbioticVaultSettings(opts *bind.TransactOpts, _symbioticVaultSettings SymbioticVaultSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "updateSymbioticVaultSettings", _symbioticVaultSettings)
}

// UpdateSymbioticVaultSettings is a paid mutator transaction binding the contract method 0x7a6efb69.
//
// Solidity: function updateSymbioticVaultSettings((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) _symbioticVaultSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) UpdateSymbioticVaultSettings(_symbioticVaultSettings SymbioticVaultSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UpdateSymbioticVaultSettings(&_MantaStakingMiddleware.TransactOpts, _symbioticVaultSettings)
}

// UpdateSymbioticVaultSettings is a paid mutator transaction binding the contract method 0x7a6efb69.
//
// Solidity: function updateSymbioticVaultSettings((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) _symbioticVaultSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) UpdateSymbioticVaultSettings(_symbioticVaultSettings SymbioticVaultSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UpdateSymbioticVaultSettings(&_MantaStakingMiddleware.TransactOpts, _symbioticVaultSettings)
}

// MantaStakingMiddlewareInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareInitializedIterator struct {
	Event *MantaStakingMiddlewareInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareInitialized represents a Initialized event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterInitialized(opts *bind.FilterOpts) (*MantaStakingMiddlewareInitializedIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareInitializedIterator{contract: _MantaStakingMiddleware.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareInitialized) (event.Subscription, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareInitialized)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseInitialized(log types.Log) (*MantaStakingMiddlewareInitialized, error) {
	event := new(MantaStakingMiddlewareInitialized)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareOperatorPausedIterator is returned from FilterOperatorPaused and is used to iterate over the raw logs and unpacked data for OperatorPaused events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorPausedIterator struct {
	Event *MantaStakingMiddlewareOperatorPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareOperatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareOperatorPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareOperatorPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareOperatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareOperatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareOperatorPaused represents a OperatorPaused event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorPaused struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorPaused is a free log retrieval operation binding the contract event 0xc5437eb8dd091f69800961953f2bb0bc16ae1ff2d3e52caa96796db65f8271da.
//
// Solidity: event OperatorPaused(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterOperatorPaused(opts *bind.FilterOpts) (*MantaStakingMiddlewareOperatorPausedIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "OperatorPaused")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareOperatorPausedIterator{contract: _MantaStakingMiddleware.contract, event: "OperatorPaused", logs: logs, sub: sub}, nil
}

// WatchOperatorPaused is a free log subscription operation binding the contract event 0xc5437eb8dd091f69800961953f2bb0bc16ae1ff2d3e52caa96796db65f8271da.
//
// Solidity: event OperatorPaused(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchOperatorPaused(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareOperatorPaused) (event.Subscription, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "OperatorPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareOperatorPaused)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorPaused is a log parse operation binding the contract event 0xc5437eb8dd091f69800961953f2bb0bc16ae1ff2d3e52caa96796db65f8271da.
//
// Solidity: event OperatorPaused(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseOperatorPaused(log types.Log) (*MantaStakingMiddlewareOperatorPaused, error) {
	event := new(MantaStakingMiddlewareOperatorPaused)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorRegisteredIterator struct {
	Event *MantaStakingMiddlewareOperatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareOperatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareOperatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareOperatorRegistered represents a OperatorRegistered event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorRegistered struct {
	Operator          common.Address
	OperatorPublicKey []byte
	OperatorName      string
	RewardAddress     common.Address
	Commission        *big.Int
	Vault             common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x82ef2e4bdc58c22c96126c5d61bc476199209d94dc13d6cca98424e7264ee16f.
//
// Solidity: event OperatorRegistered(address operator, bytes operatorPublicKey, string operatorName, address rewardAddress, uint48 commission, address vault)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterOperatorRegistered(opts *bind.FilterOpts) (*MantaStakingMiddlewareOperatorRegisteredIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "OperatorRegistered")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareOperatorRegisteredIterator{contract: _MantaStakingMiddleware.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x82ef2e4bdc58c22c96126c5d61bc476199209d94dc13d6cca98424e7264ee16f.
//
// Solidity: event OperatorRegistered(address operator, bytes operatorPublicKey, string operatorName, address rewardAddress, uint48 commission, address vault)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareOperatorRegistered) (event.Subscription, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "OperatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareOperatorRegistered)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRegistered is a log parse operation binding the contract event 0x82ef2e4bdc58c22c96126c5d61bc476199209d94dc13d6cca98424e7264ee16f.
//
// Solidity: event OperatorRegistered(address operator, bytes operatorPublicKey, string operatorName, address rewardAddress, uint48 commission, address vault)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseOperatorRegistered(log types.Log) (*MantaStakingMiddlewareOperatorRegistered, error) {
	event := new(MantaStakingMiddlewareOperatorRegistered)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareOperatorUnpausedIterator is returned from FilterOperatorUnpaused and is used to iterate over the raw logs and unpacked data for OperatorUnpaused events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorUnpausedIterator struct {
	Event *MantaStakingMiddlewareOperatorUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareOperatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareOperatorUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareOperatorUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareOperatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareOperatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareOperatorUnpaused represents a OperatorUnpaused event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorUnpaused struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnpaused is a free log retrieval operation binding the contract event 0xae02c1bd695006b6d891af37fdeefea45a10ebcc17071e3471787db4f1772885.
//
// Solidity: event OperatorUnpaused(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterOperatorUnpaused(opts *bind.FilterOpts) (*MantaStakingMiddlewareOperatorUnpausedIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "OperatorUnpaused")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareOperatorUnpausedIterator{contract: _MantaStakingMiddleware.contract, event: "OperatorUnpaused", logs: logs, sub: sub}, nil
}

// WatchOperatorUnpaused is a free log subscription operation binding the contract event 0xae02c1bd695006b6d891af37fdeefea45a10ebcc17071e3471787db4f1772885.
//
// Solidity: event OperatorUnpaused(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchOperatorUnpaused(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareOperatorUnpaused) (event.Subscription, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "OperatorUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareOperatorUnpaused)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorUnpaused is a log parse operation binding the contract event 0xae02c1bd695006b6d891af37fdeefea45a10ebcc17071e3471787db4f1772885.
//
// Solidity: event OperatorUnpaused(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseOperatorUnpaused(log types.Log) (*MantaStakingMiddlewareOperatorUnpaused, error) {
	event := new(MantaStakingMiddlewareOperatorUnpaused)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareOperatorUnregisteredIterator is returned from FilterOperatorUnregistered and is used to iterate over the raw logs and unpacked data for OperatorUnregistered events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorUnregisteredIterator struct {
	Event *MantaStakingMiddlewareOperatorUnregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareOperatorUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareOperatorUnregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareOperatorUnregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareOperatorUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareOperatorUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareOperatorUnregistered represents a OperatorUnregistered event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorUnregistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnregistered is a free log retrieval operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterOperatorUnregistered(opts *bind.FilterOpts) (*MantaStakingMiddlewareOperatorUnregisteredIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "OperatorUnregistered")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareOperatorUnregisteredIterator{contract: _MantaStakingMiddleware.contract, event: "OperatorUnregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorUnregistered is a free log subscription operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchOperatorUnregistered(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareOperatorUnregistered) (event.Subscription, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "OperatorUnregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareOperatorUnregistered)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorUnregistered is a log parse operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseOperatorUnregistered(log types.Log) (*MantaStakingMiddlewareOperatorUnregistered, error) {
	event := new(MantaStakingMiddlewareOperatorUnregistered)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareRewardAddressSetIterator is returned from FilterRewardAddressSet and is used to iterate over the raw logs and unpacked data for RewardAddressSet events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRewardAddressSetIterator struct {
	Event *MantaStakingMiddlewareRewardAddressSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareRewardAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareRewardAddressSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareRewardAddressSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareRewardAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareRewardAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareRewardAddressSet represents a RewardAddressSet event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRewardAddressSet struct {
	Operator      common.Address
	RewardAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardAddressSet is a free log retrieval operation binding the contract event 0x490e4e7668b9ec6e5180b1fb3f783a4f81efbb160c2691224937832c195b64ce.
//
// Solidity: event RewardAddressSet(address operator, address rewardAddress)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterRewardAddressSet(opts *bind.FilterOpts) (*MantaStakingMiddlewareRewardAddressSetIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "RewardAddressSet")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareRewardAddressSetIterator{contract: _MantaStakingMiddleware.contract, event: "RewardAddressSet", logs: logs, sub: sub}, nil
}

// WatchRewardAddressSet is a free log subscription operation binding the contract event 0x490e4e7668b9ec6e5180b1fb3f783a4f81efbb160c2691224937832c195b64ce.
//
// Solidity: event RewardAddressSet(address operator, address rewardAddress)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchRewardAddressSet(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareRewardAddressSet) (event.Subscription, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "RewardAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareRewardAddressSet)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RewardAddressSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardAddressSet is a log parse operation binding the contract event 0x490e4e7668b9ec6e5180b1fb3f783a4f81efbb160c2691224937832c195b64ce.
//
// Solidity: event RewardAddressSet(address operator, address rewardAddress)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseRewardAddressSet(log types.Log) (*MantaStakingMiddlewareRewardAddressSet, error) {
	event := new(MantaStakingMiddlewareRewardAddressSet)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RewardAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRoleAdminChangedIterator struct {
	Event *MantaStakingMiddlewareRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareRoleAdminChanged represents a RoleAdminChanged event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MantaStakingMiddlewareRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareRoleAdminChangedIterator{contract: _MantaStakingMiddleware.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareRoleAdminChanged)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseRoleAdminChanged(log types.Log) (*MantaStakingMiddlewareRoleAdminChanged, error) {
	event := new(MantaStakingMiddlewareRoleAdminChanged)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRoleGrantedIterator struct {
	Event *MantaStakingMiddlewareRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareRoleGranted represents a RoleGranted event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MantaStakingMiddlewareRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareRoleGrantedIterator{contract: _MantaStakingMiddleware.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareRoleGranted)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseRoleGranted(log types.Log) (*MantaStakingMiddlewareRoleGranted, error) {
	event := new(MantaStakingMiddlewareRoleGranted)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRoleRevokedIterator struct {
	Event *MantaStakingMiddlewareRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareRoleRevoked represents a RoleRevoked event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MantaStakingMiddlewareRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareRoleRevokedIterator{contract: _MantaStakingMiddleware.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareRoleRevoked)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseRoleRevoked(log types.Log) (*MantaStakingMiddlewareRoleRevoked, error) {
	event := new(MantaStakingMiddlewareRoleRevoked)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareSupportedTokensModifiedIterator is returned from FilterSupportedTokensModified and is used to iterate over the raw logs and unpacked data for SupportedTokensModified events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareSupportedTokensModifiedIterator struct {
	Event *MantaStakingMiddlewareSupportedTokensModified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareSupportedTokensModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareSupportedTokensModified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareSupportedTokensModified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareSupportedTokensModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareSupportedTokensModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareSupportedTokensModified represents a SupportedTokensModified event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareSupportedTokensModified struct {
	Token     common.Address
	Supported bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSupportedTokensModified is a free log retrieval operation binding the contract event 0x24171c14ca42bb2ec65eba4633a235ec1ec3a9cc14fdcec3adcf574875fd0b86.
//
// Solidity: event SupportedTokensModified(address token, bool supported)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterSupportedTokensModified(opts *bind.FilterOpts) (*MantaStakingMiddlewareSupportedTokensModifiedIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "SupportedTokensModified")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareSupportedTokensModifiedIterator{contract: _MantaStakingMiddleware.contract, event: "SupportedTokensModified", logs: logs, sub: sub}, nil
}

// WatchSupportedTokensModified is a free log subscription operation binding the contract event 0x24171c14ca42bb2ec65eba4633a235ec1ec3a9cc14fdcec3adcf574875fd0b86.
//
// Solidity: event SupportedTokensModified(address token, bool supported)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchSupportedTokensModified(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareSupportedTokensModified) (event.Subscription, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "SupportedTokensModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareSupportedTokensModified)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "SupportedTokensModified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSupportedTokensModified is a log parse operation binding the contract event 0x24171c14ca42bb2ec65eba4633a235ec1ec3a9cc14fdcec3adcf574875fd0b86.
//
// Solidity: event SupportedTokensModified(address token, bool supported)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseSupportedTokensModified(log types.Log) (*MantaStakingMiddlewareSupportedTokensModified, error) {
	event := new(MantaStakingMiddlewareSupportedTokensModified)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "SupportedTokensModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
