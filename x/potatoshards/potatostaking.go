// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package potatoshards

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PotatoABI is the input ABI used to generate the binding from.
const PotatoABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"evidence\",\"type\":\"bytes[]\"}],\"name\":\"ChallengeSlashNoVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hasNotVotedSince\",\"type\":\"uint256\"}],\"name\":\"SlashNoVote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"evidence\",\"type\":\"bytes[]\"}],\"name\":\"SlashDoubleSign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"UnbondStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"FinalizeSlashNoVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shard\",\"type\":\"uint256\"}],\"name\":\"GetValidators\",\"outputs\":[{\"internalType\":\"bytes[][]\",\"name\":\"\",\"type\":\"bytes[][]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubKey\",\"type\":\"bytes[]\"}],\"name\":\"BondStake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Potato is an auto generated Go binding around an Ethereum contract.
type Potato struct {
	PotatoCaller     // Read-only binding to the contract
	PotatoTransactor // Write-only binding to the contract
	PotatoFilterer   // Log filterer for contract events
}

// PotatoCaller is an auto generated read-only Go binding around an Ethereum contract.
type PotatoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PotatoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PotatoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PotatoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PotatoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PotatoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PotatoSession struct {
	Contract     *Potato           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PotatoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PotatoCallerSession struct {
	Contract *PotatoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PotatoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PotatoTransactorSession struct {
	Contract     *PotatoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PotatoRaw is an auto generated low-level Go binding around an Ethereum contract.
type PotatoRaw struct {
	Contract *Potato // Generic contract binding to access the raw methods on
}

// PotatoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PotatoCallerRaw struct {
	Contract *PotatoCaller // Generic read-only contract binding to access the raw methods on
}

// PotatoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PotatoTransactorRaw struct {
	Contract *PotatoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPotato creates a new instance of Potato, bound to a specific deployed contract.
func NewPotato(address common.Address, backend bind.ContractBackend) (*Potato, error) {
	contract, err := bindPotato(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Potato{PotatoCaller: PotatoCaller{contract: contract}, PotatoTransactor: PotatoTransactor{contract: contract}, PotatoFilterer: PotatoFilterer{contract: contract}}, nil
}

// NewPotatoCaller creates a new read-only instance of Potato, bound to a specific deployed contract.
func NewPotatoCaller(address common.Address, caller bind.ContractCaller) (*PotatoCaller, error) {
	contract, err := bindPotato(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PotatoCaller{contract: contract}, nil
}

// NewPotatoTransactor creates a new write-only instance of Potato, bound to a specific deployed contract.
func NewPotatoTransactor(address common.Address, transactor bind.ContractTransactor) (*PotatoTransactor, error) {
	contract, err := bindPotato(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PotatoTransactor{contract: contract}, nil
}

// NewPotatoFilterer creates a new log filterer instance of Potato, bound to a specific deployed contract.
func NewPotatoFilterer(address common.Address, filterer bind.ContractFilterer) (*PotatoFilterer, error) {
	contract, err := bindPotato(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PotatoFilterer{contract: contract}, nil
}

// bindPotato binds a generic wrapper to an already deployed contract.
func bindPotato(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PotatoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Potato *PotatoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Potato.Contract.PotatoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Potato *PotatoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Potato.Contract.PotatoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Potato *PotatoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Potato.Contract.PotatoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Potato *PotatoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Potato.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Potato *PotatoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Potato.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Potato *PotatoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Potato.Contract.contract.Transact(opts, method, params...)
}

// GetValidators is a free data retrieval call binding the contract method 0xb60504b8.
//
// Solidity: function GetValidators(uint256 shard) constant returns(bytes[][])
func (_Potato *PotatoCaller) GetValidators(opts *bind.CallOpts, shard *big.Int) ([][][]byte, error) {
	var (
		ret0 = new([][][]byte)
	)
	out := ret0
	err := _Potato.contract.Call(opts, out, "GetValidators", shard)
	return *ret0, err
}

// GetValidators is a free data retrieval call binding the contract method 0xb60504b8.
//
// Solidity: function GetValidators(uint256 shard) constant returns(bytes[][])
func (_Potato *PotatoSession) GetValidators(shard *big.Int) ([][][]byte, error) {
	return _Potato.Contract.GetValidators(&_Potato.CallOpts, shard)
}

// GetValidators is a free data retrieval call binding the contract method 0xb60504b8.
//
// Solidity: function GetValidators(uint256 shard) constant returns(bytes[][])
func (_Potato *PotatoCallerSession) GetValidators(shard *big.Int) ([][][]byte, error) {
	return _Potato.Contract.GetValidators(&_Potato.CallOpts, shard)
}

// BondStake is a paid mutator transaction binding the contract method 0xf349ec48.
//
// Solidity: function BondStake(bytes[] pubKey) returns()
func (_Potato *PotatoTransactor) BondStake(opts *bind.TransactOpts, pubKey [][]byte) (*types.Transaction, error) {
	return _Potato.contract.Transact(opts, "BondStake", pubKey)
}

// BondStake is a paid mutator transaction binding the contract method 0xf349ec48.
//
// Solidity: function BondStake(bytes[] pubKey) returns()
func (_Potato *PotatoSession) BondStake(pubKey [][]byte) (*types.Transaction, error) {
	return _Potato.Contract.BondStake(&_Potato.TransactOpts, pubKey)
}

// BondStake is a paid mutator transaction binding the contract method 0xf349ec48.
//
// Solidity: function BondStake(bytes[] pubKey) returns()
func (_Potato *PotatoTransactorSession) BondStake(pubKey [][]byte) (*types.Transaction, error) {
	return _Potato.Contract.BondStake(&_Potato.TransactOpts, pubKey)
}

// ChallengeSlashNoVote is a paid mutator transaction binding the contract method 0x152d89c0.
//
// Solidity: function ChallengeSlashNoVote(address signer, bytes[] evidence) returns()
func (_Potato *PotatoTransactor) ChallengeSlashNoVote(opts *bind.TransactOpts, signer common.Address, evidence [][]byte) (*types.Transaction, error) {
	return _Potato.contract.Transact(opts, "ChallengeSlashNoVote", signer, evidence)
}

// ChallengeSlashNoVote is a paid mutator transaction binding the contract method 0x152d89c0.
//
// Solidity: function ChallengeSlashNoVote(address signer, bytes[] evidence) returns()
func (_Potato *PotatoSession) ChallengeSlashNoVote(signer common.Address, evidence [][]byte) (*types.Transaction, error) {
	return _Potato.Contract.ChallengeSlashNoVote(&_Potato.TransactOpts, signer, evidence)
}

// ChallengeSlashNoVote is a paid mutator transaction binding the contract method 0x152d89c0.
//
// Solidity: function ChallengeSlashNoVote(address signer, bytes[] evidence) returns()
func (_Potato *PotatoTransactorSession) ChallengeSlashNoVote(signer common.Address, evidence [][]byte) (*types.Transaction, error) {
	return _Potato.Contract.ChallengeSlashNoVote(&_Potato.TransactOpts, signer, evidence)
}

// FinalizeSlashNoVote is a paid mutator transaction binding the contract method 0x999f084c.
//
// Solidity: function FinalizeSlashNoVote(address signer) returns()
func (_Potato *PotatoTransactor) FinalizeSlashNoVote(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Potato.contract.Transact(opts, "FinalizeSlashNoVote", signer)
}

// FinalizeSlashNoVote is a paid mutator transaction binding the contract method 0x999f084c.
//
// Solidity: function FinalizeSlashNoVote(address signer) returns()
func (_Potato *PotatoSession) FinalizeSlashNoVote(signer common.Address) (*types.Transaction, error) {
	return _Potato.Contract.FinalizeSlashNoVote(&_Potato.TransactOpts, signer)
}

// FinalizeSlashNoVote is a paid mutator transaction binding the contract method 0x999f084c.
//
// Solidity: function FinalizeSlashNoVote(address signer) returns()
func (_Potato *PotatoTransactorSession) FinalizeSlashNoVote(signer common.Address) (*types.Transaction, error) {
	return _Potato.Contract.FinalizeSlashNoVote(&_Potato.TransactOpts, signer)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x5a06d1b2.
//
// Solidity: function SlashDoubleSign(bytes[] evidence) returns()
func (_Potato *PotatoTransactor) SlashDoubleSign(opts *bind.TransactOpts, evidence [][]byte) (*types.Transaction, error) {
	return _Potato.contract.Transact(opts, "SlashDoubleSign", evidence)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x5a06d1b2.
//
// Solidity: function SlashDoubleSign(bytes[] evidence) returns()
func (_Potato *PotatoSession) SlashDoubleSign(evidence [][]byte) (*types.Transaction, error) {
	return _Potato.Contract.SlashDoubleSign(&_Potato.TransactOpts, evidence)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x5a06d1b2.
//
// Solidity: function SlashDoubleSign(bytes[] evidence) returns()
func (_Potato *PotatoTransactorSession) SlashDoubleSign(evidence [][]byte) (*types.Transaction, error) {
	return _Potato.Contract.SlashDoubleSign(&_Potato.TransactOpts, evidence)
}

// SlashNoVote is a paid mutator transaction binding the contract method 0x2d7360a7.
//
// Solidity: function SlashNoVote(address signer, uint256 hasNotVotedSince) returns()
func (_Potato *PotatoTransactor) SlashNoVote(opts *bind.TransactOpts, signer common.Address, hasNotVotedSince *big.Int) (*types.Transaction, error) {
	return _Potato.contract.Transact(opts, "SlashNoVote", signer, hasNotVotedSince)
}

// SlashNoVote is a paid mutator transaction binding the contract method 0x2d7360a7.
//
// Solidity: function SlashNoVote(address signer, uint256 hasNotVotedSince) returns()
func (_Potato *PotatoSession) SlashNoVote(signer common.Address, hasNotVotedSince *big.Int) (*types.Transaction, error) {
	return _Potato.Contract.SlashNoVote(&_Potato.TransactOpts, signer, hasNotVotedSince)
}

// SlashNoVote is a paid mutator transaction binding the contract method 0x2d7360a7.
//
// Solidity: function SlashNoVote(address signer, uint256 hasNotVotedSince) returns()
func (_Potato *PotatoTransactorSession) SlashNoVote(signer common.Address, hasNotVotedSince *big.Int) (*types.Transaction, error) {
	return _Potato.Contract.SlashNoVote(&_Potato.TransactOpts, signer, hasNotVotedSince)
}

// UnbondStake is a paid mutator transaction binding the contract method 0x6ceb49c4.
//
// Solidity: function UnbondStake() returns()
func (_Potato *PotatoTransactor) UnbondStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Potato.contract.Transact(opts, "UnbondStake")
}

// UnbondStake is a paid mutator transaction binding the contract method 0x6ceb49c4.
//
// Solidity: function UnbondStake() returns()
func (_Potato *PotatoSession) UnbondStake() (*types.Transaction, error) {
	return _Potato.Contract.UnbondStake(&_Potato.TransactOpts)
}

// UnbondStake is a paid mutator transaction binding the contract method 0x6ceb49c4.
//
// Solidity: function UnbondStake() returns()
func (_Potato *PotatoTransactorSession) UnbondStake() (*types.Transaction, error) {
	return _Potato.Contract.UnbondStake(&_Potato.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() returns()
func (_Potato *PotatoTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Potato.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() returns()
func (_Potato *PotatoSession) Withdraw() (*types.Transaction, error) {
	return _Potato.Contract.Withdraw(&_Potato.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() returns()
func (_Potato *PotatoTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Potato.Contract.Withdraw(&_Potato.TransactOpts)
}
