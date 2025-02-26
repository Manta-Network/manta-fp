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
	UnregisterTokenUnlockWindow *big.Int
	RequiredOperatorStake       *big.Int
	MinOperatorCommission       *big.Int
	MaxOperatorCommission       *big.Int
	MaxValidOperator            *big.Int
}

// RewardSettings is an auto generated low-level Go binding around an user-defined struct.
type RewardSettings struct {
	OperatorRewardDistributor common.Address
	StakerRewardDistributor   common.Address
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REWARD_UPDATER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimUnlockedToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"distributeRewards\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorRewardAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stakerRewardAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"captureTimestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"operatorRewardRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"symbioticVaultSettings_\",\"type\":\"tuple\",\"internalType\":\"structSymbioticVaultSettings\",\"components\":[{\"name\":\"operatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultConfiguration\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"stakeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultVaultParams\",\"type\":\"tuple\",\"internalType\":\"structDefaultVaultInitParams\",\"components\":[{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delegatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slasherIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"vaultSlashBurner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultDefaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeforeSlashHook\",\"type\":\"address\",\"internalType\":\"address\"}]}]},{\"name\":\"operatorSettings_\",\"type\":\"tuple\",\"internalType\":\"structOperatorSettings\",\"components\":[{\"name\":\"unregisterTokenUnlockWindow\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"requiredOperatorStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxValidOperator\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"name\":\"rewardSettings_\",\"type\":\"tuple\",\"internalType\":\"structRewardSettings\",\"components\":[{\"name\":\"operatorRewardDistributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerRewardDistributor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorNameExists\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSettings\",\"inputs\":[],\"outputs\":[{\"name\":\"unregisterTokenUnlockWindow\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"requiredOperatorStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxValidOperator\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorTokenUnlockTimestamps\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operators\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"operatorName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"commission\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operatorName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"commission\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardSettings\",\"inputs\":[],\"outputs\":[{\"name\":\"operatorRewardDistributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerRewardDistributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slash\",\"inputs\":[{\"name\":\"subnetwork\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"captureTimestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[{\"name\":\"slashedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbioticVaultSettings\",\"inputs\":[],\"outputs\":[{\"name\":\"operatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultConfiguration\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"stakeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultVaultParams\",\"type\":\"tuple\",\"internalType\":\"structDefaultVaultInitParams\",\"components\":[{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delegatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slasherIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"vaultSlashBurner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultDefaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeforeSlashHook\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpauseOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterOperator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorSettings\",\"inputs\":[{\"name\":\"operatorSettings_\",\"type\":\"tuple\",\"internalType\":\"structOperatorSettings\",\"components\":[{\"name\":\"unregisterTokenUnlockWindow\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"requiredOperatorStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxValidOperator\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateRewardSettings\",\"inputs\":[{\"name\":\"rewardSettings_\",\"type\":\"tuple\",\"internalType\":\"structRewardSettings\",\"components\":[{\"name\":\"operatorRewardDistributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerRewardDistributor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSymbioticVaultSettings\",\"inputs\":[{\"name\":\"symbioticVaultSettings_\",\"type\":\"tuple\",\"internalType\":\"structSymbioticVaultSettings\",\"components\":[{\"name\":\"operatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultConfiguration\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"stakeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultVaultParams\",\"type\":\"tuple\",\"internalType\":\"structDefaultVaultInitParams\",\"components\":[{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delegatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slasherIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"vaultSlashBurner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultDefaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeforeSlashHook\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorPaused\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"commission\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnpaused\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CannotClaimBeforeUnlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUnregisterPausedOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedToCreateVault\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommission\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorIsNotPendingUnlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorIsPendingUnlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNameAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegisteredToSymbiotic\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b612630806100d96000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c806372aef6bf116100de578063a217fddf11610097578063c15025c611610071578063c15025c61461049f578063d547741f146104d9578063f3f2c422146104ec578063ff4e42d71461052b57600080fd5b8063a217fddf1461047c578063a314db1614610484578063a876b89a1461049757600080fd5b806372aef6bf146103bd57806372f9adab146104285780637a6efb691461043b57806391d148541461044e578063989bfec514610461578063a0f98db61461047457600080fd5b80631d6a65d7116101305780631d6a65d71461033b578063248a9ca31461035e5780632e5aaf33146103715780632f2ff15d1461038457806336568abe1461039757806360ba757d146103aa57600080fd5b806301ffc9a7146101785780630c761c97146101a057806312f09ad7146101b557806313e7c9d8146101c85780631a69d9c8146101ec5780631b72a9ff14610221575b600080fd5b61018b610186366004611a4e565b61053e565b60405190151581526020015b60405180910390f35b6101b36101ae366004611b96565b610575565b005b6101b36101c3366004611c32565b6105b6565b6101db6101d6366004611c4e565b61063a565b604051610197959493929190611cbb565b6102137f9188644bf0c7a694e572b54fd40005e1230f80a50c59be6fb567a312ab5a1d4d81565b604051908152602001610197565b6000546001546002546040805160c0810182526003546001600160401b038082168352600160401b820481166020840152600160801b90910416918101919091526004546001600160a01b03908116606083015260055481166080830152600654811660a08301526102ab9481169381811693600160a01b90910465ffffffffffff169291169085565b604080516001600160a01b03968716815294861660208087019190915265ffffffffffff9094168582015291851660608086019190915281516001600160401b0390811660808088019190915294830151811660a080880191909152938301511660c0860152810151851660e0850152918201518416610100840152015190911661012082015261014001610197565b61018b610349366004611d05565b600d6020526000908152604090205460ff1681565b61021361036c366004611d05565b610716565b6101b361037f366004611c4e565b610738565b6101b3610392366004611d1e565b610803565b6101b36103a5366004611d1e565b610825565b6101b36103b8366004611d4e565b610858565b6007546008546009546103ef9265ffffffffffff908116929180821691600160301b8204811691600160601b90041685565b6040805165ffffffffffff968716815260208101959095529285169284019290925283166060830152909116608082015260a001610197565b6101b3610436366004611c4e565b610ae3565b6101b3610449366004611f0d565b610ba9565b61018b61045c366004611d1e565b610cc4565b6101b361046f366004611f2a565b610cfc565b6101b3610e3a565b610213600081565b610213610492366004611f6a565b610ef3565b6101b3611046565b600a54600b546104b9916001600160a01b03908116911682565b604080516001600160a01b03938416815292909116602083015201610197565b6101b36104e7366004611d1e565b61118e565b6105146104fa366004611c4e565b600e6020526000908152604090205465ffffffffffff1681565b60405165ffffffffffff9091168152602001610197565b6101b3610539366004611fb2565b6111aa565b60006001600160e01b03198216637965db0b60e01b148061056f57506301ffc9a760e01b6001600160e01b03198316145b92915050565b600061058081611302565b508051600a80546001600160a01b03199081166001600160a01b0393841617909155602090920151600b80549093169116179055565b60006105c181611302565b5080516007805465ffffffffffff191665ffffffffffff928316179055602082015160085560408201516009805460608501516080909501519284166bffffffffffffffffffffffff1990911617600160301b948416949094029390931765ffffffffffff60601b1916600160601b9190921602179055565b600c60205260009081526040902080546001820180546001600160a01b03831693600160a01b90930460ff1692919061067290612015565b80601f016020809104026020016040519081016040528092919081815260200182805461069e90612015565b80156106eb5780601f106106c0576101008083540402835291602001916106eb565b820191906000526020600020905b8154815290600101906020018083116106ce57829003601f168201915b505050600290930154919250506001600160a01b0381169065ffffffffffff600160a01b9091041685565b60009081526000805160206125bb833981519152602052604090206001015490565b6001600160a01b038082166000908152600c6020526040902054829116610772576040516325ec6c1f60e01b815260040160405180910390fd5b600061077d81611302565b6001600160a01b0383166000908152600c6020526040902054600160a01b900460ff16156107fe576001600160a01b0383166000818152600c6020908152604091829020805460ff60a01b1916905590519182527fae02c1bd695006b6d891af37fdeefea45a10ebcc17071e3471787db4f177288591015b60405180910390a15b505050565b61080c82610716565b61081581611302565b61081f838361130c565b50505050565b6001600160a01b038116331461084e5760405163334bd91960e11b815260040160405180910390fd5b6107fe82826113b8565b610860611434565b336000818152600e602052604090205465ffffffffffff16156108965760405163574a3d2d60e11b815260040160405180910390fd5b6000846040516020016108a9919061204f565b60408051601f1981840301815291815281516020928301206000818152600d90935291205490915060ff16156108f2576040516326375da160e11b815260040160405180910390fd5b6000818152600d60209081526040808320805460ff19166001179055338352600c9091529020546001600160a01b031615610940576040516342ee68b560e01b815260040160405180910390fd5b60095465ffffffffffff9081169084161080610970575060095465ffffffffffff600160301b9091048116908416115b1561098e5760405163dc81db8560e01b815260040160405180910390fd5b6109973361146c565b6008546002546109b6916001600160a01b0390911690339030906114f7565b60006109c13361155e565b50506040805160a0810182526001600160a01b038084168252600060208084018281528486018d81528c8516606087015265ffffffffffff8c166080870152338452600c909252949091208351815495511515600160a01b026001600160a81b03199096169316929092179390931781559151929350916001820190610a4790826120bb565b5060608201516002909101805460809093015165ffffffffffff16600160a01b026001600160d01b03199093166001600160a01b03909216919091179190911790556040517f4f8ede3304cdb9db4ee7e3e1c7e1a12d293f88c81465a874846a844ff1de65b890610ac1903390899089908990879061217a565b60405180910390a15050506107fe60016000805160206125db83398151915255565b6001600160a01b038082166000908152600c6020526040902054829116610b1d576040516325ec6c1f60e01b815260040160405180910390fd5b6000610b2881611302565b6001600160a01b0383166000908152600c6020526040902054600160a01b900460ff166107fe576001600160a01b0383166000818152600c6020908152604091829020805460ff60a01b1916600160a01b17905590519182527fc5437eb8dd091f69800961953f2bb0bc16ae1ff2d3e52caa96796db65f8271da91016107f5565b6000610bb481611302565b508051600080546001600160a01b039283166001600160a01b0319918216179091556020808401516001805460408088015165ffffffffffff16600160a01b026001600160d01b0319909216938716939093171790556060808601516002805491871691861691909117905560809586015180516003805495830151948301516001600160401b03908116600160801b0267ffffffffffffffff60801b19968216600160401b026fffffffffffffffffffffffffffffffff1990981691909316179590951793909316929092179092559081015160048054918516918416919091179055928301516005805491841691831691909117905560a09092015160068054919092169216919091179055565b60009182526000805160206125bb833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b0316600081158015610d415750825b90506000826001600160401b03166001148015610d5d5750303b155b905081158015610d6b575080155b15610d895760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610db357845460ff60401b1916600160401b1785555b610dbb61186e565b610dc660003361130c565b50610dcf611878565b610dd888610ba9565b610de1876105b6565b610dea86610575565b8315610e3057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b336000908152600e602052604081205465ffffffffffff169003610e7157604051635e0580eb60e01b815260040160405180910390fd5b336000908152600e602052604090205465ffffffffffff16610e91611888565b65ffffffffffff161015610eb8576040516352d1770f60e01b815260040160405180910390fd5b600854600254610ed5916001600160a01b03909116903390611898565b336000908152600e60205260409020805465ffffffffffff19169055565b6001600160a01b038084166000908152600c60205260408120549091859116610f2f576040516325ec6c1f60e01b815260040160405180910390fd5b6000610f3a81611302565b6001600160a01b038087166000908152600c6020908152604080832054815163b134427160e01b81529151941693849263b134427192600480820193918290030181865afa158015610f90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fb491906121c3565b6040805160208101825260008152905163010d40ab60e11b81529192506001600160a01b0383169163021a815691610ff6918d918d918d918d916004016121e0565b6020604051808303816000875af1158015611015573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110399190612226565b9998505050505050505050565b336000818152600c60205260409020546001600160a01b031661107c576040516325ec6c1f60e01b815260040160405180910390fd5b611084611434565b336000908152600c6020526040902054600160a01b900460ff16156110bb5760405162417c2d60e11b815260040160405180910390fd5b336000908152600c6020526040812080546001600160a81b0319168155906110e66001830182611a04565b5060020180546001600160d01b031916905560075465ffffffffffff1661110b611888565b611115919061223f565b336000818152600e6020908152604091829020805465ffffffffffff191665ffffffffffff9590951694909417909355519081527f6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f910160405180910390a161118b60016000805160206125db83398151915255565b50565b61119782610716565b6111a081611302565b61081f83836113b8565b7f9188644bf0c7a694e572b54fd40005e1230f80a50c59be6fb567a312ab5a1d4d6111d481611302565b6000836000604051806020016040528060008152506040518060200160405280600081525060405160200161120c949392919061226c565b60408051808303601f1901815290829052600a546348a78da760e01b83526001600160a01b038b811660048501528a81166024850152604484018a90526064840187905291935016906348a78da790608401600060405180830381600087803b15801561127857600080fd5b505af115801561128c573d6000803e3d6000fd5b5050600b5460405163239723ed60e01b81526001600160a01b03909116925063239723ed91506112c6908b908b908a9087906004016122a8565b600060405180830381600087803b1580156112e057600080fd5b505af11580156112f4573d6000803e3d6000fd5b505050505050505050505050565b61118b81336118c9565b60006000805160206125bb8339815191526113278484610cc4565b6113a7576000848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561135d3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061056f565b600091505061056f565b5092915050565b60006000805160206125bb8339815191526113d38484610cc4565b156113a7576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061056f565b6000805160206125db83398151915280546001190161146657604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6000546040516302910f8b60e31b81526001600160a01b038381166004830152909116906314887c5890602401602060405180830381865afa1580156114b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114da91906122e5565b61118b57604051635f2b0a4960e01b815260040160405180910390fd5b6040516001600160a01b03848116602483015283811660448301526064820183905261081f9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b03838183161783525050505061190b565b6040805161016081018252600280546001600160a01b0390811683526004548116602084015260015465ffffffffffff600160a01b9091041683850152600060608085018290526080850182905260a0850182905260055490921660c085015260e0840181905261010084018190526101208401819052610140840181905284518381529182019094528392839290918391816020016020820280368337505060055482519293506001600160a01b03169183915060009061162257611622612307565b60200260200101906001600160a01b031690816001600160a01b031681525050308160018151811061165657611656612307565b6001600160a01b039283166020918202929092018101919091526040805160c0810182526005548416606082018181526006548616608084015260a083018290528252818401869052938a168183015281518083018352600181850190815281528251610100810184526003546001600160401b03168152808501959095528251919490936000939192918301916116f09189910161231d565b60408051808303601f19018152918152908252600354600160401b90046001600160401b03166020808401919091528151929091019161173291879101612405565b60408051808303601f1901815291815290825260016020808401829052600354600160801b90046001600160401b031684840152825187515115158183015283518082039092018252830183526060909301929092529054905163312249f960e21b81529192506001600160a01b03169063c48927e4906117b7908490600401612499565b6060604051808303816000875af11580156117d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117fa919061256d565b919950975095506001600160a01b038816158061181e57506001600160a01b038716155b8061183057506001600160a01b038616155b1561184e576040516307cbdadb60e11b815260040160405180910390fd5b50505050509193909250565b60016000805160206125db83398151915255565b61187661197c565b565b61188061197c565b6118766119c5565b6000611893426119cd565b905090565b6040516001600160a01b038381166024830152604482018390526107fe91859182169063a9059cbb9060640161152c565b6118d38282610cc4565b6119075760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044015b60405180910390fd5b5050565b600080602060008451602086016000885af18061192e576040513d6000823e3d81fd5b50506000513d91508115611946578060011415611953565b6001600160a01b0384163b155b1561081f57604051635274afe760e01b81526001600160a01b03851660048201526024016118fe565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661187657604051631afcd79f60e31b815260040160405180910390fd5b61185a61197c565b600065ffffffffffff821115611a00576040516306dfcc6560e41b815260306004820152602481018390526044016118fe565b5090565b508054611a1090612015565b6000825580601f10611a20575050565b601f01602090049060005260206000209081019061118b91905b80821115611a005760008155600101611a3a565b600060208284031215611a6057600080fd5b81356001600160e01b031981168114611a7857600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715611ab757611ab7611a7f565b60405290565b60405160c081016001600160401b0381118282101715611ab757611ab7611a7f565b604051601f8201601f191681016001600160401b0381118282101715611b0757611b07611a7f565b604052919050565b6001600160a01b038116811461118b57600080fd5b8035611b2f81611b0f565b919050565b600060408284031215611b4657600080fd5b604051604081018181106001600160401b0382111715611b6857611b68611a7f565b6040529050808235611b7981611b0f565b81526020830135611b8981611b0f565b6020919091015292915050565b600060408284031215611ba857600080fd5b611a788383611b34565b803565ffffffffffff81168114611b2f57600080fd5b600060a08284031215611bda57600080fd5b611be2611a95565b9050611bed82611bb2565b815260208201356020820152611c0560408301611bb2565b6040820152611c1660608301611bb2565b6060820152611c2760808301611bb2565b608082015292915050565b600060a08284031215611c4457600080fd5b611a788383611bc8565b600060208284031215611c6057600080fd5b8135611a7881611b0f565b60005b83811015611c86578181015183820152602001611c6e565b50506000910152565b60008151808452611ca7816020860160208601611c6b565b601f01601f19169290920160200192915050565b600060018060a01b038088168352861515602084015260a06040840152611ce560a0840187611c8f565b941660608301525065ffffffffffff919091166080909101529392505050565b600060208284031215611d1757600080fd5b5035919050565b60008060408385031215611d3157600080fd5b823591506020830135611d4381611b0f565b809150509250929050565b600080600060608486031215611d6357600080fd5b83356001600160401b0380821115611d7a57600080fd5b818601915086601f830112611d8e57600080fd5b8135602082821115611da257611da2611a7f565b611db4601f8301601f19168201611adf565b92508183528881838601011115611dca57600080fd5b81818501828501376000818385010152829650611de8818901611b24565b955050505050611dfa60408501611bb2565b90509250925092565b80356001600160401b0381168114611b2f57600080fd5b6000818303610140811215611e2e57600080fd5b611e36611a95565b91508235611e4381611b0f565b82526020830135611e5381611b0f565b6020830152611e6460408401611bb2565b60408301526060830135611e7781611b0f565b606083015260c0607f1982011215611e8e57600080fd5b50611e97611abd565b611ea360808401611e03565b8152611eb160a08401611e03565b6020820152611ec260c08401611e03565b604082015260e0830135611ed581611b0f565b6060820152610100830135611ee981611b0f565b6080820152610120830135611efd81611b0f565b60a0820152608082015292915050565b60006101408284031215611f2057600080fd5b611a788383611e1a565b60008060006102208486031215611f4057600080fd5b611f4a8585611e1a565b9250611f5a856101408601611bc8565b9150611dfa856101e08601611b34565b60008060008060808587031215611f8057600080fd5b843593506020850135611f9281611b0f565b925060408501359150611fa760608601611bb2565b905092959194509250565b60008060008060008060c08789031215611fcb57600080fd5b8635611fd681611b0f565b95506020870135611fe681611b0f565b9450604087013593506060870135925061200260808801611bb2565b915060a087013590509295509295509295565b600181811c9082168061202957607f821691505b60208210810361204957634e487b7160e01b600052602260045260246000fd5b50919050565b60008251612061818460208701611c6b565b9190910192915050565b601f8211156107fe576000816000526020600020601f850160051c810160208610156120945750805b601f850160051c820191505b818110156120b3578281556001016120a0565b505050505050565b81516001600160401b038111156120d4576120d4611a7f565b6120e8816120e28454612015565b8461206b565b602080601f83116001811461211d57600084156121055750858301515b600019600386901b1c1916600185901b1785556120b3565b600085815260208120601f198616915b8281101561214c5788860151825594840194600190910190840161212d565b508582101561216a5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600060018060a01b03808816835260a0602084015261219c60a0840188611c8f565b958116604084015265ffffffffffff94909416606083015250911660809091015292915050565b6000602082840312156121d557600080fd5b8151611a7881611b0f565b85815260018060a01b038516602082015283604082015265ffffffffffff8316606082015260a06080820152600061221b60a0830184611c8f565b979650505050505050565b60006020828403121561223857600080fd5b5051919050565b65ffffffffffff8181168382160190808211156113b157634e487b7160e01b600052601160045260246000fd5b65ffffffffffff8516815260ff841660208201526080604082015260006122966080830185611c8f565b828103606084015261221b8185611c8f565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906122db90830184611c8f565b9695505050505050565b6000602082840312156122f757600080fd5b81518015158114611a7857600080fd5b634e487b7160e01b600052603260045260246000fd5b81516001600160a01b031681526101608101602083015161234960208401826001600160a01b03169052565b506040830151612363604084018265ffffffffffff169052565b506060830151612377606084018215159052565b50608083015161238b608084018215159052565b5060a083015160a083015260c08301516123b060c08401826001600160a01b03169052565b5060e08301516123cb60e08401826001600160a01b03169052565b50610100838101516001600160a01b0390811691840191909152610120808501518216908401526101409384015116929091019190915290565b6020808252825180516001600160a01b039081168484015281830151811660408086019190915290910151811660608401528382015160a06080850152805160c085018190526000939291830191849160e08701905b8084101561247d5784518316825293850193600193909301929085019061245b565b5060408801516001600160a01b03811660a0890152945061221b565b602081526124b36020820183516001600160401b03169052565b600060208301516124cf60408401826001600160a01b03169052565b5060408301516101008060608501526124ec610120850183611c8f565b9150606085015161250860808601826001600160401b03169052565b506080850151601f19808685030160a08701526125258483611c8f565b935060a0870151915061253c60c087018315159052565b60c08701516001600160401b03811660e0880152915060e08701519150808685030183870152506122db8382611c8f565b60008060006060848603121561258257600080fd5b835161258d81611b0f565b602085015190935061259e81611b0f565b60408501519092506125af81611b0f565b80915050925092509256fe02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220d1d5eb7c100fa28805c0e6d7498aa798201bf74a6967a8a2464cf8d9a8f1339d64736f6c63430008190033",
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

// REWARDUPDATERROLE is a free data retrieval call binding the contract method 0x1a69d9c8.
//
// Solidity: function REWARD_UPDATER_ROLE() view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) REWARDUPDATERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "REWARD_UPDATER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDUPDATERROLE is a free data retrieval call binding the contract method 0x1a69d9c8.
//
// Solidity: function REWARD_UPDATER_ROLE() view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) REWARDUPDATERROLE() ([32]byte, error) {
	return _MantaStakingMiddleware.Contract.REWARDUPDATERROLE(&_MantaStakingMiddleware.CallOpts)
}

// REWARDUPDATERROLE is a free data retrieval call binding the contract method 0x1a69d9c8.
//
// Solidity: function REWARD_UPDATER_ROLE() view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) REWARDUPDATERROLE() ([32]byte, error) {
	return _MantaStakingMiddleware.Contract.REWARDUPDATERROLE(&_MantaStakingMiddleware.CallOpts)
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
// Solidity: function operatorSettings() view returns(uint48 unregisterTokenUnlockWindow, uint256 requiredOperatorStake, uint48 minOperatorCommission, uint48 maxOperatorCommission, uint48 maxValidOperator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) OperatorSettings(opts *bind.CallOpts) (struct {
	UnregisterTokenUnlockWindow *big.Int
	RequiredOperatorStake       *big.Int
	MinOperatorCommission       *big.Int
	MaxOperatorCommission       *big.Int
	MaxValidOperator            *big.Int
}, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "operatorSettings")

	outstruct := new(struct {
		UnregisterTokenUnlockWindow *big.Int
		RequiredOperatorStake       *big.Int
		MinOperatorCommission       *big.Int
		MaxOperatorCommission       *big.Int
		MaxValidOperator            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UnregisterTokenUnlockWindow = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RequiredOperatorStake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MinOperatorCommission = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxOperatorCommission = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxValidOperator = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OperatorSettings is a free data retrieval call binding the contract method 0x72aef6bf.
//
// Solidity: function operatorSettings() view returns(uint48 unregisterTokenUnlockWindow, uint256 requiredOperatorStake, uint48 minOperatorCommission, uint48 maxOperatorCommission, uint48 maxValidOperator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) OperatorSettings() (struct {
	UnregisterTokenUnlockWindow *big.Int
	RequiredOperatorStake       *big.Int
	MinOperatorCommission       *big.Int
	MaxOperatorCommission       *big.Int
	MaxValidOperator            *big.Int
}, error) {
	return _MantaStakingMiddleware.Contract.OperatorSettings(&_MantaStakingMiddleware.CallOpts)
}

