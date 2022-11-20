// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package daoFactory

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
)

// DaoFactoryMetaData contains all meta data concerning the DaoFactory contract.
var DaoFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_proposalRegistry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"DaoCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_proposalExpirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quorumRequired\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_parentRegistry\",\"type\":\"address\"}],\"name\":\"deployDao\",\"outputs\":[{\"internalType\":\"contractProposalRegistry\",\"name\":\"registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DaoFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use DaoFactoryMetaData.ABI instead.
var DaoFactoryABI = DaoFactoryMetaData.ABI

// DaoFactory is an auto generated Go binding around an Ethereum contract.
type DaoFactory struct {
	DaoFactoryCaller     // Read-only binding to the contract
	DaoFactoryTransactor // Write-only binding to the contract
	DaoFactoryFilterer   // Log filterer for contract events
}

// DaoFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DaoFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DaoFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DaoFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DaoFactorySession struct {
	Contract     *DaoFactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaoFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DaoFactoryCallerSession struct {
	Contract *DaoFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DaoFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DaoFactoryTransactorSession struct {
	Contract     *DaoFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DaoFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DaoFactoryRaw struct {
	Contract *DaoFactory // Generic contract binding to access the raw methods on
}

// DaoFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DaoFactoryCallerRaw struct {
	Contract *DaoFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// DaoFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DaoFactoryTransactorRaw struct {
	Contract *DaoFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDaoFactory creates a new instance of DaoFactory, bound to a specific deployed contract.
func NewDaoFactory(address common.Address, backend bind.ContractBackend) (*DaoFactory, error) {
	contract, err := bindDaoFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DaoFactory{DaoFactoryCaller: DaoFactoryCaller{contract: contract}, DaoFactoryTransactor: DaoFactoryTransactor{contract: contract}, DaoFactoryFilterer: DaoFactoryFilterer{contract: contract}}, nil
}

// NewDaoFactoryCaller creates a new read-only instance of DaoFactory, bound to a specific deployed contract.
func NewDaoFactoryCaller(address common.Address, caller bind.ContractCaller) (*DaoFactoryCaller, error) {
	contract, err := bindDaoFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaoFactoryCaller{contract: contract}, nil
}

// NewDaoFactoryTransactor creates a new write-only instance of DaoFactory, bound to a specific deployed contract.
func NewDaoFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DaoFactoryTransactor, error) {
	contract, err := bindDaoFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaoFactoryTransactor{contract: contract}, nil
}

// NewDaoFactoryFilterer creates a new log filterer instance of DaoFactory, bound to a specific deployed contract.
func NewDaoFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DaoFactoryFilterer, error) {
	contract, err := bindDaoFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaoFactoryFilterer{contract: contract}, nil
}

// bindDaoFactory binds a generic wrapper to an already deployed contract.
func bindDaoFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DaoFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaoFactory *DaoFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DaoFactory.Contract.DaoFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaoFactory *DaoFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaoFactory.Contract.DaoFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaoFactory *DaoFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaoFactory.Contract.DaoFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaoFactory *DaoFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DaoFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaoFactory *DaoFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaoFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaoFactory *DaoFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaoFactory.Contract.contract.Transact(opts, method, params...)
}

// DeployDao is a paid mutator transaction binding the contract method 0x20c893c0.
//
// Solidity: function deployDao(address _governance, uint256 _proposalExpirationTime, uint256 _quorumRequired, address _parentRegistry) returns(address registry)
func (_DaoFactory *DaoFactoryTransactor) DeployDao(opts *bind.TransactOpts, _governance common.Address, _proposalExpirationTime *big.Int, _quorumRequired *big.Int, _parentRegistry common.Address) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "deployDao", _governance, _proposalExpirationTime, _quorumRequired, _parentRegistry)
}

// DeployDao is a paid mutator transaction binding the contract method 0x20c893c0.
//
// Solidity: function deployDao(address _governance, uint256 _proposalExpirationTime, uint256 _quorumRequired, address _parentRegistry) returns(address registry)
func (_DaoFactory *DaoFactorySession) DeployDao(_governance common.Address, _proposalExpirationTime *big.Int, _quorumRequired *big.Int, _parentRegistry common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.DeployDao(&_DaoFactory.TransactOpts, _governance, _proposalExpirationTime, _quorumRequired, _parentRegistry)
}

