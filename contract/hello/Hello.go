// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hello

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HelloABI is the input ABI used to generate the binding from.
const HelloABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"name\":\"SetCounter\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Hello is an auto generated Go binding around an Ethereum contract.
type Hello struct {
	HelloCaller     // Read-only binding to the contract
	HelloTransactor // Write-only binding to the contract
	HelloFilterer   // Log filterer for contract events
}

// HelloCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloSession struct {
	Contract     *Hello            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloCallerSession struct {
	Contract *HelloCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HelloTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloTransactorSession struct {
	Contract     *HelloTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloRaw struct {
	Contract *Hello // Generic contract binding to access the raw methods on
}

// HelloCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloCallerRaw struct {
	Contract *HelloCaller // Generic read-only contract binding to access the raw methods on
}

// HelloTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloTransactorRaw struct {
	Contract *HelloTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHello creates a new instance of Hello, bound to a specific deployed contract.
func NewHello(address common.Address, backend bind.ContractBackend) (*Hello, error) {
	contract, err := bindHello(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hello{HelloCaller: HelloCaller{contract: contract}, HelloTransactor: HelloTransactor{contract: contract}, HelloFilterer: HelloFilterer{contract: contract}}, nil
}

// NewHelloCaller creates a new read-only instance of Hello, bound to a specific deployed contract.
func NewHelloCaller(address common.Address, caller bind.ContractCaller) (*HelloCaller, error) {
	contract, err := bindHello(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloCaller{contract: contract}, nil
}

// NewHelloTransactor creates a new write-only instance of Hello, bound to a specific deployed contract.
func NewHelloTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloTransactor, error) {
	contract, err := bindHello(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloTransactor{contract: contract}, nil
}

// NewHelloFilterer creates a new log filterer instance of Hello, bound to a specific deployed contract.
func NewHelloFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloFilterer, error) {
	contract, err := bindHello(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloFilterer{contract: contract}, nil
}

// bindHello binds a generic wrapper to an already deployed contract.
func bindHello(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hello *HelloRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hello.Contract.HelloCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hello *HelloRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hello.Contract.HelloTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hello *HelloRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hello.Contract.HelloTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hello *HelloCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hello.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hello *HelloTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hello.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hello *HelloTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hello.Contract.contract.Transact(opts, method, params...)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Hello *HelloCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hello.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Hello *HelloSession) Counter() (*big.Int, error) {
	return _Hello.Contract.Counter(&_Hello.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Hello *HelloCallerSession) Counter() (*big.Int, error) {
	return _Hello.Contract.Counter(&_Hello.CallOpts)
}

// GetCounter is a free data retrieval call binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() view returns(uint256)
func (_Hello *HelloCaller) GetCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hello.contract.Call(opts, &out, "getCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCounter is a free data retrieval call binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() view returns(uint256)
func (_Hello *HelloSession) GetCounter() (*big.Int, error) {
	return _Hello.Contract.GetCounter(&_Hello.CallOpts)
}

// GetCounter is a free data retrieval call binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() view returns(uint256)
func (_Hello *HelloCallerSession) GetCounter() (*big.Int, error) {
	return _Hello.Contract.GetCounter(&_Hello.CallOpts)
}

// SetCounter is a paid mutator transaction binding the contract method 0x8d3945e4.
//
// Solidity: function setCounter() returns()
func (_Hello *HelloTransactor) SetCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hello.contract.Transact(opts, "setCounter")
}

// SetCounter is a paid mutator transaction binding the contract method 0x8d3945e4.
//
// Solidity: function setCounter() returns()
func (_Hello *HelloSession) SetCounter() (*types.Transaction, error) {
	return _Hello.Contract.SetCounter(&_Hello.TransactOpts)
}

// SetCounter is a paid mutator transaction binding the contract method 0x8d3945e4.
//
// Solidity: function setCounter() returns()
func (_Hello *HelloTransactorSession) SetCounter() (*types.Transaction, error) {
	return _Hello.Contract.SetCounter(&_Hello.TransactOpts)
}

// HelloSetCounterIterator is returned from FilterSetCounter and is used to iterate over the raw logs and unpacked data for SetCounter events raised by the Hello contract.
type HelloSetCounterIterator struct {
	Event *HelloSetCounter // Event containing the contract specifics and raw log

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
func (it *HelloSetCounterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloSetCounter)
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
		it.Event = new(HelloSetCounter)
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
func (it *HelloSetCounterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloSetCounterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloSetCounter represents a SetCounter event raised by the Hello contract.
type HelloSetCounter struct {
	Account common.Address
	Counter *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetCounter is a free log retrieval operation binding the contract event 0x50ef2ca26f6df90c78b73ebc44c5627b62f425b94f8c6f6c53fca4edd4c97007.
//
// Solidity: event SetCounter(address indexed account, uint256 counter)
func (_Hello *HelloFilterer) FilterSetCounter(opts *bind.FilterOpts, account []common.Address) (*HelloSetCounterIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Hello.contract.FilterLogs(opts, "SetCounter", accountRule)
	if err != nil {
		return nil, err
	}
	return &HelloSetCounterIterator{contract: _Hello.contract, event: "SetCounter", logs: logs, sub: sub}, nil
}

// WatchSetCounter is a free log subscription operation binding the contract event 0x50ef2ca26f6df90c78b73ebc44c5627b62f425b94f8c6f6c53fca4edd4c97007.
//
// Solidity: event SetCounter(address indexed account, uint256 counter)
func (_Hello *HelloFilterer) WatchSetCounter(opts *bind.WatchOpts, sink chan<- *HelloSetCounter, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Hello.contract.WatchLogs(opts, "SetCounter", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloSetCounter)
				if err := _Hello.contract.UnpackLog(event, "SetCounter", log); err != nil {
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

// ParseSetCounter is a log parse operation binding the contract event 0x50ef2ca26f6df90c78b73ebc44c5627b62f425b94f8c6f6c53fca4edd4c97007.
//
// Solidity: event SetCounter(address indexed account, uint256 counter)
func (_Hello *HelloFilterer) ParseSetCounter(log types.Log) (*HelloSetCounter, error) {
	event := new(HelloSetCounter)
	if err := _Hello.contract.UnpackLog(event, "SetCounter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
