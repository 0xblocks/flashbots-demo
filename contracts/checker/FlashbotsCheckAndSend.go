// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FlashbotsCheckAndSend

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

// FlashbotsCheckAndSendABI is the input ABI used to generate the binding from.
const FlashbotsCheckAndSendABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_resultMatch\",\"type\":\"bytes32\"}],\"name\":\"check32BytesAndSend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_resultMatches\",\"type\":\"bytes32[]\"}],\"name\":\"check32BytesAndSendMulti\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_resultMatch\",\"type\":\"bytes\"}],\"name\":\"checkBytesAndSend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_resultMatches\",\"type\":\"bytes[]\"}],\"name\":\"checkBytesAndSendMulti\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// FlashbotsCheckAndSendBin is the compiled bytecode used for deploying new contracts.
var FlashbotsCheckAndSendBin = "0x608060405234801561001057600080fd5b50610ce7806100206000396000f3fe60806040526004361061003f5760003560e01c806324835805146100445780633676290c1461006057806386b738be1461007c5780639af96d7e14610098575b600080fd5b61005e60048036038101906100599190610798565b6100b4565b005b61007a600480360381019061007591906108ae565b61010b565b005b61009660048036038101906100919190610731565b6101d1565b005b6100b260048036038101906100ad9190610817565b610228565b005b6100bf8383836102ee565b4173ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015610105573d6000803e3d6000fd5b50505050565b815183511461011957600080fd5b805183511461012757600080fd5b60005b83518110156101845761017784828151811061014257fe5b602002602001015184838151811061015657fe5b602002602001015184848151811061016a57fe5b60200260200101516102ee565b808060010191505061012a565b504173ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050501580156101cb573d6000803e3d6000fd5b50505050565b6101dc8383836103f3565b4173ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015610222573d6000803e3d6000fd5b50505050565b815183511461023657600080fd5b805183511461024457600080fd5b60005b83518110156102a15761029484828151811061025f57fe5b602002602001015184838151811061027357fe5b602002602001015184848151811061028757fe5b60200260200101516103f3565b8080600101915050610247565b504173ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050501580156102e8573d6000803e3d6000fd5b50505050565b600060608473ffffffffffffffffffffffffffffffffffffffff16846040516103179190610a76565b600060405180830381855afa9150503d8060008114610352576040519150601f19603f3d011682016040523d82523d6000602084013e610357565b606091505b50915091508161039c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039390610acd565b60405180910390fd5b80805190602001208380519060200120146103ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e390610aed565b60405180910390fd5b5050505050565b600060608473ffffffffffffffffffffffffffffffffffffffff168460405161041c9190610a76565b600060405180830381855afa9150503d8060008114610457576040519150601f19603f3d011682016040523d82523d6000602084013e61045c565b606091505b5091509150816104a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049890610acd565b60405180910390fd5b6020815110156104e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dd90610a8d565b60405180910390fd5b600060208201519050838114610531576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052890610aad565b60405180910390fd5b505050505050565b60008135905061054881610c83565b92915050565b600082601f83011261055f57600080fd5b813561057261056d82610b3a565b610b0d565b9150818183526020840193506020810190508385602084028201111561059757600080fd5b60005b838110156105c757816105ad8882610539565b84526020840193506020830192505060018101905061059a565b5050505092915050565b600082601f8301126105e257600080fd5b81356105f56105f082610b62565b610b0d565b9150818183526020840193506020810190508385602084028201111561061a57600080fd5b60005b8381101561064a578161063088826106c8565b84526020840193506020830192505060018101905061061d565b5050505092915050565b600082601f83011261066557600080fd5b813561067861067382610b8a565b610b0d565b9150818183526020840193506020810190508360005b838110156106be57813586016106a488826106dd565b84526020840193506020830192505060018101905061068e565b5050505092915050565b6000813590506106d781610c9a565b92915050565b600082601f8301126106ee57600080fd5b81356107016106fc82610bb2565b610b0d565b9150808252602083016020830185838301111561071d57600080fd5b610728838284610c41565b50505092915050565b60008060006060848603121561074657600080fd5b600061075486828701610539565b935050602084013567ffffffffffffffff81111561077157600080fd5b61077d868287016106dd565b925050604061078e868287016106c8565b9150509250925092565b6000806000606084860312156107ad57600080fd5b60006107bb86828701610539565b935050602084013567ffffffffffffffff8111156107d857600080fd5b6107e4868287016106dd565b925050604084013567ffffffffffffffff81111561080157600080fd5b61080d868287016106dd565b9150509250925092565b60008060006060848603121561082c57600080fd5b600084013567ffffffffffffffff81111561084657600080fd5b6108528682870161054e565b935050602084013567ffffffffffffffff81111561086f57600080fd5b61087b86828701610654565b925050604084013567ffffffffffffffff81111561089857600080fd5b6108a4868287016105d1565b9150509250925092565b6000806000606084860312156108c357600080fd5b600084013567ffffffffffffffff8111156108dd57600080fd5b6108e98682870161054e565b935050602084013567ffffffffffffffff81111561090657600080fd5b61091286828701610654565b925050604084013567ffffffffffffffff81111561092f57600080fd5b61093b86828701610654565b9150509250925092565b600061095082610bde565b61095a8185610be9565b935061096a818560208601610c50565b80840191505092915050565b6000610983601b83610bf4565b91507f726573706f6e7365206c657373207468616e20333220627974657300000000006000830152602082019050919050565b60006109c3601183610bf4565b91507f726573706f6e7365206d69736d617463680000000000000000000000000000006000830152602082019050919050565b6000610a03600883610bf4565b91507f21737563636573730000000000000000000000000000000000000000000000006000830152602082019050919050565b6000610a43601783610bf4565b91507f726573706f6e7365206279746573206d69736d617463680000000000000000006000830152602082019050919050565b6000610a828284610945565b915081905092915050565b60006020820190508181036000830152610aa681610976565b9050919050565b60006020820190508181036000830152610ac6816109b6565b9050919050565b60006020820190508181036000830152610ae6816109f6565b9050919050565b60006020820190508181036000830152610b0681610a36565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715610b3057600080fd5b8060405250919050565b600067ffffffffffffffff821115610b5157600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115610b7957600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115610ba157600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115610bc957600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b6000610c1082610c21565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b82818337600083830152505050565b60005b83811015610c6e578082015181840152602081019050610c53565b83811115610c7d576000848401525b50505050565b610c8c81610c05565b8114610c9757600080fd5b50565b610ca381610c17565b8114610cae57600080fd5b5056fea2646970667358221220af5efc663f0f97f259e0eac07059258305990995568ad1057a47d479047fb9a464736f6c634300060c0033"