// DeployDao is a paid mutator transaction binding the contract method 0x20c893c0.
//
// Solidity: function deployDao(address _governance, uint256 _proposalExpirationTime, uint256 _quorumRequired, address _parentRegistry) returns(address registry)
func (_DaoFactory *DaoFactoryTransactorSession) DeployDao(_governance common.Address, _proposalExpirationTime *big.Int, _quorumRequired *big.Int, _parentRegistry common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.DeployDao(&_DaoFactory.TransactOpts, _governance, _proposalExpirationTime, _quorumRequired, _parentRegistry)
}

// DaoFactoryDaoCreatedIterator is returned from FilterDaoCreated and is used to iterate over the raw logs and unpacked data for DaoCreated events raised by the DaoFactory contract.
type DaoFactoryDaoCreatedIterator struct {
	Event *DaoFactoryDaoCreated // Event containing the contract specifics and raw log

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
func (it *DaoFactoryDaoCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryDaoCreated)
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
		it.Event = new(DaoFactoryDaoCreated)
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
func (it *DaoFactoryDaoCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryDaoCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryDaoCreated represents a DaoCreated event raised by the DaoFactory contract.
type DaoFactoryDaoCreated struct {
	ProposalRegistry common.Address
	Governance       common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDaoCreated is a free log retrieval operation binding the contract event 0xdcdaa9087962b773cd00b3088349710fd7f4c555341f3d594ac9a7a8f1a5e9ea.
//
// Solidity: event DaoCreated(address indexed _proposalRegistry, address indexed _governance)
func (_DaoFactory *DaoFactoryFilterer) FilterDaoCreated(opts *bind.FilterOpts, _proposalRegistry []common.Address, _governance []common.Address) (*DaoFactoryDaoCreatedIterator, error) {

	var _proposalRegistryRule []interface{}
	for _, _proposalRegistryItem := range _proposalRegistry {
		_proposalRegistryRule = append(_proposalRegistryRule, _proposalRegistryItem)
	}
	var _governanceRule []interface{}
	for _, _governanceItem := range _governance {
		_governanceRule = append(_governanceRule, _governanceItem)
	}

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "DaoCreated", _proposalRegistryRule, _governanceRule)
	if err != nil {
		return nil, err
	}
	return &DaoFactoryDaoCreatedIterator{contract: _DaoFactory.contract, event: "DaoCreated", logs: logs, sub: sub}, nil
}

// WatchDaoCreated is a free log subscription operation binding the contract event 0xdcdaa9087962b773cd00b3088349710fd7f4c555341f3d594ac9a7a8f1a5e9ea.
//
// Solidity: event DaoCreated(address indexed _proposalRegistry, address indexed _governance)
func (_DaoFactory *DaoFactoryFilterer) WatchDaoCreated(opts *bind.WatchOpts, sink chan<- *DaoFactoryDaoCreated, _proposalRegistry []common.Address, _governance []common.Address) (event.Subscription, error) {

	var _proposalRegistryRule []interface{}
	for _, _proposalRegistryItem := range _proposalRegistry {
		_proposalRegistryRule = append(_proposalRegistryRule, _proposalRegistryItem)
	}
	var _governanceRule []interface{}
	for _, _governanceItem := range _governance {
		_governanceRule = append(_governanceRule, _governanceItem)
	}

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "DaoCreated", _proposalRegistryRule, _governanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryDaoCreated)
				if err := _DaoFactory.contract.UnpackLog(event, "DaoCreated", log); err != nil {
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

// ParseDaoCreated is a log parse operation binding the contract event 0xdcdaa9087962b773cd00b3088349710fd7f4c555341f3d594ac9a7a8f1a5e9ea.
//
// Solidity: event DaoCreated(address indexed _proposalRegistry, address indexed _governance)
func (_DaoFactory *DaoFactoryFilterer) ParseDaoCreated(log types.Log) (*DaoFactoryDaoCreated, error) {
	event := new(DaoFactoryDaoCreated)
	if err := _DaoFactory.contract.UnpackLog(event, "DaoCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
