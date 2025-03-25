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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimUnlockedToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_symbioticVaultSettings\",\"type\":\"tuple\",\"internalType\":\"structSymbioticVaultSettings\",\"components\":[{\"name\":\"operatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultConfiguration\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"stakeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultVaultParams\",\"type\":\"tuple\",\"internalType\":\"structDefaultVaultInitParams\",\"components\":[{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delegatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slasherIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"vaultSlashBurner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultDefaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeforeSlashHook\",\"type\":\"address\",\"internalType\":\"address\"}]}]},{\"name\":\"_operatorSettings\",\"type\":\"tuple\",\"internalType\":\"structOperatorSettings\",\"components\":[{\"name\":\"unregisterTokenUnlockWindow\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"requiredOperatorStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorNameExists\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSettings\",\"inputs\":[],\"outputs\":[{\"name\":\"unregisterTokenUnlockWindow\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"requiredOperatorStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorTokenUnlockTimestamps\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operators\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"operatorName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"commission\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseOperator\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"_operatorPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_operatorName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_commission\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardAddress\",\"inputs\":[{\"name\":\"_rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slash\",\"inputs\":[{\"name\":\"_subnetwork\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_captureTimestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[{\"name\":\"slashedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbioticVaultSettings\",\"inputs\":[],\"outputs\":[{\"name\":\"operatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultConfiguration\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"stakeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultVaultParams\",\"type\":\"tuple\",\"internalType\":\"structDefaultVaultInitParams\",\"components\":[{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delegatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slasherIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"vaultSlashBurner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultDefaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeforeSlashHook\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpauseOperator\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterOperator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorSettings\",\"inputs\":[{\"name\":\"_operatorSettings\",\"type\":\"tuple\",\"internalType\":\"structOperatorSettings\",\"components\":[{\"name\":\"unregisterTokenUnlockWindow\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"requiredOperatorStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxOperatorCommission\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSymbioticVaultSettings\",\"inputs\":[{\"name\":\"_symbioticVaultSettings\",\"type\":\"tuple\",\"internalType\":\"structSymbioticVaultSettings\",\"components\":[{\"name\":\"operatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultConfiguration\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"stakeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultVaultParams\",\"type\":\"tuple\",\"internalType\":\"structDefaultVaultInitParams\",\"components\":[{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delegatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slasherIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"vaultSlashBurner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultDefaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeforeSlashHook\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorPaused\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorPublicKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"operatorName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"commission\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnpaused\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardAddressSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b612659806100d96000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806372f9adab116100b8578063a204fe551161007c578063a204fe5514610401578063a217fddf14610414578063a314db161461041c578063a876b89a1461042f578063d547741f14610437578063f3f2c4221461044a57600080fd5b806372f9adab146103ad5780637a6efb69146103c0578063829126b8146103d357806391d14854146103e6578063a0f98db6146103f957600080fd5b80632e5aaf331161010a5780632e5aaf33146102f15780632e5ea91d146103065780632f2ff15d1461031957806336568abe1461032c5780635e00e6791461033f57806372aef6bf1461035257600080fd5b806301ffc9a71461014757806313e7c9d81461016f5780631b72a9ff146101935780631d6a65d7146102ad578063248a9ca3146102d0575b600080fd5b61015a610155366004611b3d565b610489565b60405190151581526020015b60405180910390f35b61018261017d366004611b93565b6104c0565b604051610166959493929190611c00565b6000546001546002546040805160c0810182526003546001600160401b038082168352600160401b820481166020840152600160801b90910416918101919091526004546001600160a01b03908116606083015260055481166080830152600654811660a083015261021d9481169381811693600160a01b90910465ffffffffffff169291169085565b604080516001600160a01b03968716815294861660208087019190915265ffffffffffff9094168582015291851660608086019190915281516001600160401b0390811660808088019190915294830151811660a080880191909152938301511660c0860152810151851660e0850152918201518416610100840152015190911661012082015261014001610166565b61015a6102bb366004611c4a565b600b6020526000908152604090205460ff1681565b6102e36102de366004611c4a565b61059c565b604051908152602001610166565b6103046102ff366004611b93565b6105be565b005b610304610314366004611e5b565b610691565b610304610327366004611e92565b6107c5565b61030461033a366004611e92565b6107e7565b61030461034d366004611b93565b61081a565b60075460085460095461037b9265ffffffffffff908116929180821691600160301b9091041684565b6040805165ffffffffffff95861681526020810194909452918416918301919091529091166060820152608001610166565b6103046103bb366004611b93565b6108ba565b6103046103ce366004611ec2565b61097f565b6103046103e1366004611f54565b610a9a565b61015a6103f4366004611e92565b610ea4565b610304610edc565b61030461040f366004612003565b61101b565b6102e3600081565b6102e361042a36600461201f565b611081565b6103046111d3565b610304610445366004611e92565b611326565b610472610458366004611b93565b600c6020526000908152604090205465ffffffffffff1681565b60405165ffffffffffff9091168152602001610166565b60006001600160e01b03198216637965db0b60e01b14806104ba57506301ffc9a760e01b6001600160e01b03198316145b92915050565b600a60205260009081526040902080546001820180546001600160a01b03831693600160a01b90930460ff169291906104f89061205c565b80601f01602080910402602001604051908101604052809291908181526020018280546105249061205c565b80156105715780601f1061054657610100808354040283529160200191610571565b820191906000526020600020905b81548152906001019060200180831161055457829003601f168201915b505050600290930154919250506001600160a01b0381169065ffffffffffff600160a01b9091041685565b6000908152600080516020612604833981519152602052604090206001015490565b6001600160a01b038082166000908152600a60205260409020548291166106005760405162461bcd60e51b81526004016105f790612096565b60405180910390fd5b600061060b81611342565b6001600160a01b0383166000908152600a6020526040902054600160a01b900460ff161561068c576001600160a01b0383166000818152600a6020908152604091829020805460ff60a01b1916905590519182527fae02c1bd695006b6d891af37fdeefea45a10ebcc17071e3471787db4f177288591015b60405180910390a15b505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156106d65750825b90506000826001600160401b031660011480156106f25750303b155b905081158015610700575080155b1561071e5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561074857845460ff60401b1916600160401b1785555b61075061134f565b61075b600033611359565b50610764611405565b61076d8761097f565b6107768661101b565b83156107bc57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b6107ce8261059c565b6107d781611342565b6107e18383611359565b50505050565b6001600160a01b03811633146108105760405163334bd91960e11b815260040160405180910390fd5b61068c8282611415565b336000818152600a60205260409020546001600160a01b031661084f5760405162461bcd60e51b81526004016105f790612096565b336000818152600a602090815260409182902060020180546001600160a01b0319166001600160a01b0387169081179091558251938452908301527f490e4e7668b9ec6e5180b1fb3f783a4f81efbb160c2691224937832c195b64ce91015b60405180910390a15050565b6001600160a01b038082166000908152600a60205260409020548291166108f35760405162461bcd60e51b81526004016105f790612096565b60006108fe81611342565b6001600160a01b0383166000908152600a6020526040902054600160a01b900460ff1661068c576001600160a01b0383166000818152600a6020908152604091829020805460ff60a01b1916600160a01b17905590519182527fc5437eb8dd091f69800961953f2bb0bc16ae1ff2d3e52caa96796db65f8271da9101610683565b600061098a81611342565b508051600080546001600160a01b039283166001600160a01b0319918216179091556020808401516001805460408088015165ffffffffffff16600160a01b026001600160d01b0319909216938716939093171790556060808601516002805491871691861691909117905560809586015180516003805495830151948301516001600160401b03908116600160801b0267ffffffffffffffff60801b19968216600160401b026fffffffffffffffffffffffffffffffff1990981691909316179590951793909316929092179092559081015160048054918516918416919091179055928301516005805491841691831691909117905560a09092015160068054919092169216919091179055565b610aa2611491565b336000818152600c602052604090205465ffffffffffff1615610b105760405162461bcd60e51b815260206004820152603260248201526000805160206125e48339815191526044820152712069732070656e64696e6720756e6c6f636b60701b60648201526084016105f7565b845160208601206001600160a01b0381163314610b8b5760405162461bcd60e51b815260206004820152603360248201527f4d616e74615374616b696e674d6964646c65776172653a20696e76616c6964206044820152726f70657261746f72207075626c6963206b657960681b60648201526084016105f7565b600085604051602001610b9e91906120d3565b60408051601f1981840301815291815281516020928301206000818152600b90935291205490915060ff1615610c215760405162461bcd60e51b815260206004820152603460248201526000805160206125e4833981519152604482015273206e616d6520616c72656164792065786973747360601b60648201526084016105f7565b6000818152600b60209081526040808320805460ff19166001179055338352600a9091529020546001600160a01b031615610ca85760405162461bcd60e51b815260206004820152603360248201526000805160206125e483398151915260448201527208185b1c9958591e481c9959da5cdd195c9959606a1b60648201526084016105f7565b60095465ffffffffffff90811690851610801590610cdb575060095465ffffffffffff600160301b909104811690851611155b610d3a5760405162461bcd60e51b815260206004820152602a60248201527f4d616e74615374616b696e674d6964646c65776172653a20696e76616c69642060448201526931b7b6b6b4b9b9b4b7b760b11b60648201526084016105f7565b610d43336114db565b600854600254610d62916001600160a01b0390911690339030906115a9565b6000610d6d33611610565b50506040805160a0810182526001600160a01b038084168252600060208084018281528486018e81528d8516606087015265ffffffffffff8d166080870152338452600a909252949091208351815495511515600160a01b026001600160a81b03199096169316929092179390931781559151929350916001820190610df3908261213f565b5060608201516002909101805460809093015165ffffffffffff16600160a01b026001600160d01b03199093166001600160a01b03909216919091179190911790556040517f82ef2e4bdc58c22c96126c5d61bc476199209d94dc13d6cca98424e7264ee16f90610e6f9033908b908b908b908b9088906121fe565b60405180910390a1505050506107e160017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6000918252600080516020612604833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b336000908152600c602052604081205465ffffffffffff169003610f4f5760405162461bcd60e51b815260206004820152603660248201526000805160206125e483398151915260448201527520686173206e6f2070656e64696e6720756e6c6f636b60501b60648201526084016105f7565b336000908152600c602052604090205465ffffffffffff16610f6f61197c565b65ffffffffffff161015610fe05760405162461bcd60e51b815260206004820152603260248201527f4d616e74615374616b696e674d6964646c65776172653a2063616e6e6f7420636044820152716c61696d206265666f726520756e6c6f636b60701b60648201526084016105f7565b600854600254610ffd916001600160a01b0390911690339061198c565b336000908152600c60205260409020805465ffffffffffff19169055565b600061102681611342565b5080516007805465ffffffffffff191665ffffffffffff92831617905560208201516008556040820151600980546060909401519183166bffffffffffffffffffffffff1990941693909317600160301b9190921602179055565b6001600160a01b038084166000908152600a602052604081205490918591166110bc5760405162461bcd60e51b81526004016105f790612096565b60006110c781611342565b6001600160a01b038087166000908152600a6020908152604080832054815163b134427160e01b81529151941693849263b134427192600480820193918290030181865afa15801561111d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611141919061225b565b6040805160208101825260008152905163010d40ab60e11b81529192506001600160a01b0383169163021a815691611183918d918d918d918d91600401612278565b6020604051808303816000875af11580156111a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111c691906122be565b9998505050505050505050565b336000818152600a60205260409020546001600160a01b03166112085760405162461bcd60e51b81526004016105f790612096565b336000818152600a6020526040902054600160a01b900460ff16156112705760405162461bcd60e51b815260206004820152602a60248201526000805160206125e4833981519152604482015269081a5cc81c185d5cd95960b21b60648201526084016105f7565b336000908152600a6020526040812080546001600160a81b03191681559061129b6001830182611af3565b5060020180546001600160d01b031916905560075465ffffffffffff166112c061197c565b6112ca91906122d7565b336000818152600c6020908152604091829020805465ffffffffffff191665ffffffffffff9590951694909417909355519081527f6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f91016108ae565b61132f8261059c565b61133881611342565b6107e18383611415565b61134c81336119bd565b50565b6113576119fa565b565b60006000805160206126048339815191526113748484610ea4565b6113f4576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556113aa3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506104ba565b60009150506104ba565b5092915050565b61140d6119fa565b611357611a43565b60006000805160206126048339815191526114308484610ea4565b156113f4576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506104ba565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f008054600119016114d557604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6000546040516302910f8b60e31b81526001600160a01b038381166004830152909116906314887c5890602401602060405180830381865afa158015611525573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115499190612304565b61134c5760405162461bcd60e51b815260206004820152603c60248201526000805160206125e483398151915260448201527f206e6f74207265676973746572656420746f2073796d62696f7469630000000060648201526084016105f7565b6040516001600160a01b0384811660248301528381166044830152606482018390526107e19186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050611a4b565b6040805161016081018252600280546001600160a01b0390811683526004548116602084015260015465ffffffffffff600160a01b9091041683850152600060608085018290526080850182905260a0850182905260055490921660c085015260e0840181905261010084018190526101208401819052610140840181905284518381529182019094528392839290918391816020016020820280368337505060055482519293506001600160a01b0316918391506000906116d4576116d4612326565b60200260200101906001600160a01b031690816001600160a01b031681525050308160018151811061170857611708612326565b6001600160a01b039283166020918202929092018101919091526040805160c0810182526005548416606082018181526006548616608084015260a083018290528252818401869052938a168183015281518083018352600181850190815281528251610100810184526003546001600160401b03168152808501959095528251919490936000939192918301916117a29189910161233c565b60408051808303601f19018152918152908252600354600160401b90046001600160401b0316602080840191909152815192909101916117e491879101612424565b60408051808303601f1901815291815290825260016020808401829052600354600160801b90046001600160401b031684840152825187515115158183015283518082039092018252830183526060909301929092529054905163312249f960e21b81529192506001600160a01b03169063c48927e4906118699084906004016124b8565b6060604051808303816000875af1158015611888573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118ac9190612596565b919950975095506001600160a01b038816158015906118d357506001600160a01b03871615155b80156118e757506001600160a01b03861615155b61194a5760405162461bcd60e51b815260206004820152602e60248201527f4d616e74615374616b696e674d6964646c65776172653a206661696c6564207460448201526d1bc818dc99585d19481d985d5b1d60921b60648201526084016105f7565b50505050509193909250565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b600061198742611abc565b905090565b6040516001600160a01b0383811660248301526044820183905261068c91859182169063a9059cbb906064016115de565b6119c78282610ea4565b6119f65760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016105f7565b5050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661135757604051631afcd79f60e31b815260040160405180910390fd5b6119566119fa565b600080602060008451602086016000885af180611a6e576040513d6000823e3d81fd5b50506000513d91508115611a86578060011415611a93565b6001600160a01b0384163b155b156107e157604051635274afe760e01b81526001600160a01b03851660048201526024016105f7565b600065ffffffffffff821115611aef576040516306dfcc6560e41b815260306004820152602481018390526044016105f7565b5090565b508054611aff9061205c565b6000825580601f10611b0f575050565b601f01602090049060005260206000209081019061134c91905b80821115611aef5760008155600101611b29565b600060208284031215611b4f57600080fd5b81356001600160e01b031981168114611b6757600080fd5b9392505050565b6001600160a01b038116811461134c57600080fd5b8035611b8e81611b6e565b919050565b600060208284031215611ba557600080fd5b8135611b6781611b6e565b60005b83811015611bcb578181015183820152602001611bb3565b50506000910152565b60008151808452611bec816020860160208601611bb0565b601f01601f19169290920160200192915050565b600060018060a01b038088168352861515602084015260a06040840152611c2a60a0840187611bd4565b941660608301525065ffffffffffff919091166080909101529392505050565b600060208284031215611c5c57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715611c9b57611c9b611c63565b60405290565b60405160c081016001600160401b0381118282101715611c9b57611c9b611c63565b803565ffffffffffff81168114611b8e57600080fd5b80356001600160401b0381168114611b8e57600080fd5b6000818303610140811215611d0457600080fd5b611d0c611c79565b91508235611d1981611b6e565b82526020830135611d2981611b6e565b6020830152611d3a60408401611cc3565b60408301526060830135611d4d81611b6e565b606083015260c0607f1982011215611d6457600080fd5b50611d6d611ca1565b611d7960808401611cd9565b8152611d8760a08401611cd9565b6020820152611d9860c08401611cd9565b604082015260e0830135611dab81611b6e565b6060820152610100830135611dbf81611b6e565b6080820152610120830135611dd381611b6e565b60a0820152608082015292915050565b600060808284031215611df557600080fd5b604051608081018181106001600160401b0382111715611e1757611e17611c63565b604052905080611e2683611cc3565b815260208301356020820152611e3e60408401611cc3565b6040820152611e4f60608401611cc3565b60608201525092915050565b6000806101c08385031215611e6f57600080fd5b611e798484611cf0565b9150611e89846101408501611de3565b90509250929050565b60008060408385031215611ea557600080fd5b823591506020830135611eb781611b6e565b809150509250929050565b60006101408284031215611ed557600080fd5b611b678383611cf0565b60006001600160401b0380841115611ef957611ef9611c63565b604051601f8501601f19908116603f01168101908282118183101715611f2157611f21611c63565b81604052809350858152868686011115611f3a57600080fd5b858560208301376000602087830101525050509392505050565b60008060008060808587031215611f6a57600080fd5b84356001600160401b0380821115611f8157600080fd5b818701915087601f830112611f9557600080fd5b611fa488833560208501611edf565b95506020870135915080821115611fba57600080fd5b508501601f81018713611fcc57600080fd5b611fdb87823560208401611edf565b935050611fea60408601611b83565b9150611ff860608601611cc3565b905092959194509250565b60006080828403121561201557600080fd5b611b678383611de3565b6000806000806080858703121561203557600080fd5b84359350602085013561204781611b6e565b925060408501359150611ff860608601611cc3565b600181811c9082168061207057607f821691505b60208210810361209057634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602f908201526000805160206125e483398151915260408201526e081b9bdd081c9959da5cdd195c9959608a1b606082015260800190565b600082516120e5818460208701611bb0565b9190910192915050565b601f82111561068c576000816000526020600020601f850160051c810160208610156121185750805b601f850160051c820191505b8181101561213757828155600101612124565b505050505050565b81516001600160401b0381111561215857612158611c63565b61216c81612166845461205c565b846120ef565b602080601f8311600181146121a157600084156121895750858301515b600019600386901b1c1916600185901b178555612137565b600085815260208120601f198616915b828110156121d0578886015182559484019460019091019084016121b1565b50858210156121ee5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600060018060a01b03808916835260c0602084015261222060c0840189611bd4565b83810360408501526122328189611bd4565b96821660608501525065ffffffffffff94909416608083015250911660a0909101529392505050565b60006020828403121561226d57600080fd5b8151611b6781611b6e565b85815260018060a01b038516602082015283604082015265ffffffffffff8316606082015260a0608082015260006122b360a0830184611bd4565b979650505050505050565b6000602082840312156122d057600080fd5b5051919050565b65ffffffffffff8181168382160190808211156113fe57634e487b7160e01b600052601160045260246000fd5b60006020828403121561231657600080fd5b81518015158114611b6757600080fd5b634e487b7160e01b600052603260045260246000fd5b81516001600160a01b031681526101608101602083015161236860208401826001600160a01b03169052565b506040830151612382604084018265ffffffffffff169052565b506060830151612396606084018215159052565b5060808301516123aa608084018215159052565b5060a083015160a083015260c08301516123cf60c08401826001600160a01b03169052565b5060e08301516123ea60e08401826001600160a01b03169052565b50610100838101516001600160a01b0390811691840191909152610120808501518216908401526101409384015116929091019190915290565b6020808252825180516001600160a01b039081168484015281830151811660408086019190915290910151811660608401528382015160a06080850152805160c085018190526000939291830191849160e08701905b8084101561249c5784518316825293850193600193909301929085019061247a565b5060408801516001600160a01b03811660a089015294506122b3565b602081526124d26020820183516001600160401b03169052565b600060208301516124ee60408401826001600160a01b03169052565b50604083015161010080606085015261250b610120850183611bd4565b9150606085015161252760808601826001600160401b03169052565b506080850151601f19808685030160a08701526125448483611bd4565b935060a0870151915061255b60c087018315159052565b60c08701516001600160401b03811660e0880152915060e087015191508086850301838701525061258c8382611bd4565b9695505050505050565b6000806000606084860312156125ab57600080fd5b83516125b681611b6e565b60208501519093506125c781611b6e565b60408501519092506125d881611b6e565b80915050925092509256fe4d616e74615374616b696e674d6964646c65776172653a206f70657261746f7202dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a2646970667358221220a2ebef894fa76fb9f4b9727442b95453ee20b11d81e765c11ee2ea4182a2c87b64736f6c63430008190033",
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
// Solidity: function operatorSettings() view returns(uint48 unregisterTokenUnlockWindow, uint256 requiredOperatorStake, uint48 minOperatorCommission, uint48 maxOperatorCommission)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) OperatorSettings(opts *bind.CallOpts) (struct {
	UnregisterTokenUnlockWindow *big.Int
	RequiredOperatorStake       *big.Int
	MinOperatorCommission       *big.Int
	MaxOperatorCommission       *big.Int
}, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "operatorSettings")

	outstruct := new(struct {
		UnregisterTokenUnlockWindow *big.Int
		RequiredOperatorStake       *big.Int
		MinOperatorCommission       *big.Int
		MaxOperatorCommission       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UnregisterTokenUnlockWindow = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RequiredOperatorStake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MinOperatorCommission = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxOperatorCommission = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OperatorSettings is a free data retrieval call binding the contract method 0x72aef6bf.
//
// Solidity: function operatorSettings() view returns(uint48 unregisterTokenUnlockWindow, uint256 requiredOperatorStake, uint48 minOperatorCommission, uint48 maxOperatorCommission)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) OperatorSettings() (struct {
	UnregisterTokenUnlockWindow *big.Int
	RequiredOperatorStake       *big.Int
	MinOperatorCommission       *big.Int
	MaxOperatorCommission       *big.Int
}, error) {
	return _MantaStakingMiddleware.Contract.OperatorSettings(&_MantaStakingMiddleware.CallOpts)
}