// OperatorSettings is a free data retrieval call binding the contract method 0x72aef6bf.
//
// Solidity: function operatorSettings() view returns(uint48 unregisterTokenUnlockWindow, uint256 requiredOperatorStake, uint48 minOperatorCommission, uint48 maxOperatorCommission, uint48 maxValidOperator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) OperatorSettings() (struct {
	UnregisterTokenUnlockWindow *big.Int
	RequiredOperatorStake       *big.Int
	MinOperatorCommission       *big.Int
	MaxOperatorCommission       *big.Int
	MaxValidOperator            *big.Int
}, error) {
	return _MantaStakingMiddleware.Contract.OperatorSettings(&_MantaStakingMiddleware.CallOpts)
}

// OperatorTokenUnlockTimestamps is a free data retrieval call binding the contract method 0xf3f2c422.
//
// Solidity: function operatorTokenUnlockTimestamps(address ) view returns(uint48)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) OperatorTokenUnlockTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "operatorTokenUnlockTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorTokenUnlockTimestamps is a free data retrieval call binding the contract method 0xf3f2c422.
//
// Solidity: function operatorTokenUnlockTimestamps(address ) view returns(uint48)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) OperatorTokenUnlockTimestamps(arg0 common.Address) (*big.Int, error) {
	return _MantaStakingMiddleware.Contract.OperatorTokenUnlockTimestamps(&_MantaStakingMiddleware.CallOpts, arg0)
}

