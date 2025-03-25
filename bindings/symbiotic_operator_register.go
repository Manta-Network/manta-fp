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

// SymbioticOperatorRegisterMetaData contains all meta data concerning the SymbioticOperatorRegister contract.
var SymbioticOperatorRegisterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"entity\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isEntity\",\"inputs\":[{\"name\":\"entity_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalEntities\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddEntity\",\"inputs\":[{\"name\":\"entity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EntityNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegistered\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557610241908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c806314887c581461013d5780632acde098146100d15780635cd8b15e146100b55763b42ba2a214610045575f80fd5b346100b15760203660031901126100b1576004355f5481101561009d575f80527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301546040516001600160a01b039091168152602090f35b634e487b7160e01b5f52603260045260245ffd5b5f80fd5b346100b1575f3660031901126100b15760205f54604051908152f35b346100b1575f3660031901126100b1576100f6335f52600160205260405f2054151590565b61012b5761010333610184565b50337fb919910dcefbf753bfd926ab3b1d3f85d877190c3d01ba1bd585047b99b99f0b5f80a2005b6040516342ee68b560e01b8152600490fd5b346100b15760203660031901126100b1576004356001600160a01b038116908190036100b15761017a6020915f52600160205260405f2054151590565b6040519015158152f35b805f52600160205260405f2054155f14610206575f54680100000000000000008110156101f25760018101805f5581101561009d5781907f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301555f54905f52600160205260405f2055600190565b634e487b7160e01b5f52604160045260245ffd5b505f9056fea264697066735822122077e93113fb982ae02c0a41b82d2748dbea47776eb9ecfa1fc44d11a0333ed06864736f6c63430008190033",
}

// SymbioticOperatorRegisterABI is the input ABI used to generate the binding from.
// Deprecated: Use SymbioticOperatorRegisterMetaData.ABI instead.
var SymbioticOperatorRegisterABI = SymbioticOperatorRegisterMetaData.ABI

// SymbioticOperatorRegisterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SymbioticOperatorRegisterMetaData.Bin instead.
var SymbioticOperatorRegisterBin = SymbioticOperatorRegisterMetaData.Bin

// DeploySymbioticOperatorRegister deploys a new Ethereum contract, binding an instance of SymbioticOperatorRegister to it.
func DeploySymbioticOperatorRegister(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SymbioticOperatorRegister, error) {
	parsed, err := SymbioticOperatorRegisterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SymbioticOperatorRegisterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SymbioticOperatorRegister{SymbioticOperatorRegisterCaller: SymbioticOperatorRegisterCaller{contract: contract}, SymbioticOperatorRegisterTransactor: SymbioticOperatorRegisterTransactor{contract: contract}, SymbioticOperatorRegisterFilterer: SymbioticOperatorRegisterFilterer{contract: contract}}, nil
}

// SymbioticOperatorRegister is an auto generated Go binding around an Ethereum contract.
type SymbioticOperatorRegister struct {
	SymbioticOperatorRegisterCaller     // Read-only binding to the contract
	SymbioticOperatorRegisterTransactor // Write-only binding to the contract
	SymbioticOperatorRegisterFilterer   // Log filterer for contract events
}

// SymbioticOperatorRegisterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SymbioticOperatorRegisterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SymbioticOperatorRegisterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SymbioticOperatorRegisterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SymbioticOperatorRegisterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SymbioticOperatorRegisterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SymbioticOperatorRegisterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SymbioticOperatorRegisterSession struct {
	Contract     *SymbioticOperatorRegister // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SymbioticOperatorRegisterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SymbioticOperatorRegisterCallerSession struct {
	Contract *SymbioticOperatorRegisterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// SymbioticOperatorRegisterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SymbioticOperatorRegisterTransactorSession struct {
	Contract     *SymbioticOperatorRegisterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// SymbioticOperatorRegisterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SymbioticOperatorRegisterRaw struct {
	Contract *SymbioticOperatorRegister // Generic contract binding to access the raw methods on
}

// SymbioticOperatorRegisterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SymbioticOperatorRegisterCallerRaw struct {
	Contract *SymbioticOperatorRegisterCaller // Generic read-only contract binding to access the raw methods on
}

// SymbioticOperatorRegisterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SymbioticOperatorRegisterTransactorRaw struct {
	Contract *SymbioticOperatorRegisterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSymbioticOperatorRegister creates a new instance of SymbioticOperatorRegister, bound to a specific deployed contract.