// OperatorSettings is a free data retrieval call binding the contract method 0x72aef6bf.
//
// Solidity: function operatorSettings() view returns(uint48 unregisterTokenUnlockWindow, uint256 requiredOperatorStake, uint48 minOperatorCommission, uint48 maxOperatorCommission)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) OperatorSettings() (struct {
	UnregisterTokenUnlockWindow *big.Int
	RequiredOperatorStake       *big.Int
	MinOperatorCommission       *big.Int
	MaxOperatorCommission       *big.Int
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

// Initialize is a paid mutator transaction binding the contract method 0x2e5ea91d.
//
// Solidity: function initialize((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) _symbioticVaultSettings, (uint48,uint256,uint48,uint48) _operatorSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) Initialize(opts *bind.TransactOpts, _symbioticVaultSettings SymbioticVaultSettings, _operatorSettings OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "initialize", _symbioticVaultSettings, _operatorSettings)
}

// Initialize is a paid mutator transaction binding the contract method 0x2e5ea91d.
//
// Solidity: function initialize((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) _symbioticVaultSettings, (uint48,uint256,uint48,uint48) _operatorSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) Initialize(_symbioticVaultSettings SymbioticVaultSettings, _operatorSettings OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Initialize(&_MantaStakingMiddleware.TransactOpts, _symbioticVaultSettings, _operatorSettings)
}