// OperatorTokenUnlockTimestamps is a free data retrieval call binding the contract method 0xf3f2c422.
//
// Solidity: function operatorTokenUnlockTimestamps(address ) view returns(uint48)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) OperatorTokenUnlockTimestamps(arg0 common.Address) (*big.Int, error) {
	return _MantaStakingMiddleware.Contract.OperatorTokenUnlockTimestamps(&_MantaStakingMiddleware.CallOpts, arg0)
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

// RewardSettings is a free data retrieval call binding the contract method 0xc15025c6.
//
// Solidity: function rewardSettings() view returns(address operatorRewardDistributor, address stakerRewardDistributor)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) RewardSettings(opts *bind.CallOpts) (struct {
	OperatorRewardDistributor common.Address
	StakerRewardDistributor   common.Address
}, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "rewardSettings")

	outstruct := new(struct {
		OperatorRewardDistributor common.Address
		StakerRewardDistributor   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorRewardDistributor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StakerRewardDistributor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// RewardSettings is a free data retrieval call binding the contract method 0xc15025c6.
//
// Solidity: function rewardSettings() view returns(address operatorRewardDistributor, address stakerRewardDistributor)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) RewardSettings() (struct {
	OperatorRewardDistributor common.Address
	StakerRewardDistributor   common.Address
}, error) {
	return _MantaStakingMiddleware.Contract.RewardSettings(&_MantaStakingMiddleware.CallOpts)
}

