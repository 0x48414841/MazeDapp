// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// GameMetaData contains all meta data concerning the Game contract.
var GameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"applyWager\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"chooseWinner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currentBets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"89a26528": "applyWager()",
		"1e7a0ac6": "chooseWinner(address,address)",
		"977d25b4": "currentBets(address)",
		"27cabc3f": "scOwner()",
	},
	Bin: "0x6080604052600080546001600160a01b03191633179055610351806100256000396000f3fe60806040526004361061003f5760003560e01c80631e7a0ac61461004457806327cabc3f1461006657806389a26528146100a3578063977d25b4146100ab575b600080fd5b34801561005057600080fd5b5061006461005f3660046102c2565b6100e6565b005b34801561007257600080fd5b50600054610086906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100646101d9565b3480156100b757600080fd5b506100d86100c63660046102a0565b60016020526000908152604090205481565b60405190815260200161009a565b6000546001600160a01b031633146101505760405162461bcd60e51b815260206004820152602260248201527f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f604482015261371760f11b60648201526084015b60405180910390fd5b6001600160a01b03808216600090815260016020526040808220549285168252812054909161017e916102f5565b6001600160a01b03808516600081815260016020526040808220829055928616815282812081905591519293509183156108fc0291849190818181858888f193505050501580156101d3573d6000803e3d6000fd5b50505050565b670de0b6b3a764000034101580156101fe575033600090815260016020526040902054155b6102705760405162461bcd60e51b815260206004820152603b60248201527f446964206e6f74207761676520656e6f756768206574686572204f5220616c7260448201527f65616479206861766520616e206578697374696e6720776167657200000000006064820152608401610147565b336000908152600160205260409020349055565b80356001600160a01b038116811461029b57600080fd5b919050565b6000602082840312156102b257600080fd5b6102bb82610284565b9392505050565b600080604083850312156102d557600080fd5b6102de83610284565b91506102ec60208401610284565b90509250929050565b6000821982111561031657634e487b7160e01b600052601160045260246000fd5b50019056fea26469706673582212206209de2c098473ffbe32b34253597f0010ac93d605de4abc73435563a495129f64736f6c63430008070033",
}

// GameABI is the input ABI used to generate the binding from.
// Deprecated: Use GameMetaData.ABI instead.
var GameABI = GameMetaData.ABI

// Deprecated: Use GameMetaData.Sigs instead.
// GameFuncSigs maps the 4-byte function signature to its string representation.
var GameFuncSigs = GameMetaData.Sigs

// GameBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GameMetaData.Bin instead.
var GameBin = GameMetaData.Bin

// DeployGame deploys a new Ethereum contract, binding an instance of Game to it.
func DeployGame(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Game, error) {
	parsed, err := GameMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GameBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Game{GameCaller: GameCaller{contract: contract}, GameTransactor: GameTransactor{contract: contract}, GameFilterer: GameFilterer{contract: contract}}, nil
}

// Game is an auto generated Go binding around an Ethereum contract.
type Game struct {
	GameCaller     // Read-only binding to the contract
	GameTransactor // Write-only binding to the contract
	GameFilterer   // Log filterer for contract events
}

// GameCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameSession struct {
	Contract     *Game             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameCallerSession struct {
	Contract *GameCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameTransactorSession struct {
	Contract     *GameTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameRaw struct {
	Contract *Game // Generic contract binding to access the raw methods on
}

// GameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameCallerRaw struct {
	Contract *GameCaller // Generic read-only contract binding to access the raw methods on
}

// GameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameTransactorRaw struct {
	Contract *GameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGame creates a new instance of Game, bound to a specific deployed contract.
func NewGame(address common.Address, backend bind.ContractBackend) (*Game, error) {
	contract, err := bindGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Game{GameCaller: GameCaller{contract: contract}, GameTransactor: GameTransactor{contract: contract}, GameFilterer: GameFilterer{contract: contract}}, nil
}

// NewGameCaller creates a new read-only instance of Game, bound to a specific deployed contract.
func NewGameCaller(address common.Address, caller bind.ContractCaller) (*GameCaller, error) {
	contract, err := bindGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameCaller{contract: contract}, nil
}

// NewGameTransactor creates a new write-only instance of Game, bound to a specific deployed contract.
func NewGameTransactor(address common.Address, transactor bind.ContractTransactor) (*GameTransactor, error) {
	contract, err := bindGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameTransactor{contract: contract}, nil
}

// NewGameFilterer creates a new log filterer instance of Game, bound to a specific deployed contract.
func NewGameFilterer(address common.Address, filterer bind.ContractFilterer) (*GameFilterer, error) {
	contract, err := bindGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameFilterer{contract: contract}, nil
}

// bindGame binds a generic wrapper to an already deployed contract.
func bindGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Game *GameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Game.Contract.GameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Game *GameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.Contract.GameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Game *GameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Game.Contract.GameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Game *GameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Game.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Game *GameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Game *GameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Game.Contract.contract.Transact(opts, method, params...)
}