// Initialize is a paid mutator transaction binding the contract method 0x2e5ea91d.
//
// Solidity: function initialize((address,address,uint48,address,(uint64,uint64,uint64,address,address,address)) _symbioticVaultSettings, (uint48,uint256,uint48,uint48) _operatorSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) Initialize(_symbioticVaultSettings SymbioticVaultSettings, _operatorSettings OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Initialize(&_MantaStakingMiddleware.TransactOpts, _symbioticVaultSettings, _operatorSettings)
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

// RegisterOperator is a paid mutator transaction binding the contract method 0x829126b8.
//
// Solidity: function registerOperator(bytes _operatorPublicKey, string _operatorName, address _rewardAddress, uint48 _commission) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) RegisterOperator(opts *bind.TransactOpts, _operatorPublicKey []byte, _operatorName string, _rewardAddress common.Address, _commission *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "registerOperator", _operatorPublicKey, _operatorName, _rewardAddress, _commission)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x829126b8.
//
// Solidity: function registerOperator(bytes _operatorPublicKey, string _operatorName, address _rewardAddress, uint48 _commission) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) RegisterOperator(_operatorPublicKey []byte, _operatorName string, _rewardAddress common.Address, _commission *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RegisterOperator(&_MantaStakingMiddleware.TransactOpts, _operatorPublicKey, _operatorName, _rewardAddress, _commission)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x829126b8.
//
// Solidity: function registerOperator(bytes _operatorPublicKey, string _operatorName, address _rewardAddress, uint48 _commission) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) RegisterOperator(_operatorPublicKey []byte, _operatorName string, _rewardAddress common.Address, _commission *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RegisterOperator(&_MantaStakingMiddleware.TransactOpts, _operatorPublicKey, _operatorName, _rewardAddress, _commission)
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