// RewardSettings is a free data retrieval call binding the contract method 0xc15025c6.
//
// Solidity: function rewardSettings() view returns(address operatorRewardDistributor, address stakerRewardDistributor)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) RewardSettings() (struct {
	OperatorRewardDistributor common.Address
	StakerRewardDistributor   common.Address
}, error) {
	return _MantaStakingMiddleware.Contract.RewardSettings(&_MantaStakingMiddleware.CallOpts)
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

// ClaimUnlockedToken is a paid mutator transaction binding the contract method 0xa0f98db6.
//
// Solidity: function claimUnlockedToken() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) ClaimUnlockedToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "claimUnlockedToken")
}

// ClaimUnlockedToken is a paid mutator transaction binding the contract method 0xa0f98db6.
//
// Solidity: function claimUnlockedToken() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) ClaimUnlockedToken() (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.ClaimUnlockedToken(&_MantaStakingMiddleware.TransactOpts)
}

// ClaimUnlockedToken is a paid mutator transaction binding the contract method 0xa0f98db6.
//
// Solidity: function claimUnlockedToken() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) ClaimUnlockedToken() (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.ClaimUnlockedToken(&_MantaStakingMiddleware.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0xff4e42d7.
//
// Solidity: function distributeRewards(address network, address token, uint256 operatorRewardAmount, uint256 stakerRewardAmount, uint48 captureTimestamp, bytes32 operatorRewardRoot) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) DistributeRewards(opts *bind.TransactOpts, network common.Address, token common.Address, operatorRewardAmount *big.Int, stakerRewardAmount *big.Int, captureTimestamp *big.Int, operatorRewardRoot [32]byte) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "distributeRewards", network, token, operatorRewardAmount, stakerRewardAmount, captureTimestamp, operatorRewardRoot)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0xff4e42d7.
//
// Solidity: function distributeRewards(address network, address token, uint256 operatorRewardAmount, uint256 stakerRewardAmount, uint48 captureTimestamp, bytes32 operatorRewardRoot) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) DistributeRewards(network common.Address, token common.Address, operatorRewardAmount *big.Int, stakerRewardAmount *big.Int, captureTimestamp *big.Int, operatorRewardRoot [32]byte) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.DistributeRewards(&_MantaStakingMiddleware.TransactOpts, network, token, operatorRewardAmount, stakerRewardAmount, captureTimestamp, operatorRewardRoot)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0xff4e42d7.
//
// Solidity: function distributeRewards(address network, address token, uint256 operatorRewardAmount, uint256 stakerRewardAmount, uint48 captureTimestamp, bytes32 operatorRewardRoot) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) DistributeRewards(network common.Address, token common.Address, operatorRewardAmount *big.Int, stakerRewardAmount *big.Int, captureTimestamp *big.Int, operatorRewardRoot [32]byte) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.DistributeRewards(&_MantaStakingMiddleware.TransactOpts, network, token, operatorRewardAmount, stakerRewardAmount, captureTimestamp, operatorRewardRoot)
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

