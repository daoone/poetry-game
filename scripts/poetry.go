// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ERC20TokenABI is the input ABI used to generate the binding from.
const ERC20TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20TokenBin is the compiled bytecode used for deploying new contracts.
const ERC20TokenBin = `0x`

// DeployERC20Token deploys a new Ethereum contract, binding an instance of ERC20Token to it.
func DeployERC20Token(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Token, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Token{ERC20TokenCaller: ERC20TokenCaller{contract: contract}, ERC20TokenTransactor: ERC20TokenTransactor{contract: contract}, ERC20TokenFilterer: ERC20TokenFilterer{contract: contract}}, nil
}

// ERC20Token is an auto generated Go binding around an Ethereum contract.
type ERC20Token struct {
	ERC20TokenCaller     // Read-only binding to the contract
	ERC20TokenTransactor // Write-only binding to the contract
	ERC20TokenFilterer   // Log filterer for contract events
}

// ERC20TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenSession struct {
	Contract     *ERC20Token       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenCallerSession struct {
	Contract *ERC20TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenTransactorSession struct {
	Contract     *ERC20TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenRaw struct {
	Contract *ERC20Token // Generic contract binding to access the raw methods on
}

// ERC20TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenCallerRaw struct {
	Contract *ERC20TokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenTransactorRaw struct {
	Contract *ERC20TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Token creates a new instance of ERC20Token, bound to a specific deployed contract.
func NewERC20Token(address common.Address, backend bind.ContractBackend) (*ERC20Token, error) {
	contract, err := bindERC20Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Token{ERC20TokenCaller: ERC20TokenCaller{contract: contract}, ERC20TokenTransactor: ERC20TokenTransactor{contract: contract}, ERC20TokenFilterer: ERC20TokenFilterer{contract: contract}}, nil
}

// NewERC20TokenCaller creates a new read-only instance of ERC20Token, bound to a specific deployed contract.
func NewERC20TokenCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenCaller, error) {
	contract, err := bindERC20Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenCaller{contract: contract}, nil
}

// NewERC20TokenTransactor creates a new write-only instance of ERC20Token, bound to a specific deployed contract.
func NewERC20TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenTransactor, error) {
	contract, err := bindERC20Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenTransactor{contract: contract}, nil
}

// NewERC20TokenFilterer creates a new log filterer instance of ERC20Token, bound to a specific deployed contract.
func NewERC20TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenFilterer, error) {
	contract, err := bindERC20Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenFilterer{contract: contract}, nil
}