// SetRewardAddress is a paid mutator transaction binding the contract method 0x5e00e679.
//
// Solidity: function setRewardAddress(address _rewardAddress) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) SetRewardAddress(opts *bind.TransactOpts, _rewardAddress common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "setRewardAddress", _rewardAddress)
}

// SetRewardAddress is a paid mutator transaction binding the contract method 0x5e00e679.
//
// Solidity: function setRewardAddress(address _rewardAddress) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) SetRewardAddress(_rewardAddress common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.SetRewardAddress(&_MantaStakingMiddleware.TransactOpts, _rewardAddress)
}

// SetRewardAddress is a paid mutator transaction binding the contract method 0x5e00e679.
//
// Solidity: function setRewardAddress(address _rewardAddress) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) SetRewardAddress(_rewardAddress common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.SetRewardAddress(&_MantaStakingMiddleware.TransactOpts, _rewardAddress)
}

// Slash is a paid mutator transaction binding the contract method 0xa314db16.
//
// Solidity: function slash(bytes32 _subnetwork, address _operator, uint256 _amount, uint48 _captureTimestamp) returns(uint256 slashedAmount)
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) Slash(opts *bind.TransactOpts, _subnetwork [32]byte, _operator common.Address, _amount *big.Int, _captureTimestamp *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "slash", _subnetwork, _operator, _amount, _captureTimestamp)
}