// Initialize is a paid mutator transaction binding the contract method 0x989bfec5.
//
// Solidity: function initialize((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) symbioticVaultSettings_, (uint48,uint256,uint48,uint48,uint48) operatorSettings_, (address,address) rewardSettings_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) Initialize(opts *bind.TransactOpts, symbioticVaultSettings_ SymbioticVaultSettings, operatorSettings_ OperatorSettings, rewardSettings_ RewardSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "initialize", symbioticVaultSettings_, operatorSettings_, rewardSettings_)
}

// Initialize is a paid mutator transaction binding the contract method 0x989bfec5.
//
// Solidity: function initialize((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) symbioticVaultSettings_, (uint48,uint256,uint48,uint48,uint48) operatorSettings_, (address,address) rewardSettings_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) Initialize(symbioticVaultSettings_ SymbioticVaultSettings, operatorSettings_ OperatorSettings, rewardSettings_ RewardSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Initialize(&_MantaStakingMiddleware.TransactOpts, symbioticVaultSettings_, operatorSettings_, rewardSettings_)
}

// Initialize is a paid mutator transaction binding the contract method 0x989bfec5.
//
// Solidity: function initialize((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) symbioticVaultSettings_, (uint48,uint256,uint48,uint48,uint48) operatorSettings_, (address,address) rewardSettings_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) Initialize(symbioticVaultSettings_ SymbioticVaultSettings, operatorSettings_ OperatorSettings, rewardSettings_ RewardSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Initialize(&_MantaStakingMiddleware.TransactOpts, symbioticVaultSettings_, operatorSettings_, rewardSettings_)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) PauseOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "pauseOperator", operator)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) PauseOperator(operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.PauseOperator(&_MantaStakingMiddleware.TransactOpts, operator)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) PauseOperator(operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.PauseOperator(&_MantaStakingMiddleware.TransactOpts, operator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x60ba757d.
//
// Solidity: function registerOperator(string operatorName, address rewardAddress, uint48 commission) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) RegisterOperator(opts *bind.TransactOpts, operatorName string, rewardAddress common.Address, commission *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "registerOperator", operatorName, rewardAddress, commission)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x60ba757d.
//
// Solidity: function registerOperator(string operatorName, address rewardAddress, uint48 commission) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) RegisterOperator(operatorName string, rewardAddress common.Address, commission *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RegisterOperator(&_MantaStakingMiddleware.TransactOpts, operatorName, rewardAddress, commission)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x60ba757d.
//
// Solidity: function registerOperator(string operatorName, address rewardAddress, uint48 commission) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) RegisterOperator(operatorName string, rewardAddress common.Address, commission *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RegisterOperator(&_MantaStakingMiddleware.TransactOpts, operatorName, rewardAddress, commission)
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

