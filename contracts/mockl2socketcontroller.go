// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ISocketFees is an auto generated low-level Go binding around an user-defined struct.
type ISocketFees struct {
	TransmissionFees *big.Int
	ExecutionFee     *big.Int
	SwitchboardFees  *big.Int
}

// MockL2SocketControllerMetaData contains all meta data concerning the MockL2SocketController contract.
var MockL2SocketControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"}],\"name\":\"ExecutionSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"localChainSlug\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"localPlug\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dstChainSlug\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstPlug\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minMsgGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionParams\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transmissionParams\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"transmissionFees\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"executionFee\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"switchboardFees\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structISocket.Fees\",\"name\":\"fees\",\"type\":\"tuple\"}],\"name\":\"MessageOutbound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"}],\"name\":\"SocketMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"connecter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"TokensMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"connector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_msgId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"finalizeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"burnAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"withdrawFromAppChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516104a43803806104a483398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610411806100936000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806321d7667c146100515780634518387314610066578063ecc7042814610079578063fc0c546a1461008f575b600080fd5b61006461005f3660046103be565b6100ba565b005b6100646100743660046103be565b6101a5565b6000546040519081526020015b60405180910390f35b6001546100a2906001600160a01b031681565b6040516001600160a01b039091168152602001610086565b6001546040516340c10f1960e01b81526001600160a01b03868116600483015260248201869052909116906340c10f1990604401600060405180830381600087803b15801561010857600080fd5b505af115801561011c573d6000803e3d6000fd5b5050604080516001600160a01b038086168252881660208201529081018690527fbf67ec129007be07f346d9d2933215293a6612ce7ff17d6b479a5a5cae72890c9250606001905060405180910390a16040518281527fdc29884a71d2bb98d3c53dc09718be05c7bfd142b7773a5c5cf2517629290ac09060200160405180910390a150505050565b600154604051632770a7eb60e21b8152336004820152602481018590526001600160a01b0390911690639dc29fac90604401600060405180830381600087803b1580156101f157600080fd5b505af1158015610205573d6000803e3d6000fd5b5050604080516001600160a01b0385811682523360208301528816818301526060810187905290517fedf7bea45e16025d7f82902171a24376f5f3a2c06d9d8c2be4d41bbc7292f74a9350908190036080019150a17f45d70e73f9a26cea18d4009d1ad25ffbdc4114ce9d8ceb15959ae9eadb0a3a2c60008060008061028a60005490565b604080516060808201835260008083526020808401828152848601838152865163ffffffff9c8d1681526001600160a01b039b8c169381019390935298909a16818601529590971690850152608084019290925260a0830185905260c0830185905260e083018590526101806101008401819052830194909452516fffffffffffffffffffffffffffffffff90811661012083015293518416610140820152905190921661016083015251908190036101a00190a17fc419bcb5e6804701c0d077263cd1a3a021af5c35e88c64784f6981388248217d848461036b60005490565b604080516001600160a01b03909416845260208401929092529082015260600160405180910390a150506000805460010190555050565b80356001600160a01b03811681146103b957600080fd5b919050565b600080600080608085870312156103d457600080fd5b6103dd856103a2565b935060208501359250604085013591506103f9606086016103a2565b90509295919450925056fea164736f6c634300080f000a",
}

// MockL2SocketControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use MockL2SocketControllerMetaData.ABI instead.
var MockL2SocketControllerABI = MockL2SocketControllerMetaData.ABI

// MockL2SocketControllerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockL2SocketControllerMetaData.Bin instead.
var MockL2SocketControllerBin = MockL2SocketControllerMetaData.Bin

// DeployMockL2SocketController deploys a new Ethereum contract, binding an instance of MockL2SocketController to it.
func DeployMockL2SocketController(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *MockL2SocketController, error) {
	parsed, err := MockL2SocketControllerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockL2SocketControllerBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockL2SocketController{MockL2SocketControllerCaller: MockL2SocketControllerCaller{contract: contract}, MockL2SocketControllerTransactor: MockL2SocketControllerTransactor{contract: contract}, MockL2SocketControllerFilterer: MockL2SocketControllerFilterer{contract: contract}}, nil
}

// MockL2SocketController is an auto generated Go binding around an Ethereum contract.
type MockL2SocketController struct {
	MockL2SocketControllerCaller     // Read-only binding to the contract
	MockL2SocketControllerTransactor // Write-only binding to the contract
	MockL2SocketControllerFilterer   // Log filterer for contract events
}