// DeployFlashbotsCheckAndSend deploys a new Ethereum contract, binding an instance of FlashbotsCheckAndSend to it.
func DeployFlashbotsCheckAndSend(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FlashbotsCheckAndSend, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashbotsCheckAndSendABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FlashbotsCheckAndSendBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FlashbotsCheckAndSend{FlashbotsCheckAndSendCaller: FlashbotsCheckAndSendCaller{contract: contract}, FlashbotsCheckAndSendTransactor: FlashbotsCheckAndSendTransactor{contract: contract}, FlashbotsCheckAndSendFilterer: FlashbotsCheckAndSendFilterer{contract: contract}}, nil
}

// FlashbotsCheckAndSend is an auto generated Go binding around an Ethereum contract.
type FlashbotsCheckAndSend struct {
	FlashbotsCheckAndSendCaller     // Read-only binding to the contract
	FlashbotsCheckAndSendTransactor // Write-only binding to the contract
	FlashbotsCheckAndSendFilterer   // Log filterer for contract events
}

// FlashbotsCheckAndSendCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashbotsCheckAndSendCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashbotsCheckAndSendTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashbotsCheckAndSendTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashbotsCheckAndSendFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashbotsCheckAndSendFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashbotsCheckAndSendSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashbotsCheckAndSendSession struct {
	Contract     *FlashbotsCheckAndSend // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FlashbotsCheckAndSendCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashbotsCheckAndSendCallerSession struct {
	Contract *FlashbotsCheckAndSendCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// FlashbotsCheckAndSendTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashbotsCheckAndSendTransactorSession struct {
	Contract     *FlashbotsCheckAndSendTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// FlashbotsCheckAndSendRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashbotsCheckAndSendRaw struct {
	Contract *FlashbotsCheckAndSend // Generic contract binding to access the raw methods on
}

// FlashbotsCheckAndSendCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashbotsCheckAndSendCallerRaw struct {
	Contract *FlashbotsCheckAndSendCaller // Generic read-only contract binding to access the raw methods on
}

// FlashbotsCheckAndSendTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashbotsCheckAndSendTransactorRaw struct {
	Contract *FlashbotsCheckAndSendTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashbotsCheckAndSend creates a new instance of FlashbotsCheckAndSend, bound to a specific deployed contract.
func NewFlashbotsCheckAndSend(address common.Address, backend bind.ContractBackend) (*FlashbotsCheckAndSend, error) {
	contract, err := bindFlashbotsCheckAndSend(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlashbotsCheckAndSend{FlashbotsCheckAndSendCaller: FlashbotsCheckAndSendCaller{contract: contract}, FlashbotsCheckAndSendTransactor: FlashbotsCheckAndSendTransactor{contract: contract}, FlashbotsCheckAndSendFilterer: FlashbotsCheckAndSendFilterer{contract: contract}}, nil
}

// NewFlashbotsCheckAndSendCaller creates a new read-only instance of FlashbotsCheckAndSend, bound to a specific deployed contract.
func NewFlashbotsCheckAndSendCaller(address common.Address, caller bind.ContractCaller) (*FlashbotsCheckAndSendCaller, error) {
	contract, err := bindFlashbotsCheckAndSend(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashbotsCheckAndSendCaller{contract: contract}, nil
}

// NewFlashbotsCheckAndSendTransactor creates a new write-only instance of FlashbotsCheckAndSend, bound to a specific deployed contract.
func NewFlashbotsCheckAndSendTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashbotsCheckAndSendTransactor, error) {
	contract, err := bindFlashbotsCheckAndSend(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashbotsCheckAndSendTransactor{contract: contract}, nil
}

// NewFlashbotsCheckAndSendFilterer creates a new log filterer instance of FlashbotsCheckAndSend, bound to a specific deployed contract.
func NewFlashbotsCheckAndSendFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashbotsCheckAndSendFilterer, error) {
	contract, err := bindFlashbotsCheckAndSend(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashbotsCheckAndSendFilterer{contract: contract}, nil
}

// bindFlashbotsCheckAndSend binds a generic wrapper to an already deployed contract.
func bindFlashbotsCheckAndSend(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashbotsCheckAndSendABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashbotsCheckAndSend.Contract.FlashbotsCheckAndSendCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.Contract.FlashbotsCheckAndSendTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.Contract.FlashbotsCheckAndSendTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashbotsCheckAndSend.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.Contract.contract.Transact(opts, method, params...)
}

// Check32BytesAndSend is a paid mutator transaction binding the contract method 0x86b738be.
//
// Solidity: function check32BytesAndSend(address _target, bytes _payload, bytes32 _resultMatch) payable returns()
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendTransactor) Check32BytesAndSend(opts *bind.TransactOpts, _target common.Address, _payload []byte, _resultMatch [32]byte) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.contract.Transact(opts, "check32BytesAndSend", _target, _payload, _resultMatch)
}

// Check32BytesAndSend is a paid mutator transaction binding the contract method 0x86b738be.
//
// Solidity: function check32BytesAndSend(address _target, bytes _payload, bytes32 _resultMatch) payable returns()
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendSession) Check32BytesAndSend(_target common.Address, _payload []byte, _resultMatch [32]byte) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.Contract.Check32BytesAndSend(&_FlashbotsCheckAndSend.TransactOpts, _target, _payload, _resultMatch)
}

// Check32BytesAndSend is a paid mutator transaction binding the contract method 0x86b738be.
//
// Solidity: function check32BytesAndSend(address _target, bytes _payload, bytes32 _resultMatch) payable returns()
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendTransactorSession) Check32BytesAndSend(_target common.Address, _payload []byte, _resultMatch [32]byte) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.Contract.Check32BytesAndSend(&_FlashbotsCheckAndSend.TransactOpts, _target, _payload, _resultMatch)
}

// Check32BytesAndSendMulti is a paid mutator transaction binding the contract method 0x9af96d7e.
//
// Solidity: function check32BytesAndSendMulti(address[] _targets, bytes[] _payloads, bytes32[] _resultMatches) payable returns()
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendTransactor) Check32BytesAndSendMulti(opts *bind.TransactOpts, _targets []common.Address, _payloads [][]byte, _resultMatches [][32]byte) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.contract.Transact(opts, "check32BytesAndSendMulti", _targets, _payloads, _resultMatches)
}