func NewSymbioticOperatorRegister(address common.Address, backend bind.ContractBackend) (*SymbioticOperatorRegister, error) {
	contract, err := bindSymbioticOperatorRegister(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SymbioticOperatorRegister{SymbioticOperatorRegisterCaller: SymbioticOperatorRegisterCaller{contract: contract}, SymbioticOperatorRegisterTransactor: SymbioticOperatorRegisterTransactor{contract: contract}, SymbioticOperatorRegisterFilterer: SymbioticOperatorRegisterFilterer{contract: contract}}, nil
}

// NewSymbioticOperatorRegisterCaller creates a new read-only instance of SymbioticOperatorRegister, bound to a specific deployed contract.
func NewSymbioticOperatorRegisterCaller(address common.Address, caller bind.ContractCaller) (*SymbioticOperatorRegisterCaller, error) {
	contract, err := bindSymbioticOperatorRegister(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SymbioticOperatorRegisterCaller{contract: contract}, nil
}

// NewSymbioticOperatorRegisterTransactor creates a new write-only instance of SymbioticOperatorRegister, bound to a specific deployed contract.
func NewSymbioticOperatorRegisterTransactor(address common.Address, transactor bind.ContractTransactor) (*SymbioticOperatorRegisterTransactor, error) {
	contract, err := bindSymbioticOperatorRegister(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SymbioticOperatorRegisterTransactor{contract: contract}, nil
}

// NewSymbioticOperatorRegisterFilterer creates a new log filterer instance of SymbioticOperatorRegister, bound to a specific deployed contract.
func NewSymbioticOperatorRegisterFilterer(address common.Address, filterer bind.ContractFilterer) (*SymbioticOperatorRegisterFilterer, error) {
	contract, err := bindSymbioticOperatorRegister(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SymbioticOperatorRegisterFilterer{contract: contract}, nil
}

// bindSymbioticOperatorRegister binds a generic wrapper to an already deployed contract.
func bindSymbioticOperatorRegister(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SymbioticOperatorRegisterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SymbioticOperatorRegister.Contract.SymbioticOperatorRegisterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SymbioticOperatorRegister.Contract.SymbioticOperatorRegisterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SymbioticOperatorRegister.Contract.SymbioticOperatorRegisterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SymbioticOperatorRegister.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SymbioticOperatorRegister.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SymbioticOperatorRegister.Contract.contract.Transact(opts, method, params...)
}

// Entity is a free data retrieval call binding the contract method 0xb42ba2a2.
//
// Solidity: function entity(uint256 index) view returns(address)
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterCaller) Entity(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SymbioticOperatorRegister.contract.Call(opts, &out, "entity", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Entity is a free data retrieval call binding the contract method 0xb42ba2a2.
//
// Solidity: function entity(uint256 index) view returns(address)
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterSession) Entity(index *big.Int) (common.Address, error) {
	return _SymbioticOperatorRegister.Contract.Entity(&_SymbioticOperatorRegister.CallOpts, index)
}

// Entity is a free data retrieval call binding the contract method 0xb42ba2a2.
//
// Solidity: function entity(uint256 index) view returns(address)
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterCallerSession) Entity(index *big.Int) (common.Address, error) {
	return _SymbioticOperatorRegister.Contract.Entity(&_SymbioticOperatorRegister.CallOpts, index)
}

// IsEntity is a free data retrieval call binding the contract method 0x14887c58.
//
// Solidity: function isEntity(address entity_) view returns(bool)
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterCaller) IsEntity(opts *bind.CallOpts, entity_ common.Address) (bool, error) {
	var out []interface{}
	err := _SymbioticOperatorRegister.contract.Call(opts, &out, "isEntity", entity_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEntity is a free data retrieval call binding the contract method 0x14887c58.
//
// Solidity: function isEntity(address entity_) view returns(bool)
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterSession) IsEntity(entity_ common.Address) (bool, error) {
	return _SymbioticOperatorRegister.Contract.IsEntity(&_SymbioticOperatorRegister.CallOpts, entity_)
}

// IsEntity is a free data retrieval call binding the contract method 0x14887c58.
//
// Solidity: function isEntity(address entity_) view returns(bool)
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterCallerSession) IsEntity(entity_ common.Address) (bool, error) {
	return _SymbioticOperatorRegister.Contract.IsEntity(&_SymbioticOperatorRegister.CallOpts, entity_)
}

// TotalEntities is a free data retrieval call binding the contract method 0x5cd8b15e.
//
// Solidity: function totalEntities() view returns(uint256)
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterCaller) TotalEntities(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticOperatorRegister.contract.Call(opts, &out, "totalEntities")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEntities is a free data retrieval call binding the contract method 0x5cd8b15e.
//
// Solidity: function totalEntities() view returns(uint256)
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterSession) TotalEntities() (*big.Int, error) {
	return _SymbioticOperatorRegister.Contract.TotalEntities(&_SymbioticOperatorRegister.CallOpts)
}

// TotalEntities is a free data retrieval call binding the contract method 0x5cd8b15e.
//
// Solidity: function totalEntities() view returns(uint256)
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterCallerSession) TotalEntities() (*big.Int, error) {
	return _SymbioticOperatorRegister.Contract.TotalEntities(&_SymbioticOperatorRegister.CallOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x2acde098.
//
// Solidity: function registerOperator() returns()
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterTransactor) RegisterOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SymbioticOperatorRegister.contract.Transact(opts, "registerOperator")
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x2acde098.
//
// Solidity: function registerOperator() returns()
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterSession) RegisterOperator() (*types.Transaction, error) {
	return _SymbioticOperatorRegister.Contract.RegisterOperator(&_SymbioticOperatorRegister.TransactOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x2acde098.
//
// Solidity: function registerOperator() returns()
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterTransactorSession) RegisterOperator() (*types.Transaction, error) {
	return _SymbioticOperatorRegister.Contract.RegisterOperator(&_SymbioticOperatorRegister.TransactOpts)
}

// SymbioticOperatorRegisterAddEntityIterator is returned from FilterAddEntity and is used to iterate over the raw logs and unpacked data for AddEntity events raised by the SymbioticOperatorRegister contract.
type SymbioticOperatorRegisterAddEntityIterator struct {
	Event *SymbioticOperatorRegisterAddEntity // Event containing the contract specifics and raw log

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
func (it *SymbioticOperatorRegisterAddEntityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticOperatorRegisterAddEntity)
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
		it.Event = new(SymbioticOperatorRegisterAddEntity)
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
func (it *SymbioticOperatorRegisterAddEntityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticOperatorRegisterAddEntityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticOperatorRegisterAddEntity represents a AddEntity event raised by the SymbioticOperatorRegister contract.
type SymbioticOperatorRegisterAddEntity struct {
	Entity common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddEntity is a free log retrieval operation binding the contract event 0xb919910dcefbf753bfd926ab3b1d3f85d877190c3d01ba1bd585047b99b99f0b.
//
// Solidity: event AddEntity(address indexed entity)
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterFilterer) FilterAddEntity(opts *bind.FilterOpts, entity []common.Address) (*SymbioticOperatorRegisterAddEntityIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _SymbioticOperatorRegister.contract.FilterLogs(opts, "AddEntity", entityRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticOperatorRegisterAddEntityIterator{contract: _SymbioticOperatorRegister.contract, event: "AddEntity", logs: logs, sub: sub}, nil
}

// WatchAddEntity is a free log subscription operation binding the contract event 0xb919910dcefbf753bfd926ab3b1d3f85d877190c3d01ba1bd585047b99b99f0b.
//
// Solidity: event AddEntity(address indexed entity)
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterFilterer) WatchAddEntity(opts *bind.WatchOpts, sink chan<- *SymbioticOperatorRegisterAddEntity, entity []common.Address) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _SymbioticOperatorRegister.contract.WatchLogs(opts, "AddEntity", entityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticOperatorRegisterAddEntity)
				if err := _SymbioticOperatorRegister.contract.UnpackLog(event, "AddEntity", log); err != nil {
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

// ParseAddEntity is a log parse operation binding the contract event 0xb919910dcefbf753bfd926ab3b1d3f85d877190c3d01ba1bd585047b99b99f0b.
//
// Solidity: event AddEntity(address indexed entity)
func (_SymbioticOperatorRegister *SymbioticOperatorRegisterFilterer) ParseAddEntity(log types.Log) (*SymbioticOperatorRegisterAddEntity, error) {
	event := new(SymbioticOperatorRegisterAddEntity)
	if err := _SymbioticOperatorRegister.contract.UnpackLog(event, "AddEntity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