// bindERC20Token binds a generic wrapper to an already deployed contract.
func bindERC20Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Token *ERC20TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Token.Contract.ERC20TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Token *ERC20TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Token.Contract.ERC20TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Token *ERC20TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Token.Contract.ERC20TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Token *ERC20TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Token *ERC20TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Token *ERC20TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20Token *ERC20TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20Token *ERC20TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20Token.Contract.Allowance(&_ERC20Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20Token *ERC20TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20Token.Contract.Allowance(&_ERC20Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20Token *ERC20TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20Token *ERC20TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC20Token.Contract.BalanceOf(&_ERC20Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20Token *ERC20TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC20Token.Contract.BalanceOf(&_ERC20Token.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Token *ERC20TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Token *ERC20TokenSession) TotalSupply() (*big.Int, error) {
	return _ERC20Token.Contract.TotalSupply(&_ERC20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Token *ERC20TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Token.Contract.TotalSupply(&_ERC20Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20Token *ERC20TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20Token *ERC20TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Approve(&_ERC20Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20Token *ERC20TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Approve(&_ERC20Token.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ERC20Token *ERC20TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ERC20Token *ERC20TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Transfer(&_ERC20Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ERC20Token *ERC20TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Transfer(&_ERC20Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20Token *ERC20TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20Token *ERC20TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.TransferFrom(&_ERC20Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20Token *ERC20TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.TransferFrom(&_ERC20Token.TransactOpts, _from, _to, _value)
}

// ERC20TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Token contract.
type ERC20TokenApprovalIterator struct {
	Event *ERC20TokenApproval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenApproval)
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
		it.Event = new(ERC20TokenApproval)
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
func (it *ERC20TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenApproval represents a Approval event raised by the ERC20Token contract.
type ERC20TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_ERC20Token *ERC20TokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ERC20TokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ERC20Token.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenApprovalIterator{contract: _ERC20Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_ERC20Token *ERC20TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20TokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ERC20Token.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenApproval)
				if err := _ERC20Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Token contract.
type ERC20TokenTransferIterator struct {
	Event *ERC20TokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenTransfer)
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
		it.Event = new(ERC20TokenTransfer)
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
func (it *ERC20TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenTransfer represents a Transfer event raised by the ERC20Token contract.
type ERC20TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ERC20Token *ERC20TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC20TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC20Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenTransferIterator{contract: _ERC20Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ERC20Token *ERC20TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20TokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC20Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenTransfer)
				if err := _ERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556101668061003b6000396000f30060606040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063a6f9dae11461008c575b600080fd5b341561005b57600080fd5b6100636100ba565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561009757600080fd5b6100b873ffffffffffffffffffffffffffffffffffffffff600435166100d6565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100fe57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff929092169190911790555600a165627a7a72305820f7c467b7c11583c4622d2e1ba59368cf7118b289d650865f65e933cd19e8b3ee0029`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
	OwnedFilterer   // Log filterer for contract events
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// NewOwnedFilterer creates a new log filterer instance of Owned, bound to a specific deployed contract.
func NewOwnedFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnedFilterer, error) {
	contract, err := bindOwned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedFilterer{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Owned *OwnedTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Owned *OwnedSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Owned *OwnedTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, newOwner)
}

// PoetryABI is the input ABI used to generate the binding from.
const PoetryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getVoterReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"poemContent\",\"type\":\"string\"}],\"name\":\"addPoem\",\"outputs\":[{\"name\":\"poemId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gameover\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"xmb\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voteReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poems\",\"outputs\":[{\"name\":\"content\",\"type\":\"string\"},{\"name\":\"votes\",\"type\":\"uint256\"},{\"name\":\"voteCounts\",\"type\":\"uint256\"},{\"name\":\"poetAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rechargeRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxVotes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"poemId\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"votePoem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eachVoterReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winners\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"poemReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"tokenIncrease\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buyXmb\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rechargeLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"decimalUnits\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"poemId\",\"type\":\"uint256\"}],\"name\":\"PoemAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"poemId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PoemVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RewardPushlished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ethValue\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"distrbution\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"RechargeFaith\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"TokenIncrease\",\"type\":\"event\"}]"

// PoetryBin is the compiled bytecode used for deploying new contracts.
const PoetryBin = `0x60606040526001805460a060020a60ff021916905567d02ab486cedc000060025568049b9ca9a6943400006003556801a055690d9db800006008556103e8600955341561004b57600080fd5b604051604080611ca2833981016040528080519190602001805160008054600160a060020a03191633600160a060020a031617905542600a8190559092506100a491506228de806401000000006100fe81026111ba1704565b600b5581816100b1610114565b91825260ff16602082015260409081019051809103906000f08015156100d657600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055506101249050565b60008282018381101561010d57fe5b9392505050565b60405161097a8061132883390190565b6111f5806101336000396000f3006060604052600436106101035763ffffffff60e060020a60003504166315aca6cb8114610108578063228cb7331461011d57806323c6dbc7146101255780632d175f11146101885780633197cbb6146101af5780634940abba146101c2578063542b1c5d146101f15780635b93f270146102045780636ab4d054146102c2578063777c3acf146102d557806378e97925146102e8578063871dcabe146102fb5780638da5cb5b146103145780638f8b656814610327578063a2fb11751461033a578063a6f9dae114610350578063c8c1ccce1461036f578063d96782df14610382578063d9a6788b1461038d578063e27671d914610395578063f8b2cb4f146103a8575b600080fd5b341561011357600080fd5b61011b6103c7565b005b61011b610571565b341561013057600080fd5b61017660046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061074c95505050505050565b60405190815260200160405180910390f35b341561019357600080fd5b61019b61084d565b604051901515815260200160405180910390f35b34156101ba57600080fd5b61017661086e565b34156101cd57600080fd5b6101d5610874565b604051600160a060020a03909116815260200160405180910390f35b34156101fc57600080fd5b610176610883565b341561020f57600080fd5b61021a600435610889565b6040516020810184905260408101839052600160a060020a03821660608201526080808252855460026000196101006001841615020190911604908201819052819060a0820190879080156102b05780601f10610285576101008083540402835291602001916102b0565b820191906000526020600020905b81548152906001019060200180831161029357829003601f168201915b50509550505050505060405180910390f35b34156102cd57600080fd5b6101766108c7565b34156102e057600080fd5b6101766108cd565b34156102f357600080fd5b6101766108d3565b341561030657600080fd5b61011b6004356024356108d9565b341561031f57600080fd5b6101d5610bba565b341561033257600080fd5b610176610bc9565b341561034557600080fd5b610176600435610bcf565b341561035b57600080fd5b61011b600160a060020a0360043516610bee565b341561037a57600080fd5b610176610c38565b61011b600435610c3e565b61011b610d0d565b34156103a057600080fd5b610176610ede565b34156103b357600080fd5b610176600160a060020a0360043516610ee4565b600b5460009042116103d857600080fd5b60015474010000000000000000000000000000000000000000900460ff16151561040157600080fd5b5060005b60075460001901811161056e57600560078281548110151561042357fe5b6000918252602090912001548154811061043957fe5b60009182526020808320600160a060020a03331684526003600590930201919091019052604090205460ff16156105665733600160a060020a03166108fc6004549081150290604051600060405180830381858888f19350505050151561049f57600080fd5b600060056007838154811015156104b257fe5b600091825260209091200154815481106104c857fe5b9060005260206000209060050201600301600033600160a060020a0316600160a060020a0316815260200190815260200160002060006101000a81548160ff0219169083151502179055507f94ee9456b94a89f9e752f560ea9fb37e59c7e4a9ab64ddf4ec32f90b3f16bf703033600454604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b600101610405565b50565b6000806000600b544211151561058657600080fd5b60005433600160a060020a0390811691161480156105b357506003546002540130600160a060020a031631115b15156105be57600080fd5b6007546002546105d39163ffffffff610f5f16565b925060009150600090505b6007546000190181116107315760056007828154811015156105fc57fe5b6000918252602090912001548154811061061257fe5b6000918252602090912060046005909202010154600160a060020a031683156108fc0284604051600060405180830381858888f19350505050151561065657600080fd5b7f94ee9456b94a89f9e752f560ea9fb37e59c7e4a9ab64ddf4ec32f90b3f16bf7030600560078481548110151561068957fe5b6000918252602090912001548154811061069f57fe5b6000918252602090912060046005909202010154600160a060020a031685604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a160056007828154811015156106fa57fe5b6000918252602090912001548154811061071057fe5b600091825260209091206002600590920201015491909101906001016105de565b600354610744908363ffffffff610f5f16565b600455505050565b600080600b54421115151561076057600080fd5b60058054906107729060018301611036565b915060058281548110151561078357fe5b9060005260206000209060050201905082816000019080516107a9929160200190611067565b50600060018281018290556002830182905560048301805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a038116918217909255835260038401602052604092839020805460ff19169092179091557f9da9b3063e7e42e9805b6c1ac6b0cee7236cccda5c4760384ebcb2724e538b8f91849051600160a060020a03909216825260208201526040908101905180910390a150919050565b60015474010000000000000000000000000000000000000000900460ff1681565b600b5481565b600154600160a060020a031681565b60035481565b600580548290811061089757fe5b600091825260209091206005909102016001810154600282015460048301549293509091600160a060020a031684565b60095481565b60065481565b600a5481565b600b546000904211156108eb57600080fd5b600154600090600160a060020a03166327e235e333836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561094657600080fd5b6102c65a03f1151561095757600080fd5b505050604051805190501180156109e15750600154600160a060020a03166327e235e33360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156109c257600080fd5b6102c65a03f115156109d357600080fd5b505050604051805190508211155b15156109ec57600080fd5b60058054849081106109fa57fe5b60009182526020808320600160a060020a03331684526003600590930201918201905260409091205490915060ff1615610a3357600080fd5b600160a060020a03338181166000908152600384016020526040808220805460ff19166001908117909155546004860154908516946323b872dd9493911691879190516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610ac957600080fd5b6102c65a03f11515610ada57600080fd5b5050506040518051505060018082018054840190819055600283018054909201909155600654901115610b1d576001810154600655610b1883610f76565b610b4d565b60065481600101541415610b4d576007805460018101610b3d83826110e5565b5060009182526020909120018390555b60048101547f01360d1c3a49fcb03b7dc68fa217f917ed1f2cd585bc243d8861b3c8eeed1133903390600160a060020a03168585604051600160a060020a0394851681529290931660208301526040808301919091526060820192909252608001905180910390a1505050565b600054600160a060020a031681565b60045481565b6007805482908110610bdd57fe5b600091825260209091200154905081565b60005433600160a060020a03908116911614610c0957600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60025481565b60005433600160a060020a03908116911614610c5957600080fd5b600154600160a060020a031663fea7bcfb8260006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610caa57600080fd5b6102c65a03f11515610cbb57600080fd5b50505060405180519050507fe7301b2ea8a58d6d1a3409482ce1680fe003cd53bf895947c09ef7388d3d1c2f8130600160a060020a03163160405191825260208201526040908101905180910390a150565b600080348190118015610d9a5750600154600160a060020a03166327e235e33060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610d7457600080fd5b6102c65a03f11515610d8557600080fd5b50505060405180519050610d9834610fdc565b105b15610e8957600854341115610e0257600854610dbd90349063ffffffff610ff916565b9050600160a060020a03331681156108fc0282604051600060405180830381858888f193505050501515610df057600080fd5b610dfb600854610fdc565b9150610e0e565b610e0b34610fdc565b91505b600154600160a060020a031663a9059cbb338460006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610e6d57600080fd5b6102c65a03f11515610e7e57600080fd5b505050604051805150505b33600160a060020a03167f44dde0ccedf4eae6acb2415fde025c865191dc80f4165e4ce6f1b310af03d25934848460405180848152602001838152602001828152602001935050505060405180910390a25050565b60085481565b600154600090600160a060020a03166370a0823183836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610f3f57600080fd5b6102c65a03f11515610f5057600080fd5b50505060405180519392505050565b6000808284811515610f6d57fe5b04949350505050565b60075460009081901115610fb7575060005b60075460001901811015610fb7576007805482908110610fa457fe5b6000918252602082200155600101610f88565b6007805460018101610fc983826110e5565b5060009182526020909120019190915550565b6000610ff36009548361100b90919063ffffffff16565b92915050565b60008282111561100557fe5b50900390565b6000828202831580611027575082848281151561102457fe5b04145b151561102f57fe5b9392505050565b815481835581811511611062576005028160050283600052602060002091820191016110629190611109565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110a857805160ff19168380011785556110d5565b828001600101855582156110d5579182015b828111156110d55782518255916020019190600101906110ba565b506110e192915061115c565b5090565b8154818355818115116110625760008381526020902061106291810190830161115c565b61115991905b808211156110e15760006111238282611176565b50600060018201819055600282015560048101805473ffffffffffffffffffffffffffffffffffffffff1916905560050161110f565b90565b61115991905b808211156110e15760008155600101611162565b50805460018160011615610100020316600290046000825580601f1061119c575061056e565b601f01602090049060005260206000209081019061056e919061115c565b60008282018381101561102f57fe00a165627a7a723058207d306c95d31b5fb16b3dcce283581d2b7f2111974382b77df36ff41083ffb3f300296060604052341561000f57600080fd5b60405160408061097a833981016040528080519190602001805160018054600160a060020a03191633600160a060020a03161790559150604090508051908101604052600881527f586d62546f6b656e0000000000000000000000000000000000000000000000006020820152600290805161008f92916020019061010b565b506040805190810160405260038082527f584d42000000000000000000000000000000000000000000000000000000000060208301529080516100d692916020019061010b565b5033600160a060020a03166000908152600560205260409020919091556004805460ff191660ff9092169190911790556101a6565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014c57805160ff1916838001178555610179565b82800160010185558215610179579182015b8281111561017957825182559160200191906001019061015e565b50610185929150610189565b5090565b6101a391905b80821115610185576000815560010161018f565b90565b6107c5806101b56000396000f3006060604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100c9578063095ea7b31461015357806318160ddd1461018957806323b872dd146101ae57806327e235e3146101d6578063313ce567146101f557806370a082311461021e5780638da5cb5b1461023d57806395d89b411461026c578063a6f9dae11461027f578063a9059cbb146102a0578063dd62ed3e146102c2578063fea7bcfb146102e7575b600080fd5b34156100d457600080fd5b6100dc6102fd565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610118578082015183820152602001610100565b50505050905090810190601f1680156101455780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561015e57600080fd5b610175600160a060020a036004351660243561039b565b604051901515815260200160405180910390f35b341561019457600080fd5b61019c610408565b60405190815260200160405180910390f35b34156101b957600080fd5b610175600160a060020a036004358116906024351660443561040e565b34156101e157600080fd5b61019c600160a060020a0360043516610524565b341561020057600080fd5b610208610536565b60405160ff909116815260200160405180910390f35b341561022957600080fd5b61019c600160a060020a036004351661053f565b341561024857600080fd5b61025061055e565b604051600160a060020a03909116815260200160405180910390f35b341561027757600080fd5b6100dc61056d565b341561028a57600080fd5b61029e600160a060020a03600435166105d8565b005b34156102ab57600080fd5b610175600160a060020a0360043516602435610622565b34156102cd57600080fd5b61019c600160a060020a03600435811690602435166106fe565b34156102f257600080fd5b610175600435610729565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103935780601f1061036857610100808354040283529160200191610393565b820191906000526020600020905b81548152906001019060200180831161037657829003601f168201915b505050505081565b600160a060020a03338116600081815260066020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b60015460009033600160a060020a0390811691161461042c57600080fd5b600160a060020a03841660009081526005602052604090205482901080159061047c5750600160a060020a0380851660009081526006602090815260408083203390941683529290522054829010155b80156104885750600082115b1561051957600160a060020a03808416600081815260056020908152604080832080548801905588851680845281842080548990039055600683528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600161051d565b5060005b9392505050565b60056020526000908152604090205481565b60045460ff1681565b600160a060020a0381166000908152600560205260409020545b919050565b600154600160a060020a031681565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103935780601f1061036857610100808354040283529160200191610393565b60015433600160a060020a039081169116146105f357600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60015460009033600160a060020a0390811691161461064057600080fd5b600160a060020a0333166000908152600560205260409020548290108015906106835750600160a060020a03831660009081526005602052604090205482810110155b156106f557600160a060020a033381166000818152600560205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610402565b50600092915050565b600160a060020a03918216600090815260066020908152604080832093909416825291909152205490565b60015460009033600160a060020a0390811691161461074757600080fd5b600154600160a060020a031660009081526005602052604090205482810110610791575060018054600160a060020a03166000908152600560205260409020805483019055610559565b5060009190505600a165627a7a723058200f6907c9644d7948336a4802d1fb01237bb8deba5f2d29319b74f84086c720000029`

// DeployPoetry deploys a new Ethereum contract, binding an instance of Poetry to it.
func DeployPoetry(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, decimalUnits uint8) (common.Address, *types.Transaction, *Poetry, error) {
	parsed, err := abi.JSON(strings.NewReader(PoetryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PoetryBin), backend, initialSupply, decimalUnits)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Poetry{PoetryCaller: PoetryCaller{contract: contract}, PoetryTransactor: PoetryTransactor{contract: contract}, PoetryFilterer: PoetryFilterer{contract: contract}}, nil
}

// Poetry is an auto generated Go binding around an Ethereum contract.
type Poetry struct {
	PoetryCaller     // Read-only binding to the contract
	PoetryTransactor // Write-only binding to the contract
	PoetryFilterer   // Log filterer for contract events
}

// PoetryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoetryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoetryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoetryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoetryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoetryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoetrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoetrySession struct {
	Contract     *Poetry           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoetryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoetryCallerSession struct {
	Contract *PoetryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoetryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoetryTransactorSession struct {
	Contract     *PoetryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoetryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoetryRaw struct {
	Contract *Poetry // Generic contract binding to access the raw methods on
}

// PoetryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoetryCallerRaw struct {
	Contract *PoetryCaller // Generic read-only contract binding to access the raw methods on
}

// PoetryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoetryTransactorRaw struct {
	Contract *PoetryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoetry creates a new instance of Poetry, bound to a specific deployed contract.
func NewPoetry(address common.Address, backend bind.ContractBackend) (*Poetry, error) {
	contract, err := bindPoetry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Poetry{PoetryCaller: PoetryCaller{contract: contract}, PoetryTransactor: PoetryTransactor{contract: contract}, PoetryFilterer: PoetryFilterer{contract: contract}}, nil
}

// NewPoetryCaller creates a new read-only instance of Poetry, bound to a specific deployed contract.
func NewPoetryCaller(address common.Address, caller bind.ContractCaller) (*PoetryCaller, error) {
	contract, err := bindPoetry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoetryCaller{contract: contract}, nil
}

// NewPoetryTransactor creates a new write-only instance of Poetry, bound to a specific deployed contract.
func NewPoetryTransactor(address common.Address, transactor bind.ContractTransactor) (*PoetryTransactor, error) {
	contract, err := bindPoetry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoetryTransactor{contract: contract}, nil
}

// NewPoetryFilterer creates a new log filterer instance of Poetry, bound to a specific deployed contract.
func NewPoetryFilterer(address common.Address, filterer bind.ContractFilterer) (*PoetryFilterer, error) {
	contract, err := bindPoetry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoetryFilterer{contract: contract}, nil
}

// bindPoetry binds a generic wrapper to an already deployed contract.
func bindPoetry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoetryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poetry *PoetryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Poetry.Contract.PoetryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poetry *PoetryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poetry.Contract.PoetryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poetry *PoetryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poetry.Contract.PoetryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poetry *PoetryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Poetry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poetry *PoetryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poetry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poetry *PoetryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poetry.Contract.contract.Transact(opts, method, params...)
}

// EachVoterReward is a free data retrieval call binding the contract method 0x8f8b6568.
//
// Solidity: function eachVoterReward() constant returns(uint256)
func (_Poetry *PoetryCaller) EachVoterReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "eachVoterReward")
	return *ret0, err
}

// EachVoterReward is a free data retrieval call binding the contract method 0x8f8b6568.
//
// Solidity: function eachVoterReward() constant returns(uint256)
func (_Poetry *PoetrySession) EachVoterReward() (*big.Int, error) {
	return _Poetry.Contract.EachVoterReward(&_Poetry.CallOpts)
}

// EachVoterReward is a free data retrieval call binding the contract method 0x8f8b6568.
//
// Solidity: function eachVoterReward() constant returns(uint256)
func (_Poetry *PoetryCallerSession) EachVoterReward() (*big.Int, error) {
	return _Poetry.Contract.EachVoterReward(&_Poetry.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_Poetry *PoetryCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "endTime")
	return *ret0, err
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_Poetry *PoetrySession) EndTime() (*big.Int, error) {
	return _Poetry.Contract.EndTime(&_Poetry.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_Poetry *PoetryCallerSession) EndTime() (*big.Int, error) {
	return _Poetry.Contract.EndTime(&_Poetry.CallOpts)
}

// Gameover is a free data retrieval call binding the contract method 0x2d175f11.
//
// Solidity: function gameover() constant returns(bool)
func (_Poetry *PoetryCaller) Gameover(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "gameover")
	return *ret0, err
}

// Gameover is a free data retrieval call binding the contract method 0x2d175f11.
//
// Solidity: function gameover() constant returns(bool)
func (_Poetry *PoetrySession) Gameover() (bool, error) {
	return _Poetry.Contract.Gameover(&_Poetry.CallOpts)
}

// Gameover is a free data retrieval call binding the contract method 0x2d175f11.
//
// Solidity: function gameover() constant returns(bool)
func (_Poetry *PoetryCallerSession) Gameover() (bool, error) {
	return _Poetry.Contract.Gameover(&_Poetry.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(owner address) constant returns(uint256)
func (_Poetry *PoetryCaller) GetBalance(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "getBalance", owner)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(owner address) constant returns(uint256)
func (_Poetry *PoetrySession) GetBalance(owner common.Address) (*big.Int, error) {
	return _Poetry.Contract.GetBalance(&_Poetry.CallOpts, owner)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(owner address) constant returns(uint256)
func (_Poetry *PoetryCallerSession) GetBalance(owner common.Address) (*big.Int, error) {
	return _Poetry.Contract.GetBalance(&_Poetry.CallOpts, owner)
}

// MaxVotes is a free data retrieval call binding the contract method 0x777c3acf.
//
// Solidity: function maxVotes() constant returns(uint256)
func (_Poetry *PoetryCaller) MaxVotes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "maxVotes")
	return *ret0, err
}

// MaxVotes is a free data retrieval call binding the contract method 0x777c3acf.
//
// Solidity: function maxVotes() constant returns(uint256)
func (_Poetry *PoetrySession) MaxVotes() (*big.Int, error) {
	return _Poetry.Contract.MaxVotes(&_Poetry.CallOpts)
}

// MaxVotes is a free data retrieval call binding the contract method 0x777c3acf.
//
// Solidity: function maxVotes() constant returns(uint256)
func (_Poetry *PoetryCallerSession) MaxVotes() (*big.Int, error) {
	return _Poetry.Contract.MaxVotes(&_Poetry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Poetry *PoetryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Poetry *PoetrySession) Owner() (common.Address, error) {
	return _Poetry.Contract.Owner(&_Poetry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Poetry *PoetryCallerSession) Owner() (common.Address, error) {
	return _Poetry.Contract.Owner(&_Poetry.CallOpts)
}

// PoemReward is a free data retrieval call binding the contract method 0xc8c1ccce.
//
// Solidity: function poemReward() constant returns(uint256)
func (_Poetry *PoetryCaller) PoemReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "poemReward")
	return *ret0, err
}

// PoemReward is a free data retrieval call binding the contract method 0xc8c1ccce.
//
// Solidity: function poemReward() constant returns(uint256)
func (_Poetry *PoetrySession) PoemReward() (*big.Int, error) {
	return _Poetry.Contract.PoemReward(&_Poetry.CallOpts)
}

// PoemReward is a free data retrieval call binding the contract method 0xc8c1ccce.
//
// Solidity: function poemReward() constant returns(uint256)
func (_Poetry *PoetryCallerSession) PoemReward() (*big.Int, error) {
	return _Poetry.Contract.PoemReward(&_Poetry.CallOpts)
}

// Poems is a free data retrieval call binding the contract method 0x5b93f270.
//
// Solidity: function poems( uint256) constant returns(content string, votes uint256, voteCounts uint256, poetAddr address)
func (_Poetry *PoetryCaller) Poems(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Content    string
	Votes      *big.Int
	VoteCounts *big.Int
	PoetAddr   common.Address
}, error) {
	ret := new(struct {
		Content    string
		Votes      *big.Int
		VoteCounts *big.Int
		PoetAddr   common.Address
	})
	out := ret
	err := _Poetry.contract.Call(opts, out, "poems", arg0)
	return *ret, err
}

// Poems is a free data retrieval call binding the contract method 0x5b93f270.
//
// Solidity: function poems( uint256) constant returns(content string, votes uint256, voteCounts uint256, poetAddr address)
func (_Poetry *PoetrySession) Poems(arg0 *big.Int) (struct {
	Content    string
	Votes      *big.Int
	VoteCounts *big.Int
	PoetAddr   common.Address
}, error) {
	return _Poetry.Contract.Poems(&_Poetry.CallOpts, arg0)
}

// Poems is a free data retrieval call binding the contract method 0x5b93f270.
//
// Solidity: function poems( uint256) constant returns(content string, votes uint256, voteCounts uint256, poetAddr address)
func (_Poetry *PoetryCallerSession) Poems(arg0 *big.Int) (struct {
	Content    string
	Votes      *big.Int
	VoteCounts *big.Int
	PoetAddr   common.Address
}, error) {
	return _Poetry.Contract.Poems(&_Poetry.CallOpts, arg0)
}

// RechargeLimit is a free data retrieval call binding the contract method 0xe27671d9.
//
// Solidity: function rechargeLimit() constant returns(uint256)
func (_Poetry *PoetryCaller) RechargeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "rechargeLimit")
	return *ret0, err
}

// RechargeLimit is a free data retrieval call binding the contract method 0xe27671d9.
//
// Solidity: function rechargeLimit() constant returns(uint256)
func (_Poetry *PoetrySession) RechargeLimit() (*big.Int, error) {
	return _Poetry.Contract.RechargeLimit(&_Poetry.CallOpts)
}

// RechargeLimit is a free data retrieval call binding the contract method 0xe27671d9.
//
// Solidity: function rechargeLimit() constant returns(uint256)
func (_Poetry *PoetryCallerSession) RechargeLimit() (*big.Int, error) {
	return _Poetry.Contract.RechargeLimit(&_Poetry.CallOpts)
}

// RechargeRate is a free data retrieval call binding the contract method 0x6ab4d054.
//
// Solidity: function rechargeRate() constant returns(uint256)
func (_Poetry *PoetryCaller) RechargeRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "rechargeRate")
	return *ret0, err
}

// RechargeRate is a free data retrieval call binding the contract method 0x6ab4d054.
//
// Solidity: function rechargeRate() constant returns(uint256)
func (_Poetry *PoetrySession) RechargeRate() (*big.Int, error) {
	return _Poetry.Contract.RechargeRate(&_Poetry.CallOpts)
}

// RechargeRate is a free data retrieval call binding the contract method 0x6ab4d054.
//
// Solidity: function rechargeRate() constant returns(uint256)
func (_Poetry *PoetryCallerSession) RechargeRate() (*big.Int, error) {
	return _Poetry.Contract.RechargeRate(&_Poetry.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_Poetry *PoetryCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "startTime")
	return *ret0, err
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_Poetry *PoetrySession) StartTime() (*big.Int, error) {
	return _Poetry.Contract.StartTime(&_Poetry.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_Poetry *PoetryCallerSession) StartTime() (*big.Int, error) {
	return _Poetry.Contract.StartTime(&_Poetry.CallOpts)
}

// VoteReward is a free data retrieval call binding the contract method 0x542b1c5d.
//
// Solidity: function voteReward() constant returns(uint256)
func (_Poetry *PoetryCaller) VoteReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "voteReward")
	return *ret0, err
}

// VoteReward is a free data retrieval call binding the contract method 0x542b1c5d.
//
// Solidity: function voteReward() constant returns(uint256)
func (_Poetry *PoetrySession) VoteReward() (*big.Int, error) {
	return _Poetry.Contract.VoteReward(&_Poetry.CallOpts)
}

// VoteReward is a free data retrieval call binding the contract method 0x542b1c5d.
//
// Solidity: function voteReward() constant returns(uint256)
func (_Poetry *PoetryCallerSession) VoteReward() (*big.Int, error) {
	return _Poetry.Contract.VoteReward(&_Poetry.CallOpts)
}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners( uint256) constant returns(uint256)
func (_Poetry *PoetryCaller) Winners(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "winners", arg0)
	return *ret0, err
}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners( uint256) constant returns(uint256)
func (_Poetry *PoetrySession) Winners(arg0 *big.Int) (*big.Int, error) {
	return _Poetry.Contract.Winners(&_Poetry.CallOpts, arg0)
}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners( uint256) constant returns(uint256)
func (_Poetry *PoetryCallerSession) Winners(arg0 *big.Int) (*big.Int, error) {
	return _Poetry.Contract.Winners(&_Poetry.CallOpts, arg0)
}

// Xmb is a free data retrieval call binding the contract method 0x4940abba.
//
// Solidity: function xmb() constant returns(address)
func (_Poetry *PoetryCaller) Xmb(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Poetry.contract.Call(opts, out, "xmb")
	return *ret0, err
}

// Xmb is a free data retrieval call binding the contract method 0x4940abba.
//
// Solidity: function xmb() constant returns(address)
func (_Poetry *PoetrySession) Xmb() (common.Address, error) {
	return _Poetry.Contract.Xmb(&_Poetry.CallOpts)
}

// Xmb is a free data retrieval call binding the contract method 0x4940abba.
//
// Solidity: function xmb() constant returns(address)
func (_Poetry *PoetryCallerSession) Xmb() (common.Address, error) {
	return _Poetry.Contract.Xmb(&_Poetry.CallOpts)
}

// AddPoem is a paid mutator transaction binding the contract method 0x23c6dbc7.
//
// Solidity: function addPoem(poemContent string) returns(poemId uint256)
func (_Poetry *PoetryTransactor) AddPoem(opts *bind.TransactOpts, poemContent string) (*types.Transaction, error) {
	return _Poetry.contract.Transact(opts, "addPoem", poemContent)
}

// AddPoem is a paid mutator transaction binding the contract method 0x23c6dbc7.
//
// Solidity: function addPoem(poemContent string) returns(poemId uint256)
func (_Poetry *PoetrySession) AddPoem(poemContent string) (*types.Transaction, error) {
	return _Poetry.Contract.AddPoem(&_Poetry.TransactOpts, poemContent)
}

// AddPoem is a paid mutator transaction binding the contract method 0x23c6dbc7.
//
// Solidity: function addPoem(poemContent string) returns(poemId uint256)
func (_Poetry *PoetryTransactorSession) AddPoem(poemContent string) (*types.Transaction, error) {
	return _Poetry.Contract.AddPoem(&_Poetry.TransactOpts, poemContent)
}

// BuyXmb is a paid mutator transaction binding the contract method 0xd9a6788b.
//
// Solidity: function buyXmb() returns()
func (_Poetry *PoetryTransactor) BuyXmb(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poetry.contract.Transact(opts, "buyXmb")
}

// BuyXmb is a paid mutator transaction binding the contract method 0xd9a6788b.
//
// Solidity: function buyXmb() returns()
func (_Poetry *PoetrySession) BuyXmb() (*types.Transaction, error) {
	return _Poetry.Contract.BuyXmb(&_Poetry.TransactOpts)
}

// BuyXmb is a paid mutator transaction binding the contract method 0xd9a6788b.
//
// Solidity: function buyXmb() returns()
func (_Poetry *PoetryTransactorSession) BuyXmb() (*types.Transaction, error) {
	return _Poetry.Contract.BuyXmb(&_Poetry.TransactOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Poetry *PoetryTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Poetry.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Poetry *PoetrySession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Poetry.Contract.ChangeOwner(&_Poetry.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Poetry *PoetryTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Poetry.Contract.ChangeOwner(&_Poetry.TransactOpts, newOwner)
}

// GetVoterReward is a paid mutator transaction binding the contract method 0x15aca6cb.
//
// Solidity: function getVoterReward() returns()
func (_Poetry *PoetryTransactor) GetVoterReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poetry.contract.Transact(opts, "getVoterReward")
}

// GetVoterReward is a paid mutator transaction binding the contract method 0x15aca6cb.
//
// Solidity: function getVoterReward() returns()
func (_Poetry *PoetrySession) GetVoterReward() (*types.Transaction, error) {
	return _Poetry.Contract.GetVoterReward(&_Poetry.TransactOpts)
}

// GetVoterReward is a paid mutator transaction binding the contract method 0x15aca6cb.
//
// Solidity: function getVoterReward() returns()
func (_Poetry *PoetryTransactorSession) GetVoterReward() (*types.Transaction, error) {
	return _Poetry.Contract.GetVoterReward(&_Poetry.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_Poetry *PoetryTransactor) Reward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poetry.contract.Transact(opts, "reward")
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_Poetry *PoetrySession) Reward() (*types.Transaction, error) {
	return _Poetry.Contract.Reward(&_Poetry.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_Poetry *PoetryTransactorSession) Reward() (*types.Transaction, error) {
	return _Poetry.Contract.Reward(&_Poetry.TransactOpts)
}

// TokenIncrease is a paid mutator transaction binding the contract method 0xd96782df.
//
// Solidity: function tokenIncrease(value uint256) returns()
func (_Poetry *PoetryTransactor) TokenIncrease(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Poetry.contract.Transact(opts, "tokenIncrease", value)
}

// TokenIncrease is a paid mutator transaction binding the contract method 0xd96782df.
//
// Solidity: function tokenIncrease(value uint256) returns()
func (_Poetry *PoetrySession) TokenIncrease(value *big.Int) (*types.Transaction, error) {
	return _Poetry.Contract.TokenIncrease(&_Poetry.TransactOpts, value)
}

// TokenIncrease is a paid mutator transaction binding the contract method 0xd96782df.
//
// Solidity: function tokenIncrease(value uint256) returns()
func (_Poetry *PoetryTransactorSession) TokenIncrease(value *big.Int) (*types.Transaction, error) {
	return _Poetry.Contract.TokenIncrease(&_Poetry.TransactOpts, value)
}

// VotePoem is a paid mutator transaction binding the contract method 0x871dcabe.
//
// Solidity: function votePoem(poemId uint256, _value uint256) returns()
func (_Poetry *PoetryTransactor) VotePoem(opts *bind.TransactOpts, poemId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Poetry.contract.Transact(opts, "votePoem", poemId, _value)
}

// VotePoem is a paid mutator transaction binding the contract method 0x871dcabe.
//
// Solidity: function votePoem(poemId uint256, _value uint256) returns()
func (_Poetry *PoetrySession) VotePoem(poemId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Poetry.Contract.VotePoem(&_Poetry.TransactOpts, poemId, _value)
}

// VotePoem is a paid mutator transaction binding the contract method 0x871dcabe.
//
// Solidity: function votePoem(poemId uint256, _value uint256) returns()
func (_Poetry *PoetryTransactorSession) VotePoem(poemId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Poetry.Contract.VotePoem(&_Poetry.TransactOpts, poemId, _value)
}

// PoetryPoemAddedIterator is returned from FilterPoemAdded and is used to iterate over the raw logs and unpacked data for PoemAdded events raised by the Poetry contract.
type PoetryPoemAddedIterator struct {
	Event *PoetryPoemAdded // Event containing the contract specifics and raw log

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
func (it *PoetryPoemAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoetryPoemAdded)
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
		it.Event = new(PoetryPoemAdded)
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
func (it *PoetryPoemAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoetryPoemAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoetryPoemAdded represents a PoemAdded event raised by the Poetry contract.
type PoetryPoemAdded struct {
	From   common.Address
	PoemId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoemAdded is a free log retrieval operation binding the contract event 0x9da9b3063e7e42e9805b6c1ac6b0cee7236cccda5c4760384ebcb2724e538b8f.
//
// Solidity: event PoemAdded(from address, poemId uint256)
func (_Poetry *PoetryFilterer) FilterPoemAdded(opts *bind.FilterOpts) (*PoetryPoemAddedIterator, error) {

	logs, sub, err := _Poetry.contract.FilterLogs(opts, "PoemAdded")
	if err != nil {
		return nil, err
	}
	return &PoetryPoemAddedIterator{contract: _Poetry.contract, event: "PoemAdded", logs: logs, sub: sub}, nil
}

// WatchPoemAdded is a free log subscription operation binding the contract event 0x9da9b3063e7e42e9805b6c1ac6b0cee7236cccda5c4760384ebcb2724e538b8f.
//
// Solidity: event PoemAdded(from address, poemId uint256)
func (_Poetry *PoetryFilterer) WatchPoemAdded(opts *bind.WatchOpts, sink chan<- *PoetryPoemAdded) (event.Subscription, error) {

	logs, sub, err := _Poetry.contract.WatchLogs(opts, "PoemAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoetryPoemAdded)
				if err := _Poetry.contract.UnpackLog(event, "PoemAdded", log); err != nil {
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

// PoetryPoemVotedIterator is returned from FilterPoemVoted and is used to iterate over the raw logs and unpacked data for PoemVoted events raised by the Poetry contract.
type PoetryPoemVotedIterator struct {
	Event *PoetryPoemVoted // Event containing the contract specifics and raw log

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
func (it *PoetryPoemVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoetryPoemVoted)
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
		it.Event = new(PoetryPoemVoted)
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
func (it *PoetryPoemVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoetryPoemVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoetryPoemVoted represents a PoemVoted event raised by the Poetry contract.
type PoetryPoemVoted struct {
	From   common.Address
	To     common.Address
	PoemId *big.Int
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoemVoted is a free log retrieval operation binding the contract event 0x01360d1c3a49fcb03b7dc68fa217f917ed1f2cd585bc243d8861b3c8eeed1133.
//
// Solidity: event PoemVoted(from address, to address, poemId uint256, value uint256)
func (_Poetry *PoetryFilterer) FilterPoemVoted(opts *bind.FilterOpts) (*PoetryPoemVotedIterator, error) {

	logs, sub, err := _Poetry.contract.FilterLogs(opts, "PoemVoted")
	if err != nil {
		return nil, err
	}
	return &PoetryPoemVotedIterator{contract: _Poetry.contract, event: "PoemVoted", logs: logs, sub: sub}, nil
}

// WatchPoemVoted is a free log subscription operation binding the contract event 0x01360d1c3a49fcb03b7dc68fa217f917ed1f2cd585bc243d8861b3c8eeed1133.
//
// Solidity: event PoemVoted(from address, to address, poemId uint256, value uint256)
func (_Poetry *PoetryFilterer) WatchPoemVoted(opts *bind.WatchOpts, sink chan<- *PoetryPoemVoted) (event.Subscription, error) {

	logs, sub, err := _Poetry.contract.WatchLogs(opts, "PoemVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoetryPoemVoted)
				if err := _Poetry.contract.UnpackLog(event, "PoemVoted", log); err != nil {
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

// PoetryRechargeFaithIterator is returned from FilterRechargeFaith and is used to iterate over the raw logs and unpacked data for RechargeFaith events raised by the Poetry contract.
type PoetryRechargeFaithIterator struct {
	Event *PoetryRechargeFaith // Event containing the contract specifics and raw log

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
func (it *PoetryRechargeFaithIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoetryRechargeFaith)
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
		it.Event = new(PoetryRechargeFaith)
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
func (it *PoetryRechargeFaithIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoetryRechargeFaithIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoetryRechargeFaith represents a RechargeFaith event raised by the Poetry contract.
type PoetryRechargeFaith struct {
	To          common.Address
	EthValue    *big.Int
	Distrbution *big.Int
	Refund      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRechargeFaith is a free log retrieval operation binding the contract event 0x44dde0ccedf4eae6acb2415fde025c865191dc80f4165e4ce6f1b310af03d259.
//
// Solidity: event RechargeFaith(to indexed address, ethValue uint256, distrbution uint256, refund uint256)
func (_Poetry *PoetryFilterer) FilterRechargeFaith(opts *bind.FilterOpts, to []common.Address) (*PoetryRechargeFaithIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Poetry.contract.FilterLogs(opts, "RechargeFaith", toRule)
	if err != nil {
		return nil, err
	}
	return &PoetryRechargeFaithIterator{contract: _Poetry.contract, event: "RechargeFaith", logs: logs, sub: sub}, nil
}

// WatchRechargeFaith is a free log subscription operation binding the contract event 0x44dde0ccedf4eae6acb2415fde025c865191dc80f4165e4ce6f1b310af03d259.
//
// Solidity: event RechargeFaith(to indexed address, ethValue uint256, distrbution uint256, refund uint256)
func (_Poetry *PoetryFilterer) WatchRechargeFaith(opts *bind.WatchOpts, sink chan<- *PoetryRechargeFaith, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Poetry.contract.WatchLogs(opts, "RechargeFaith", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoetryRechargeFaith)
				if err := _Poetry.contract.UnpackLog(event, "RechargeFaith", log); err != nil {
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

// PoetryRewardPushlishedIterator is returned from FilterRewardPushlished and is used to iterate over the raw logs and unpacked data for RewardPushlished events raised by the Poetry contract.
type PoetryRewardPushlishedIterator struct {
	Event *PoetryRewardPushlished // Event containing the contract specifics and raw log

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
func (it *PoetryRewardPushlishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoetryRewardPushlished)
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
		it.Event = new(PoetryRewardPushlished)
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
func (it *PoetryRewardPushlishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoetryRewardPushlishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoetryRewardPushlished represents a RewardPushlished event raised by the Poetry contract.
type PoetryRewardPushlished struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRewardPushlished is a free log retrieval operation binding the contract event 0x94ee9456b94a89f9e752f560ea9fb37e59c7e4a9ab64ddf4ec32f90b3f16bf70.
//
// Solidity: event RewardPushlished(from address, to address, value uint256)
func (_Poetry *PoetryFilterer) FilterRewardPushlished(opts *bind.FilterOpts) (*PoetryRewardPushlishedIterator, error) {

	logs, sub, err := _Poetry.contract.FilterLogs(opts, "RewardPushlished")
	if err != nil {
		return nil, err
	}
	return &PoetryRewardPushlishedIterator{contract: _Poetry.contract, event: "RewardPushlished", logs: logs, sub: sub}, nil
}

// WatchRewardPushlished is a free log subscription operation binding the contract event 0x94ee9456b94a89f9e752f560ea9fb37e59c7e4a9ab64ddf4ec32f90b3f16bf70.
//
// Solidity: event RewardPushlished(from address, to address, value uint256)
func (_Poetry *PoetryFilterer) WatchRewardPushlished(opts *bind.WatchOpts, sink chan<- *PoetryRewardPushlished) (event.Subscription, error) {

	logs, sub, err := _Poetry.contract.WatchLogs(opts, "RewardPushlished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoetryRewardPushlished)
				if err := _Poetry.contract.UnpackLog(event, "RewardPushlished", log); err != nil {
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

// PoetryTokenIncreaseIterator is returned from FilterTokenIncrease and is used to iterate over the raw logs and unpacked data for TokenIncrease events raised by the Poetry contract.
type PoetryTokenIncreaseIterator struct {
	Event *PoetryTokenIncrease // Event containing the contract specifics and raw log

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
func (it *PoetryTokenIncreaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoetryTokenIncrease)
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
		it.Event = new(PoetryTokenIncrease)
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
func (it *PoetryTokenIncreaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoetryTokenIncreaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoetryTokenIncrease represents a TokenIncrease event raised by the Poetry contract.
type PoetryTokenIncrease struct {
	Value   *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenIncrease is a free log retrieval operation binding the contract event 0xe7301b2ea8a58d6d1a3409482ce1680fe003cd53bf895947c09ef7388d3d1c2f.
//
// Solidity: event TokenIncrease(value uint256, balance uint256)
func (_Poetry *PoetryFilterer) FilterTokenIncrease(opts *bind.FilterOpts) (*PoetryTokenIncreaseIterator, error) {

	logs, sub, err := _Poetry.contract.FilterLogs(opts, "TokenIncrease")
	if err != nil {
		return nil, err
	}
	return &PoetryTokenIncreaseIterator{contract: _Poetry.contract, event: "TokenIncrease", logs: logs, sub: sub}, nil
}

// WatchTokenIncrease is a free log subscription operation binding the contract event 0xe7301b2ea8a58d6d1a3409482ce1680fe003cd53bf895947c09ef7388d3d1c2f.
//
// Solidity: event TokenIncrease(value uint256, balance uint256)
func (_Poetry *PoetryFilterer) WatchTokenIncrease(opts *bind.WatchOpts, sink chan<- *PoetryTokenIncrease) (event.Subscription, error) {

	logs, sub, err := _Poetry.contract.WatchLogs(opts, "TokenIncrease")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoetryTokenIncrease)
				if err := _Poetry.contract.UnpackLog(event, "TokenIncrease", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a7230582054c8abb38fb4e87a2331d53a1fc7627cd07e5aeaa864574971b34b2dea8d2ec60029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// XmbTokenABI is the input ABI used to generate the binding from.
const XmbTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"additional\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"decimalUnits\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// XmbTokenBin is the compiled bytecode used for deploying new contracts.
const XmbTokenBin = `0x6060604052341561000f57600080fd5b60405160408061097a833981016040528080519190602001805160018054600160a060020a03191633600160a060020a03161790559150604090508051908101604052600881527f586d62546f6b656e0000000000000000000000000000000000000000000000006020820152600290805161008f92916020019061010b565b506040805190810160405260038082527f584d42000000000000000000000000000000000000000000000000000000000060208301529080516100d692916020019061010b565b5033600160a060020a03166000908152600560205260409020919091556004805460ff191660ff9092169190911790556101a6565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014c57805160ff1916838001178555610179565b82800160010185558215610179579182015b8281111561017957825182559160200191906001019061015e565b50610185929150610189565b5090565b6101a391905b80821115610185576000815560010161018f565b90565b6107c5806101b56000396000f3006060604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100c9578063095ea7b31461015357806318160ddd1461018957806323b872dd146101ae57806327e235e3146101d6578063313ce567146101f557806370a082311461021e5780638da5cb5b1461023d57806395d89b411461026c578063a6f9dae11461027f578063a9059cbb146102a0578063dd62ed3e146102c2578063fea7bcfb146102e7575b600080fd5b34156100d457600080fd5b6100dc6102fd565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610118578082015183820152602001610100565b50505050905090810190601f1680156101455780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561015e57600080fd5b610175600160a060020a036004351660243561039b565b604051901515815260200160405180910390f35b341561019457600080fd5b61019c610408565b60405190815260200160405180910390f35b34156101b957600080fd5b610175600160a060020a036004358116906024351660443561040e565b34156101e157600080fd5b61019c600160a060020a0360043516610524565b341561020057600080fd5b610208610536565b60405160ff909116815260200160405180910390f35b341561022957600080fd5b61019c600160a060020a036004351661053f565b341561024857600080fd5b61025061055e565b604051600160a060020a03909116815260200160405180910390f35b341561027757600080fd5b6100dc61056d565b341561028a57600080fd5b61029e600160a060020a03600435166105d8565b005b34156102ab57600080fd5b610175600160a060020a0360043516602435610622565b34156102cd57600080fd5b61019c600160a060020a03600435811690602435166106fe565b34156102f257600080fd5b610175600435610729565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103935780601f1061036857610100808354040283529160200191610393565b820191906000526020600020905b81548152906001019060200180831161037657829003601f168201915b505050505081565b600160a060020a03338116600081815260066020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b60015460009033600160a060020a0390811691161461042c57600080fd5b600160a060020a03841660009081526005602052604090205482901080159061047c5750600160a060020a0380851660009081526006602090815260408083203390941683529290522054829010155b80156104885750600082115b1561051957600160a060020a03808416600081815260056020908152604080832080548801905588851680845281842080548990039055600683528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600161051d565b5060005b9392505050565b60056020526000908152604090205481565b60045460ff1681565b600160a060020a0381166000908152600560205260409020545b919050565b600154600160a060020a031681565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103935780601f1061036857610100808354040283529160200191610393565b60015433600160a060020a039081169116146105f357600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60015460009033600160a060020a0390811691161461064057600080fd5b600160a060020a0333166000908152600560205260409020548290108015906106835750600160a060020a03831660009081526005602052604090205482810110155b156106f557600160a060020a033381166000818152600560205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610402565b50600092915050565b600160a060020a03918216600090815260066020908152604080832093909416825291909152205490565b60015460009033600160a060020a0390811691161461074757600080fd5b600154600160a060020a031660009081526005602052604090205482810110610791575060018054600160a060020a03166000908152600560205260409020805483019055610559565b5060009190505600a165627a7a723058200f6907c9644d7948336a4802d1fb01237bb8deba5f2d29319b74f84086c720000029`

// DeployXmbToken deploys a new Ethereum contract, binding an instance of XmbToken to it.
func DeployXmbToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, decimalUnits uint8) (common.Address, *types.Transaction, *XmbToken, error) {
	parsed, err := abi.JSON(strings.NewReader(XmbTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XmbTokenBin), backend, initialSupply, decimalUnits)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XmbToken{XmbTokenCaller: XmbTokenCaller{contract: contract}, XmbTokenTransactor: XmbTokenTransactor{contract: contract}, XmbTokenFilterer: XmbTokenFilterer{contract: contract}}, nil
}

// XmbToken is an auto generated Go binding around an Ethereum contract.
type XmbToken struct {
	XmbTokenCaller     // Read-only binding to the contract
	XmbTokenTransactor // Write-only binding to the contract
	XmbTokenFilterer   // Log filterer for contract events
}

// XmbTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type XmbTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XmbTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XmbTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XmbTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XmbTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XmbTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XmbTokenSession struct {
	Contract     *XmbToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XmbTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XmbTokenCallerSession struct {
	Contract *XmbTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// XmbTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XmbTokenTransactorSession struct {
	Contract     *XmbTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// XmbTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type XmbTokenRaw struct {
	Contract *XmbToken // Generic contract binding to access the raw methods on
}

// XmbTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XmbTokenCallerRaw struct {
	Contract *XmbTokenCaller // Generic read-only contract binding to access the raw methods on
}

// XmbTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XmbTokenTransactorRaw struct {
	Contract *XmbTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXmbToken creates a new instance of XmbToken, bound to a specific deployed contract.
func NewXmbToken(address common.Address, backend bind.ContractBackend) (*XmbToken, error) {
	contract, err := bindXmbToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XmbToken{XmbTokenCaller: XmbTokenCaller{contract: contract}, XmbTokenTransactor: XmbTokenTransactor{contract: contract}, XmbTokenFilterer: XmbTokenFilterer{contract: contract}}, nil
}

// NewXmbTokenCaller creates a new read-only instance of XmbToken, bound to a specific deployed contract.
func NewXmbTokenCaller(address common.Address, caller bind.ContractCaller) (*XmbTokenCaller, error) {
	contract, err := bindXmbToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XmbTokenCaller{contract: contract}, nil
}

// NewXmbTokenTransactor creates a new write-only instance of XmbToken, bound to a specific deployed contract.
func NewXmbTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*XmbTokenTransactor, error) {
	contract, err := bindXmbToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XmbTokenTransactor{contract: contract}, nil
}

// NewXmbTokenFilterer creates a new log filterer instance of XmbToken, bound to a specific deployed contract.
func NewXmbTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*XmbTokenFilterer, error) {
	contract, err := bindXmbToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XmbTokenFilterer{contract: contract}, nil
}

// bindXmbToken binds a generic wrapper to an already deployed contract.
func bindXmbToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XmbTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XmbToken *XmbTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XmbToken.Contract.XmbTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XmbToken *XmbTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XmbToken.Contract.XmbTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XmbToken *XmbTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XmbToken.Contract.XmbTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XmbToken *XmbTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XmbToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XmbToken *XmbTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XmbToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XmbToken *XmbTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XmbToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_XmbToken *XmbTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XmbToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_XmbToken *XmbTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _XmbToken.Contract.Allowance(&_XmbToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_XmbToken *XmbTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _XmbToken.Contract.Allowance(&_XmbToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_XmbToken *XmbTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XmbToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_XmbToken *XmbTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _XmbToken.Contract.BalanceOf(&_XmbToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_XmbToken *XmbTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _XmbToken.Contract.BalanceOf(&_XmbToken.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_XmbToken *XmbTokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XmbToken.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_XmbToken *XmbTokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _XmbToken.Contract.Balances(&_XmbToken.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_XmbToken *XmbTokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _XmbToken.Contract.Balances(&_XmbToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_XmbToken *XmbTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _XmbToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_XmbToken *XmbTokenSession) Decimals() (uint8, error) {
	return _XmbToken.Contract.Decimals(&_XmbToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_XmbToken *XmbTokenCallerSession) Decimals() (uint8, error) {
	return _XmbToken.Contract.Decimals(&_XmbToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_XmbToken *XmbTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _XmbToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_XmbToken *XmbTokenSession) Name() (string, error) {
	return _XmbToken.Contract.Name(&_XmbToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_XmbToken *XmbTokenCallerSession) Name() (string, error) {
	return _XmbToken.Contract.Name(&_XmbToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_XmbToken *XmbTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XmbToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_XmbToken *XmbTokenSession) Owner() (common.Address, error) {
	return _XmbToken.Contract.Owner(&_XmbToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_XmbToken *XmbTokenCallerSession) Owner() (common.Address, error) {
	return _XmbToken.Contract.Owner(&_XmbToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_XmbToken *XmbTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _XmbToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_XmbToken *XmbTokenSession) Symbol() (string, error) {
	return _XmbToken.Contract.Symbol(&_XmbToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_XmbToken *XmbTokenCallerSession) Symbol() (string, error) {
	return _XmbToken.Contract.Symbol(&_XmbToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_XmbToken *XmbTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XmbToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_XmbToken *XmbTokenSession) TotalSupply() (*big.Int, error) {
	return _XmbToken.Contract.TotalSupply(&_XmbToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_XmbToken *XmbTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _XmbToken.Contract.TotalSupply(&_XmbToken.CallOpts)
}

// Additional is a paid mutator transaction binding the contract method 0xfea7bcfb.
//
// Solidity: function additional(_value uint256) returns(success bool)
func (_XmbToken *XmbTokenTransactor) Additional(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _XmbToken.contract.Transact(opts, "additional", _value)
}

// Additional is a paid mutator transaction binding the contract method 0xfea7bcfb.
//
// Solidity: function additional(_value uint256) returns(success bool)
func (_XmbToken *XmbTokenSession) Additional(_value *big.Int) (*types.Transaction, error) {
	return _XmbToken.Contract.Additional(&_XmbToken.TransactOpts, _value)
}

// Additional is a paid mutator transaction binding the contract method 0xfea7bcfb.
//
// Solidity: function additional(_value uint256) returns(success bool)
func (_XmbToken *XmbTokenTransactorSession) Additional(_value *big.Int) (*types.Transaction, error) {
	return _XmbToken.Contract.Additional(&_XmbToken.TransactOpts, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_XmbToken *XmbTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _XmbToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_XmbToken *XmbTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _XmbToken.Contract.Approve(&_XmbToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_XmbToken *XmbTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _XmbToken.Contract.Approve(&_XmbToken.TransactOpts, _spender, _value)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_XmbToken *XmbTokenTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _XmbToken.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_XmbToken *XmbTokenSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _XmbToken.Contract.ChangeOwner(&_XmbToken.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_XmbToken *XmbTokenTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _XmbToken.Contract.ChangeOwner(&_XmbToken.TransactOpts, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_XmbToken *XmbTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _XmbToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_XmbToken *XmbTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _XmbToken.Contract.Transfer(&_XmbToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_XmbToken *XmbTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _XmbToken.Contract.Transfer(&_XmbToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_XmbToken *XmbTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _XmbToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_XmbToken *XmbTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _XmbToken.Contract.TransferFrom(&_XmbToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_XmbToken *XmbTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _XmbToken.Contract.TransferFrom(&_XmbToken.TransactOpts, _from, _to, _value)
}

// XmbTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the XmbToken contract.
type XmbTokenApprovalIterator struct {
	Event *XmbTokenApproval // Event containing the contract specifics and raw log

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
func (it *XmbTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmbTokenApproval)
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
		it.Event = new(XmbTokenApproval)
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
func (it *XmbTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmbTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmbTokenApproval represents a Approval event raised by the XmbToken contract.
type XmbTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_XmbToken *XmbTokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*XmbTokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _XmbToken.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &XmbTokenApprovalIterator{contract: _XmbToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_XmbToken *XmbTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *XmbTokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _XmbToken.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmbTokenApproval)
				if err := _XmbToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// XmbTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the XmbToken contract.
type XmbTokenTransferIterator struct {
	Event *XmbTokenTransfer // Event containing the contract specifics and raw log

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
func (it *XmbTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmbTokenTransfer)
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
		it.Event = new(XmbTokenTransfer)
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
func (it *XmbTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmbTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmbTokenTransfer represents a Transfer event raised by the XmbToken contract.
type XmbTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_XmbToken *XmbTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*XmbTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _XmbToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &XmbTokenTransferIterator{contract: _XmbToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_XmbToken *XmbTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *XmbTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _XmbToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmbTokenTransfer)
				if err := _XmbToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