// Slash is a paid mutator transaction binding the contract method 0xa314db16.
//
// Solidity: function slash(bytes32 subnetwork, address operator, uint256 amount, uint48 captureTimestamp) returns(uint256 slashedAmount)
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) Slash(opts *bind.TransactOpts, subnetwork [32]byte, operator common.Address, amount *big.Int, captureTimestamp *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "slash", subnetwork, operator, amount, captureTimestamp)
}

// Slash is a paid mutator transaction binding the contract method 0xa314db16.
//
// Solidity: function slash(bytes32 subnetwork, address operator, uint256 amount, uint48 captureTimestamp) returns(uint256 slashedAmount)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) Slash(subnetwork [32]byte, operator common.Address, amount *big.Int, captureTimestamp *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Slash(&_MantaStakingMiddleware.TransactOpts, subnetwork, operator, amount, captureTimestamp)
}

// Slash is a paid mutator transaction binding the contract method 0xa314db16.
//
// Solidity: function slash(bytes32 subnetwork, address operator, uint256 amount, uint48 captureTimestamp) returns(uint256 slashedAmount)
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) Slash(subnetwork [32]byte, operator common.Address, amount *big.Int, captureTimestamp *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Slash(&_MantaStakingMiddleware.TransactOpts, subnetwork, operator, amount, captureTimestamp)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) UnpauseOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "unpauseOperator", operator)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) UnpauseOperator(operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UnpauseOperator(&_MantaStakingMiddleware.TransactOpts, operator)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) UnpauseOperator(operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UnpauseOperator(&_MantaStakingMiddleware.TransactOpts, operator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0xa876b89a.
//
// Solidity: function unregisterOperator() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) UnregisterOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "unregisterOperator")
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0xa876b89a.
//
// Solidity: function unregisterOperator() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) UnregisterOperator() (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UnregisterOperator(&_MantaStakingMiddleware.TransactOpts)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0xa876b89a.
//
// Solidity: function unregisterOperator() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) UnregisterOperator() (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UnregisterOperator(&_MantaStakingMiddleware.TransactOpts)
}

// UpdateOperatorSettings is a paid mutator transaction binding the contract method 0x12f09ad7.
//
// Solidity: function updateOperatorSettings((uint48,uint256,uint48,uint48,uint48) operatorSettings_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) UpdateOperatorSettings(opts *bind.TransactOpts, operatorSettings_ OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "updateOperatorSettings", operatorSettings_)
}