// MockL2SocketControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockL2SocketControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockL2SocketControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockL2SocketControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockL2SocketControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockL2SocketControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockL2SocketControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockL2SocketControllerSession struct {
	Contract     *MockL2SocketController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MockL2SocketControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockL2SocketControllerCallerSession struct {
	Contract *MockL2SocketControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MockL2SocketControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockL2SocketControllerTransactorSession struct {
	Contract     *MockL2SocketControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MockL2SocketControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockL2SocketControllerRaw struct {
	Contract *MockL2SocketController // Generic contract binding to access the raw methods on
}

// MockL2SocketControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockL2SocketControllerCallerRaw struct {
	Contract *MockL2SocketControllerCaller // Generic read-only contract binding to access the raw methods on
}

// MockL2SocketControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockL2SocketControllerTransactorRaw struct {
	Contract *MockL2SocketControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockL2SocketController creates a new instance of MockL2SocketController, bound to a specific deployed contract.
func NewMockL2SocketController(address common.Address, backend bind.ContractBackend) (*MockL2SocketController, error) {
	contract, err := bindMockL2SocketController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockL2SocketController{MockL2SocketControllerCaller: MockL2SocketControllerCaller{contract: contract}, MockL2SocketControllerTransactor: MockL2SocketControllerTransactor{contract: contract}, MockL2SocketControllerFilterer: MockL2SocketControllerFilterer{contract: contract}}, nil
}

// NewMockL2SocketControllerCaller creates a new read-only instance of MockL2SocketController, bound to a specific deployed contract.
func NewMockL2SocketControllerCaller(address common.Address, caller bind.ContractCaller) (*MockL2SocketControllerCaller, error) {
	contract, err := bindMockL2SocketController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockL2SocketControllerCaller{contract: contract}, nil
}

// NewMockL2SocketControllerTransactor creates a new write-only instance of MockL2SocketController, bound to a specific deployed contract.
func NewMockL2SocketControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*MockL2SocketControllerTransactor, error) {
	contract, err := bindMockL2SocketController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockL2SocketControllerTransactor{contract: contract}, nil
}

// NewMockL2SocketControllerFilterer creates a new log filterer instance of MockL2SocketController, bound to a specific deployed contract.
func NewMockL2SocketControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*MockL2SocketControllerFilterer, error) {
	contract, err := bindMockL2SocketController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockL2SocketControllerFilterer{contract: contract}, nil
}