// Slash is a paid mutator transaction binding the contract method 0xa314db16.
//
// Solidity: function slash(bytes32 _subnetwork, address _operator, uint256 _amount, uint48 _captureTimestamp) returns(uint256 slashedAmount)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) Slash(_subnetwork [32]byte, _operator common.Address, _amount *big.Int, _captureTimestamp *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Slash(&_MantaStakingMiddleware.TransactOpts, _subnetwork, _operator, _amount, _captureTimestamp)
}

// Slash is a paid mutator transaction binding the contract method 0xa314db16.
//
// Solidity: function slash(bytes32 _subnetwork, address _operator, uint256 _amount, uint48 _captureTimestamp) returns(uint256 slashedAmount)
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) Slash(_subnetwork [32]byte, _operator common.Address, _amount *big.Int, _captureTimestamp *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Slash(&_MantaStakingMiddleware.TransactOpts, _subnetwork, _operator, _amount, _captureTimestamp)
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

// UpdateOperatorSettings is a paid mutator transaction binding the contract method 0xa204fe55.
//
// Solidity: function updateOperatorSettings((uint48,uint256,uint48,uint48) _operatorSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) UpdateOperatorSettings(opts *bind.TransactOpts, _operatorSettings OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "updateOperatorSettings", _operatorSettings)
}

// UpdateOperatorSettings is a paid mutator transaction binding the contract method 0xa204fe55.
//
// Solidity: function updateOperatorSettings((uint48,uint256,uint48,uint48) _operatorSettings) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) UpdateOperatorSettings(_operatorSettings OperatorSettings) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UpdateOperatorSettings(&_MantaStakingMiddleware.TransactOpts, _operatorSettings)
}

// UpdateOperatorSettings is a paid mutator transaction binding the contract method 0xa204fe55.
//
// Solidity: function updateOperatorSettings((uint48,uint256,uint48,uint48) _operatorSettings) returns()
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