// CurrentBets is a free data retrieval call binding the contract method 0x977d25b4.
//
// Solidity: function currentBets(address ) view returns(uint256)
func (_Game *GameCaller) CurrentBets(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "currentBets", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentBets is a free data retrieval call binding the contract method 0x977d25b4.
//
// Solidity: function currentBets(address ) view returns(uint256)
func (_Game *GameSession) CurrentBets(arg0 common.Address) (*big.Int, error) {
	return _Game.Contract.CurrentBets(&_Game.CallOpts, arg0)
}

// CurrentBets is a free data retrieval call binding the contract method 0x977d25b4.
//
// Solidity: function currentBets(address ) view returns(uint256)
func (_Game *GameCallerSession) CurrentBets(arg0 common.Address) (*big.Int, error) {
	return _Game.Contract.CurrentBets(&_Game.CallOpts, arg0)
}

// ScOwner is a free data retrieval call binding the contract method 0x27cabc3f.
//
// Solidity: function scOwner() view returns(address)
func (_Game *GameCaller) ScOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "scOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ScOwner is a free data retrieval call binding the contract method 0x27cabc3f.
//
// Solidity: function scOwner() view returns(address)
func (_Game *GameSession) ScOwner() (common.Address, error) {
	return _Game.Contract.ScOwner(&_Game.CallOpts)
}

// ScOwner is a free data retrieval call binding the contract method 0x27cabc3f.
//
// Solidity: function scOwner() view returns(address)
func (_Game *GameCallerSession) ScOwner() (common.Address, error) {
	return _Game.Contract.ScOwner(&_Game.CallOpts)
}

// ApplyWager is a paid mutator transaction binding the contract method 0x89a26528.
//
// Solidity: function applyWager() payable returns()
func (_Game *GameTransactor) ApplyWager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "applyWager")
}

// ApplyWager is a paid mutator transaction binding the contract method 0x89a26528.
//
// Solidity: function applyWager() payable returns()
func (_Game *GameSession) ApplyWager() (*types.Transaction, error) {
	return _Game.Contract.ApplyWager(&_Game.TransactOpts)
}

// ApplyWager is a paid mutator transaction binding the contract method 0x89a26528.
//
// Solidity: function applyWager() payable returns()
func (_Game *GameTransactorSession) ApplyWager() (*types.Transaction, error) {
	return _Game.Contract.ApplyWager(&_Game.TransactOpts)
}

// ChooseWinner is a paid mutator transaction binding the contract method 0x1e7a0ac6.
//
// Solidity: function chooseWinner(address winner, address loser) returns()
func (_Game *GameTransactor) ChooseWinner(opts *bind.TransactOpts, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "chooseWinner", winner, loser)
}

// ChooseWinner is a paid mutator transaction binding the contract method 0x1e7a0ac6.
//
// Solidity: function chooseWinner(address winner, address loser) returns()
func (_Game *GameSession) ChooseWinner(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _Game.Contract.ChooseWinner(&_Game.TransactOpts, winner, loser)
}

// ChooseWinner is a paid mutator transaction binding the contract method 0x1e7a0ac6.
//
// Solidity: function chooseWinner(address winner, address loser) returns()
func (_Game *GameTransactorSession) ChooseWinner(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _Game.Contract.ChooseWinner(&_Game.TransactOpts, winner, loser)
}