// bindMockL2SocketController binds a generic wrapper to an already deployed contract.
func bindMockL2SocketController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockL2SocketControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockL2SocketController *MockL2SocketControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL2SocketController.Contract.MockL2SocketControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockL2SocketController *MockL2SocketControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL2SocketController.Contract.MockL2SocketControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockL2SocketController *MockL2SocketControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL2SocketController.Contract.MockL2SocketControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockL2SocketController *MockL2SocketControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL2SocketController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockL2SocketController *MockL2SocketControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL2SocketController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockL2SocketController *MockL2SocketControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL2SocketController.Contract.contract.Transact(opts, method, params...)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(bytes32)
func (_MockL2SocketController *MockL2SocketControllerCaller) MessageNonce(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockL2SocketController.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(bytes32)
func (_MockL2SocketController *MockL2SocketControllerSession) MessageNonce() ([32]byte, error) {
	return _MockL2SocketController.Contract.MessageNonce(&_MockL2SocketController.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(bytes32)
func (_MockL2SocketController *MockL2SocketControllerCallerSession) MessageNonce() ([32]byte, error) {
	return _MockL2SocketController.Contract.MessageNonce(&_MockL2SocketController.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MockL2SocketController *MockL2SocketControllerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockL2SocketController.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MockL2SocketController *MockL2SocketControllerSession) Token() (common.Address, error) {
	return _MockL2SocketController.Contract.Token(&_MockL2SocketController.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MockL2SocketController *MockL2SocketControllerCallerSession) Token() (common.Address, error) {
	return _MockL2SocketController.Contract.Token(&_MockL2SocketController.CallOpts)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x21d7667c.
//
// Solidity: function finalizeDeposit(address _to, uint256 _amount, bytes32 _msgId, address connector) returns()
func (_MockL2SocketController *MockL2SocketControllerTransactor) FinalizeDeposit(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _msgId [32]byte, connector common.Address) (*types.Transaction, error) {
	return _MockL2SocketController.contract.Transact(opts, "finalizeDeposit", _to, _amount, _msgId, connector)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x21d7667c.
//
// Solidity: function finalizeDeposit(address _to, uint256 _amount, bytes32 _msgId, address connector) returns()
func (_MockL2SocketController *MockL2SocketControllerSession) FinalizeDeposit(_to common.Address, _amount *big.Int, _msgId [32]byte, connector common.Address) (*types.Transaction, error) {
	return _MockL2SocketController.Contract.FinalizeDeposit(&_MockL2SocketController.TransactOpts, _to, _amount, _msgId, connector)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x21d7667c.
//
// Solidity: function finalizeDeposit(address _to, uint256 _amount, bytes32 _msgId, address connector) returns()
func (_MockL2SocketController *MockL2SocketControllerTransactorSession) FinalizeDeposit(_to common.Address, _amount *big.Int, _msgId [32]byte, connector common.Address) (*types.Transaction, error) {
	return _MockL2SocketController.Contract.FinalizeDeposit(&_MockL2SocketController.TransactOpts, _to, _amount, _msgId, connector)
}

// WithdrawFromAppChain is a paid mutator transaction binding the contract method 0x45183873.
//
// Solidity: function withdrawFromAppChain(address receiver_, uint256 burnAmount_, uint256 , address connector) returns()
func (_MockL2SocketController *MockL2SocketControllerTransactor) WithdrawFromAppChain(opts *bind.TransactOpts, receiver_ common.Address, burnAmount_ *big.Int, arg2 *big.Int, connector common.Address) (*types.Transaction, error) {
	return _MockL2SocketController.contract.Transact(opts, "withdrawFromAppChain", receiver_, burnAmount_, arg2, connector)
}

// WithdrawFromAppChain is a paid mutator transaction binding the contract method 0x45183873.
//
// Solidity: function withdrawFromAppChain(address receiver_, uint256 burnAmount_, uint256 , address connector) returns()
func (_MockL2SocketController *MockL2SocketControllerSession) WithdrawFromAppChain(receiver_ common.Address, burnAmount_ *big.Int, arg2 *big.Int, connector common.Address) (*types.Transaction, error) {
	return _MockL2SocketController.Contract.WithdrawFromAppChain(&_MockL2SocketController.TransactOpts, receiver_, burnAmount_, arg2, connector)
}

// WithdrawFromAppChain is a paid mutator transaction binding the contract method 0x45183873.
//
// Solidity: function withdrawFromAppChain(address receiver_, uint256 burnAmount_, uint256 , address connector) returns()
func (_MockL2SocketController *MockL2SocketControllerTransactorSession) WithdrawFromAppChain(receiver_ common.Address, burnAmount_ *big.Int, arg2 *big.Int, connector common.Address) (*types.Transaction, error) {
	return _MockL2SocketController.Contract.WithdrawFromAppChain(&_MockL2SocketController.TransactOpts, receiver_, burnAmount_, arg2, connector)
}

// MockL2SocketControllerExecutionSuccessIterator is returned from FilterExecutionSuccess and is used to iterate over the raw logs and unpacked data for ExecutionSuccess events raised by the MockL2SocketController contract.
type MockL2SocketControllerExecutionSuccessIterator struct {
	Event *MockL2SocketControllerExecutionSuccess // Event containing the contract specifics and raw log

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
func (it *MockL2SocketControllerExecutionSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL2SocketControllerExecutionSuccess)
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
		it.Event = new(MockL2SocketControllerExecutionSuccess)
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
func (it *MockL2SocketControllerExecutionSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL2SocketControllerExecutionSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL2SocketControllerExecutionSuccess represents a ExecutionSuccess event raised by the MockL2SocketController contract.
type MockL2SocketControllerExecutionSuccess struct {
	MsgId [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExecutionSuccess is a free log retrieval operation binding the contract event 0xdc29884a71d2bb98d3c53dc09718be05c7bfd142b7773a5c5cf2517629290ac0.
//
// Solidity: event ExecutionSuccess(bytes32 msgId)
func (_MockL2SocketController *MockL2SocketControllerFilterer) FilterExecutionSuccess(opts *bind.FilterOpts) (*MockL2SocketControllerExecutionSuccessIterator, error) {

	logs, sub, err := _MockL2SocketController.contract.FilterLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return &MockL2SocketControllerExecutionSuccessIterator{contract: _MockL2SocketController.contract, event: "ExecutionSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionSuccess is a free log subscription operation binding the contract event 0xdc29884a71d2bb98d3c53dc09718be05c7bfd142b7773a5c5cf2517629290ac0.
//
// Solidity: event ExecutionSuccess(bytes32 msgId)
func (_MockL2SocketController *MockL2SocketControllerFilterer) WatchExecutionSuccess(opts *bind.WatchOpts, sink chan<- *MockL2SocketControllerExecutionSuccess) (event.Subscription, error) {

	logs, sub, err := _MockL2SocketController.contract.WatchLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL2SocketControllerExecutionSuccess)
				if err := _MockL2SocketController.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
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

// ParseExecutionSuccess is a log parse operation binding the contract event 0xdc29884a71d2bb98d3c53dc09718be05c7bfd142b7773a5c5cf2517629290ac0.
//
// Solidity: event ExecutionSuccess(bytes32 msgId)
func (_MockL2SocketController *MockL2SocketControllerFilterer) ParseExecutionSuccess(log types.Log) (*MockL2SocketControllerExecutionSuccess, error) {
	event := new(MockL2SocketControllerExecutionSuccess)
	if err := _MockL2SocketController.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL2SocketControllerMessageOutboundIterator is returned from FilterMessageOutbound and is used to iterate over the raw logs and unpacked data for MessageOutbound events raised by the MockL2SocketController contract.
type MockL2SocketControllerMessageOutboundIterator struct {
	Event *MockL2SocketControllerMessageOutbound // Event containing the contract specifics and raw log

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
func (it *MockL2SocketControllerMessageOutboundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL2SocketControllerMessageOutbound)
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
		it.Event = new(MockL2SocketControllerMessageOutbound)
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
func (it *MockL2SocketControllerMessageOutboundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL2SocketControllerMessageOutboundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL2SocketControllerMessageOutbound represents a MessageOutbound event raised by the MockL2SocketController contract.
type MockL2SocketControllerMessageOutbound struct {
	LocalChainSlug     uint32
	LocalPlug          common.Address
	DstChainSlug       uint32
	DstPlug            common.Address
	MsgId              [32]byte
	MinMsgGasLimit     *big.Int
	ExecutionParams    [32]byte
	TransmissionParams [32]byte
	Payload            []byte
	Fees               ISocketFees
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageOutbound is a free log retrieval operation binding the contract event 0x45d70e73f9a26cea18d4009d1ad25ffbdc4114ce9d8ceb15959ae9eadb0a3a2c.
//
// Solidity: event MessageOutbound(uint32 localChainSlug, address localPlug, uint32 dstChainSlug, address dstPlug, bytes32 msgId, uint256 minMsgGasLimit, bytes32 executionParams, bytes32 transmissionParams, bytes payload, (uint128,uint128,uint128) fees)
func (_MockL2SocketController *MockL2SocketControllerFilterer) FilterMessageOutbound(opts *bind.FilterOpts) (*MockL2SocketControllerMessageOutboundIterator, error) {

	logs, sub, err := _MockL2SocketController.contract.FilterLogs(opts, "MessageOutbound")
	if err != nil {
		return nil, err
	}
	return &MockL2SocketControllerMessageOutboundIterator{contract: _MockL2SocketController.contract, event: "MessageOutbound", logs: logs, sub: sub}, nil
}

// WatchMessageOutbound is a free log subscription operation binding the contract event 0x45d70e73f9a26cea18d4009d1ad25ffbdc4114ce9d8ceb15959ae9eadb0a3a2c.
//
// Solidity: event MessageOutbound(uint32 localChainSlug, address localPlug, uint32 dstChainSlug, address dstPlug, bytes32 msgId, uint256 minMsgGasLimit, bytes32 executionParams, bytes32 transmissionParams, bytes payload, (uint128,uint128,uint128) fees)
func (_MockL2SocketController *MockL2SocketControllerFilterer) WatchMessageOutbound(opts *bind.WatchOpts, sink chan<- *MockL2SocketControllerMessageOutbound) (event.Subscription, error) {

	logs, sub, err := _MockL2SocketController.contract.WatchLogs(opts, "MessageOutbound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL2SocketControllerMessageOutbound)
				if err := _MockL2SocketController.contract.UnpackLog(event, "MessageOutbound", log); err != nil {
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

// ParseMessageOutbound is a log parse operation binding the contract event 0x45d70e73f9a26cea18d4009d1ad25ffbdc4114ce9d8ceb15959ae9eadb0a3a2c.
//
// Solidity: event MessageOutbound(uint32 localChainSlug, address localPlug, uint32 dstChainSlug, address dstPlug, bytes32 msgId, uint256 minMsgGasLimit, bytes32 executionParams, bytes32 transmissionParams, bytes payload, (uint128,uint128,uint128) fees)
func (_MockL2SocketController *MockL2SocketControllerFilterer) ParseMessageOutbound(log types.Log) (*MockL2SocketControllerMessageOutbound, error) {
	event := new(MockL2SocketControllerMessageOutbound)
	if err := _MockL2SocketController.contract.UnpackLog(event, "MessageOutbound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL2SocketControllerSocketMessageIterator is returned from FilterSocketMessage and is used to iterate over the raw logs and unpacked data for SocketMessage events raised by the MockL2SocketController contract.
type MockL2SocketControllerSocketMessageIterator struct {
	Event *MockL2SocketControllerSocketMessage // Event containing the contract specifics and raw log

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
func (it *MockL2SocketControllerSocketMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL2SocketControllerSocketMessage)
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
		it.Event = new(MockL2SocketControllerSocketMessage)
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
func (it *MockL2SocketControllerSocketMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL2SocketControllerSocketMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL2SocketControllerSocketMessage represents a SocketMessage event raised by the MockL2SocketController contract.
type MockL2SocketControllerSocketMessage struct {
	To     common.Address
	Amount *big.Int
	MsgId  [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSocketMessage is a free log retrieval operation binding the contract event 0xc419bcb5e6804701c0d077263cd1a3a021af5c35e88c64784f6981388248217d.
//
// Solidity: event SocketMessage(address to, uint256 amount, bytes32 msgId)
func (_MockL2SocketController *MockL2SocketControllerFilterer) FilterSocketMessage(opts *bind.FilterOpts) (*MockL2SocketControllerSocketMessageIterator, error) {

	logs, sub, err := _MockL2SocketController.contract.FilterLogs(opts, "SocketMessage")
	if err != nil {
		return nil, err
	}
	return &MockL2SocketControllerSocketMessageIterator{contract: _MockL2SocketController.contract, event: "SocketMessage", logs: logs, sub: sub}, nil
}

// WatchSocketMessage is a free log subscription operation binding the contract event 0xc419bcb5e6804701c0d077263cd1a3a021af5c35e88c64784f6981388248217d.
//
// Solidity: event SocketMessage(address to, uint256 amount, bytes32 msgId)
func (_MockL2SocketController *MockL2SocketControllerFilterer) WatchSocketMessage(opts *bind.WatchOpts, sink chan<- *MockL2SocketControllerSocketMessage) (event.Subscription, error) {

	logs, sub, err := _MockL2SocketController.contract.WatchLogs(opts, "SocketMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL2SocketControllerSocketMessage)
				if err := _MockL2SocketController.contract.UnpackLog(event, "SocketMessage", log); err != nil {
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

// ParseSocketMessage is a log parse operation binding the contract event 0xc419bcb5e6804701c0d077263cd1a3a021af5c35e88c64784f6981388248217d.
//
// Solidity: event SocketMessage(address to, uint256 amount, bytes32 msgId)
func (_MockL2SocketController *MockL2SocketControllerFilterer) ParseSocketMessage(log types.Log) (*MockL2SocketControllerSocketMessage, error) {
	event := new(MockL2SocketControllerSocketMessage)
	if err := _MockL2SocketController.contract.UnpackLog(event, "SocketMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL2SocketControllerTokensMintedIterator is returned from FilterTokensMinted and is used to iterate over the raw logs and unpacked data for TokensMinted events raised by the MockL2SocketController contract.
type MockL2SocketControllerTokensMintedIterator struct {
	Event *MockL2SocketControllerTokensMinted // Event containing the contract specifics and raw log

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
func (it *MockL2SocketControllerTokensMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL2SocketControllerTokensMinted)
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
		it.Event = new(MockL2SocketControllerTokensMinted)
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
func (it *MockL2SocketControllerTokensMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL2SocketControllerTokensMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL2SocketControllerTokensMinted represents a TokensMinted event raised by the MockL2SocketController contract.
type MockL2SocketControllerTokensMinted struct {
	Connecter  common.Address
	Receiver   common.Address
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokensMinted is a free log retrieval operation binding the contract event 0xbf67ec129007be07f346d9d2933215293a6612ce7ff17d6b479a5a5cae72890c.
//
// Solidity: event TokensMinted(address connecter, address receiver, uint256 mintAmount)
func (_MockL2SocketController *MockL2SocketControllerFilterer) FilterTokensMinted(opts *bind.FilterOpts) (*MockL2SocketControllerTokensMintedIterator, error) {

	logs, sub, err := _MockL2SocketController.contract.FilterLogs(opts, "TokensMinted")
	if err != nil {
		return nil, err
	}
	return &MockL2SocketControllerTokensMintedIterator{contract: _MockL2SocketController.contract, event: "TokensMinted", logs: logs, sub: sub}, nil
}

// WatchTokensMinted is a free log subscription operation binding the contract event 0xbf67ec129007be07f346d9d2933215293a6612ce7ff17d6b479a5a5cae72890c.
//
// Solidity: event TokensMinted(address connecter, address receiver, uint256 mintAmount)
func (_MockL2SocketController *MockL2SocketControllerFilterer) WatchTokensMinted(opts *bind.WatchOpts, sink chan<- *MockL2SocketControllerTokensMinted) (event.Subscription, error) {

	logs, sub, err := _MockL2SocketController.contract.WatchLogs(opts, "TokensMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL2SocketControllerTokensMinted)
				if err := _MockL2SocketController.contract.UnpackLog(event, "TokensMinted", log); err != nil {
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

// ParseTokensMinted is a log parse operation binding the contract event 0xbf67ec129007be07f346d9d2933215293a6612ce7ff17d6b479a5a5cae72890c.
//
// Solidity: event TokensMinted(address connecter, address receiver, uint256 mintAmount)
func (_MockL2SocketController *MockL2SocketControllerFilterer) ParseTokensMinted(log types.Log) (*MockL2SocketControllerTokensMinted, error) {
	event := new(MockL2SocketControllerTokensMinted)
	if err := _MockL2SocketController.contract.UnpackLog(event, "TokensMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL2SocketControllerTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the MockL2SocketController contract.
type MockL2SocketControllerTokensWithdrawnIterator struct {
	Event *MockL2SocketControllerTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *MockL2SocketControllerTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL2SocketControllerTokensWithdrawn)
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
		it.Event = new(MockL2SocketControllerTokensWithdrawn)
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
func (it *MockL2SocketControllerTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL2SocketControllerTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL2SocketControllerTokensWithdrawn represents a TokensWithdrawn event raised by the MockL2SocketController contract.
type MockL2SocketControllerTokensWithdrawn struct {
	Connector  common.Address
	Withdrawer common.Address
	Receiver   common.Address
	BurnAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0xedf7bea45e16025d7f82902171a24376f5f3a2c06d9d8c2be4d41bbc7292f74a.
//
// Solidity: event TokensWithdrawn(address connector, address withdrawer, address receiver, uint256 burnAmount)
func (_MockL2SocketController *MockL2SocketControllerFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts) (*MockL2SocketControllerTokensWithdrawnIterator, error) {

	logs, sub, err := _MockL2SocketController.contract.FilterLogs(opts, "TokensWithdrawn")
	if err != nil {
		return nil, err
	}
	return &MockL2SocketControllerTokensWithdrawnIterator{contract: _MockL2SocketController.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0xedf7bea45e16025d7f82902171a24376f5f3a2c06d9d8c2be4d41bbc7292f74a.
//
// Solidity: event TokensWithdrawn(address connector, address withdrawer, address receiver, uint256 burnAmount)
func (_MockL2SocketController *MockL2SocketControllerFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *MockL2SocketControllerTokensWithdrawn) (event.Subscription, error) {

	logs, sub, err := _MockL2SocketController.contract.WatchLogs(opts, "TokensWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL2SocketControllerTokensWithdrawn)
				if err := _MockL2SocketController.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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

// ParseTokensWithdrawn is a log parse operation binding the contract event 0xedf7bea45e16025d7f82902171a24376f5f3a2c06d9d8c2be4d41bbc7292f74a.
//
// Solidity: event TokensWithdrawn(address connector, address withdrawer, address receiver, uint256 burnAmount)
func (_MockL2SocketController *MockL2SocketControllerFilterer) ParseTokensWithdrawn(log types.Log) (*MockL2SocketControllerTokensWithdrawn, error) {
	event := new(MockL2SocketControllerTokensWithdrawn)
	if err := _MockL2SocketController.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