// Check32BytesAndSendMulti is a paid mutator transaction binding the contract method 0x9af96d7e.
//
// Solidity: function check32BytesAndSendMulti(address[] _targets, bytes[] _payloads, bytes32[] _resultMatches) payable returns()
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendSession) Check32BytesAndSendMulti(_targets []common.Address, _payloads [][]byte, _resultMatches [][32]byte) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.Contract.Check32BytesAndSendMulti(&_FlashbotsCheckAndSend.TransactOpts, _targets, _payloads, _resultMatches)
}

// Check32BytesAndSendMulti is a paid mutator transaction binding the contract method 0x9af96d7e.
//
// Solidity: function check32BytesAndSendMulti(address[] _targets, bytes[] _payloads, bytes32[] _resultMatches) payable returns()
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendTransactorSession) Check32BytesAndSendMulti(_targets []common.Address, _payloads [][]byte, _resultMatches [][32]byte) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.Contract.Check32BytesAndSendMulti(&_FlashbotsCheckAndSend.TransactOpts, _targets, _payloads, _resultMatches)
}

// CheckBytesAndSend is a paid mutator transaction binding the contract method 0x24835805.
//
// Solidity: function checkBytesAndSend(address _target, bytes _payload, bytes _resultMatch) payable returns()
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendTransactor) CheckBytesAndSend(opts *bind.TransactOpts, _target common.Address, _payload []byte, _resultMatch []byte) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.contract.Transact(opts, "checkBytesAndSend", _target, _payload, _resultMatch)
}

// CheckBytesAndSend is a paid mutator transaction binding the contract method 0x24835805.
//
// Solidity: function checkBytesAndSend(address _target, bytes _payload, bytes _resultMatch) payable returns()
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendSession) CheckBytesAndSend(_target common.Address, _payload []byte, _resultMatch []byte) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.Contract.CheckBytesAndSend(&_FlashbotsCheckAndSend.TransactOpts, _target, _payload, _resultMatch)
}

// CheckBytesAndSend is a paid mutator transaction binding the contract method 0x24835805.
//
// Solidity: function checkBytesAndSend(address _target, bytes _payload, bytes _resultMatch) payable returns()
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendTransactorSession) CheckBytesAndSend(_target common.Address, _payload []byte, _resultMatch []byte) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.Contract.CheckBytesAndSend(&_FlashbotsCheckAndSend.TransactOpts, _target, _payload, _resultMatch)
}

// CheckBytesAndSendMulti is a paid mutator transaction binding the contract method 0x3676290c.
//
// Solidity: function checkBytesAndSendMulti(address[] _targets, bytes[] _payloads, bytes[] _resultMatches) payable returns()
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendTransactor) CheckBytesAndSendMulti(opts *bind.TransactOpts, _targets []common.Address, _payloads [][]byte, _resultMatches [][]byte) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.contract.Transact(opts, "checkBytesAndSendMulti", _targets, _payloads, _resultMatches)
}

// CheckBytesAndSendMulti is a paid mutator transaction binding the contract method 0x3676290c.
//
// Solidity: function checkBytesAndSendMulti(address[] _targets, bytes[] _payloads, bytes[] _resultMatches) payable returns()
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendSession) CheckBytesAndSendMulti(_targets []common.Address, _payloads [][]byte, _resultMatches [][]byte) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.Contract.CheckBytesAndSendMulti(&_FlashbotsCheckAndSend.TransactOpts, _targets, _payloads, _resultMatches)
}

// CheckBytesAndSendMulti is a paid mutator transaction binding the contract method 0x3676290c.
//
// Solidity: function checkBytesAndSendMulti(address[] _targets, bytes[] _payloads, bytes[] _resultMatches) payable returns()
func (_FlashbotsCheckAndSend *FlashbotsCheckAndSendTransactorSession) CheckBytesAndSendMulti(_targets []common.Address, _payloads [][]byte, _resultMatches [][]byte) (*types.Transaction, error) {
	return _FlashbotsCheckAndSend.Contract.CheckBytesAndSendMulti(&_FlashbotsCheckAndSend.TransactOpts, _targets, _payloads, _resultMatches)
}