// UpdateOperatorSettings is a paid mutator transaction binding the contract method 0x12f09ad7.
//
// Solidity: function updateOperatorSettings((uint48,uint256,uint48,uint48,uint48) operatorSettings_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) UpdateOperatorSettings(operatorSettings_ OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UpdateOperatorSettings(&_MantaStakingMiddleware.TransactOpts, operatorSettings_)
}

// UpdateOperatorSettings is a paid mutator transaction binding the contract method 0x12f09ad7.
//
// Solidity: function updateOperatorSettings((uint48,uint256,uint48,uint48,uint48) operatorSettings_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) UpdateOperatorSettings(operatorSettings_ OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UpdateOperatorSettings(&_MantaStakingMiddleware.TransactOpts, operatorSettings_)
}

// UpdateRewardSettings is a paid mutator transaction binding the contract method 0x0c761c97.
//
// Solidity: function updateRewardSettings((address,address) rewardSettings_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) UpdateRewardSettings(opts *bind.TransactOpts, rewardSettings_ RewardSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "updateRewardSettings", rewardSettings_)
}

// UpdateRewardSettings is a paid mutator transaction binding the contract method 0x0c761c97.
//
// Solidity: function updateRewardSettings((address,address) rewardSettings_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) UpdateRewardSettings(rewardSettings_ RewardSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UpdateRewardSettings(&_MantaStakingMiddleware.TransactOpts, rewardSettings_)
}

// UpdateRewardSettings is a paid mutator transaction binding the contract method 0x0c761c97.
//
// Solidity: function updateRewardSettings((address,address) rewardSettings_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) UpdateRewardSettings(rewardSettings_ RewardSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UpdateRewardSettings(&_MantaStakingMiddleware.TransactOpts, rewardSettings_)
}

// UpdateSymbioticVaultSettings is a paid mutator transaction binding the contract method 0x7a6efb69.
//
// Solidity: function updateSymbioticVaultSettings((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) symbioticVaultSettings_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) UpdateSymbioticVaultSettings(opts *bind.TransactOpts, symbioticVaultSettings_ SymbioticVaultSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "updateSymbioticVaultSettings", symbioticVaultSettings_)
}

// UpdateSymbioticVaultSettings is a paid mutator transaction binding the contract method 0x7a6efb69.
//
// Solidity: function updateSymbioticVaultSettings((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) symbioticVaultSettings_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) UpdateSymbioticVaultSettings(symbioticVaultSettings_ SymbioticVaultSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UpdateSymbioticVaultSettings(&_MantaStakingMiddleware.TransactOpts, symbioticVaultSettings_)
}

// UpdateSymbioticVaultSettings is a paid mutator transaction binding the contract method 0x7a6efb69.
//
// Solidity: function updateSymbioticVaultSettings((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) symbioticVaultSettings_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) UpdateSymbioticVaultSettings(symbioticVaultSettings_ SymbioticVaultSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UpdateSymbioticVaultSettings(&_MantaStakingMiddleware.TransactOpts, symbioticVaultSettings_)
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
	Operator      common.Address
	OperatorName  string
	RewardAddress common.Address
	Commission    *big.Int
	Vault         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x4f8ede3304cdb9db4ee7e3e1c7e1a12d293f88c81465a874846a844ff1de65b8.
//
// Solidity: event OperatorRegistered(address operator, string operatorName, address rewardAddress, uint48 commission, address vault)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterOperatorRegistered(opts *bind.FilterOpts) (*MantaStakingMiddlewareOperatorRegisteredIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "OperatorRegistered")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareOperatorRegisteredIterator{contract: _MantaStakingMiddleware.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x4f8ede3304cdb9db4ee7e3e1c7e1a12d293f88c81465a874846a844ff1de65b8.
//
// Solidity: event OperatorRegistered(address operator, string operatorName, address rewardAddress, uint48 commission, address vault)
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x4f8ede3304cdb9db4ee7e3e1c7e1a12d293f88c81465a874846a844ff1de65b8.
//
// Solidity: event OperatorRegistered(address operator, string operatorName, address rewardAddress, uint48 commission, address vault)
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
