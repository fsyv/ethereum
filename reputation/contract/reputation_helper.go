// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reputation

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

// EnumerableMapMetaData contains all meta data concerning the EnumerableMap contract.
var EnumerableMapMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x607b6037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122006ef361d4aa72741c6185e21c78a97efe2e099210b1b8a82d4f2c4591edce36f64736f6c637827302e382e31392d646576656c6f702e323032332e332e372b636f6d6d69742e37646436643430340058",
}

// EnumerableMapABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableMapMetaData.ABI instead.
var EnumerableMapABI = EnumerableMapMetaData.ABI

// EnumerableMapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableMapMetaData.Bin instead.
var EnumerableMapBin = EnumerableMapMetaData.Bin

// DeployEnumerableMap deploys a new Ethereum contract, binding an instance of EnumerableMap to it.
func DeployEnumerableMap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableMap, error) {
	parsed, err := EnumerableMapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableMapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableMap{EnumerableMapCaller: EnumerableMapCaller{contract: contract}, EnumerableMapTransactor: EnumerableMapTransactor{contract: contract}, EnumerableMapFilterer: EnumerableMapFilterer{contract: contract}}, nil
}

// EnumerableMap is an auto generated Go binding around an Ethereum contract.
type EnumerableMap struct {
	EnumerableMapCaller     // Read-only binding to the contract
	EnumerableMapTransactor // Write-only binding to the contract
	EnumerableMapFilterer   // Log filterer for contract events
}

// EnumerableMapCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableMapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableMapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableMapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableMapSession struct {
	Contract     *EnumerableMap    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableMapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableMapCallerSession struct {
	Contract *EnumerableMapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableMapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableMapTransactorSession struct {
	Contract     *EnumerableMapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableMapRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableMapRaw struct {
	Contract *EnumerableMap // Generic contract binding to access the raw methods on
}

// EnumerableMapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableMapCallerRaw struct {
	Contract *EnumerableMapCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableMapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableMapTransactorRaw struct {
	Contract *EnumerableMapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableMap creates a new instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMap(address common.Address, backend bind.ContractBackend) (*EnumerableMap, error) {
	contract, err := bindEnumerableMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableMap{EnumerableMapCaller: EnumerableMapCaller{contract: contract}, EnumerableMapTransactor: EnumerableMapTransactor{contract: contract}, EnumerableMapFilterer: EnumerableMapFilterer{contract: contract}}, nil
}

// NewEnumerableMapCaller creates a new read-only instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapCaller(address common.Address, caller bind.ContractCaller) (*EnumerableMapCaller, error) {
	contract, err := bindEnumerableMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapCaller{contract: contract}, nil
}

// NewEnumerableMapTransactor creates a new write-only instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableMapTransactor, error) {
	contract, err := bindEnumerableMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapTransactor{contract: contract}, nil
}

// NewEnumerableMapFilterer creates a new log filterer instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableMapFilterer, error) {
	contract, err := bindEnumerableMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapFilterer{contract: contract}, nil
}

// bindEnumerableMap binds a generic wrapper to an already deployed contract.
func bindEnumerableMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableMapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableMap *EnumerableMapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableMap.Contract.EnumerableMapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableMap *EnumerableMapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableMap.Contract.EnumerableMapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableMap *EnumerableMapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableMap.Contract.EnumerableMapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableMap *EnumerableMapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableMap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableMap *EnumerableMapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableMap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableMap *EnumerableMapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableMap.Contract.contract.Transact(opts, method, params...)
}

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x607b6037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200a4124a27bd2a1421b01b45036d6746d14efd38e3cd3275ef5c2b426cf6649a564736f6c637827302e382e31392d646576656c6f702e323032332e332e372b636f6d6d69742e37646436643430340058",
}

// EnumerableSetABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableSetMetaData.ABI instead.
var EnumerableSetABI = EnumerableSetMetaData.ABI

// EnumerableSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableSetMetaData.Bin instead.
var EnumerableSetBin = EnumerableSetMetaData.Bin

// DeployEnumerableSet deploys a new Ethereum contract, binding an instance of EnumerableSet to it.
func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// EnumerableSet is an auto generated Go binding around an Ethereum contract.
type EnumerableSet struct {
	EnumerableSetCaller     // Read-only binding to the contract
	EnumerableSetTransactor // Write-only binding to the contract
	EnumerableSetFilterer   // Log filterer for contract events
}

// EnumerableSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableSetSession struct {
	Contract     *EnumerableSet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableSetRaw struct {
	Contract *EnumerableSet // Generic contract binding to access the raw methods on
}

// EnumerableSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSet creates a new instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// NewEnumerableSetCaller creates a new read-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

// NewEnumerableSetTransactor creates a new write-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

// NewEnumerableSetFilterer creates a new log filterer instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

// bindEnumerableSet binds a generic wrapper to an already deployed contract.
func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
}

// ReputationMetaData contains all meta data concerning the Reputation contract.
var ReputationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getRepThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getReputation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"}],\"name\":\"updateReputation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ef4165f2": "getRepThreshold()",
		"9c89a0e2": "getReputation(address)",
		"6a08b511": "updateReputation(address,int256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610510806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80636a08b511146100465780639c89a0e21461005b578063ef4165f214610080575b600080fd5b6100596100543660046103e3565b610088565b005b61006e61006936600461040d565b6100f6565b60405190815260200160405180910390f35b61006e610125565b6100936000836101a1565b156100e55760006100a481846101bd565b905060008212156100c7576100c1816100bc846101d2565b6101e9565b506100d3565b6100d181836101f5565b505b6100df60008483610201565b50505050565b6100f160008383610201565b505050565b600061010281836101a1565b15610118576101126000836101bd565b92915050565b506000919050565b919050565b60008080805b610135600061021f565b81101561017e576000610148818361022a565b915050801561016b5761015b818561043e565b93508261016781610451565b9350505b508061017681610451565b91505061012b565b50806000036101905760009250505090565b61019a818361046a565b9250505090565b60006101b6836001600160a01b038416610246565b9392505050565b60006101b6836001600160a01b038416610252565b6000808212156101e55781600003610112565b5090565b60006101b6828461048c565b60006101b6828461043e565b6000610217846001600160a01b038516846102c6565b949350505050565b6000610112826102e3565b600080808061023986866102ee565b9097909650945050505050565b60006101b68383610319565b60008181526002830160205260408120548015158061027657506102768484610246565b6101b65760405162461bcd60e51b815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b65790000604482015260640160405180910390fd5b600082815260028401602052604081208290556102178484610331565b60006101128261033d565b600080806102fc8585610347565b600081815260029690960160205260409095205494959350505050565b600081815260018301602052604081205415156101b6565b60006101b68383610353565b6000610112825490565b60006101b683836103a2565b600081815260018301602052604081205461039a57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610112565b506000610112565b60008260000182815481106103b9576103b961049f565b9060005260206000200154905092915050565b80356001600160a01b038116811461012057600080fd5b600080604083850312156103f657600080fd5b6103ff836103cc565b946020939093013593505050565b60006020828403121561041f57600080fd5b6101b6826103cc565b634e487b7160e01b600052601160045260246000fd5b8082018082111561011257610112610428565b60006001820161046357610463610428565b5060010190565b60008261048757634e487b7160e01b600052601260045260246000fd5b500490565b8181038181111561011257610112610428565b634e487b7160e01b600052603260045260246000fdfea264697066735822122054cf3c06ea75389d5df5de0a10c72d3d8370647a40dbd63761292c579c31d59964736f6c637827302e382e31392d646576656c6f702e323032332e332e372b636f6d6d69742e37646436643430340058",
}

// ReputationABI is the input ABI used to generate the binding from.
// Deprecated: Use ReputationMetaData.ABI instead.
var ReputationABI = ReputationMetaData.ABI

// Deprecated: Use ReputationMetaData.Sigs instead.
// ReputationFuncSigs maps the 4-byte function signature to its string representation.
var ReputationFuncSigs = ReputationMetaData.Sigs

// ReputationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReputationMetaData.Bin instead.
var ReputationBin = ReputationMetaData.Bin

// DeployReputation deploys a new Ethereum contract, binding an instance of Reputation to it.
func DeployReputation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Reputation, error) {
	parsed, err := ReputationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReputationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reputation{ReputationCaller: ReputationCaller{contract: contract}, ReputationTransactor: ReputationTransactor{contract: contract}, ReputationFilterer: ReputationFilterer{contract: contract}}, nil
}

// Reputation is an auto generated Go binding around an Ethereum contract.
type Reputation struct {
	ReputationCaller     // Read-only binding to the contract
	ReputationTransactor // Write-only binding to the contract
	ReputationFilterer   // Log filterer for contract events
}

// ReputationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReputationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReputationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReputationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReputationSession struct {
	Contract     *Reputation       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReputationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReputationCallerSession struct {
	Contract *ReputationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ReputationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReputationTransactorSession struct {
	Contract     *ReputationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ReputationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReputationRaw struct {
	Contract *Reputation // Generic contract binding to access the raw methods on
}

// ReputationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReputationCallerRaw struct {
	Contract *ReputationCaller // Generic read-only contract binding to access the raw methods on
}

// ReputationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReputationTransactorRaw struct {
	Contract *ReputationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReputation creates a new instance of Reputation, bound to a specific deployed contract.
func NewReputation(address common.Address, backend bind.ContractBackend) (*Reputation, error) {
	contract, err := bindReputation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reputation{ReputationCaller: ReputationCaller{contract: contract}, ReputationTransactor: ReputationTransactor{contract: contract}, ReputationFilterer: ReputationFilterer{contract: contract}}, nil
}

// NewReputationCaller creates a new read-only instance of Reputation, bound to a specific deployed contract.
func NewReputationCaller(address common.Address, caller bind.ContractCaller) (*ReputationCaller, error) {
	contract, err := bindReputation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationCaller{contract: contract}, nil
}

// NewReputationTransactor creates a new write-only instance of Reputation, bound to a specific deployed contract.
func NewReputationTransactor(address common.Address, transactor bind.ContractTransactor) (*ReputationTransactor, error) {
	contract, err := bindReputation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationTransactor{contract: contract}, nil
}

// NewReputationFilterer creates a new log filterer instance of Reputation, bound to a specific deployed contract.
func NewReputationFilterer(address common.Address, filterer bind.ContractFilterer) (*ReputationFilterer, error) {
	contract, err := bindReputation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReputationFilterer{contract: contract}, nil
}

// bindReputation binds a generic wrapper to an already deployed contract.
func bindReputation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReputationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reputation *ReputationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reputation.Contract.ReputationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reputation *ReputationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputation.Contract.ReputationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reputation *ReputationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reputation.Contract.ReputationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reputation *ReputationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reputation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reputation *ReputationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reputation *ReputationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reputation.Contract.contract.Transact(opts, method, params...)
}

// GetRepThreshold is a free data retrieval call binding the contract method 0xef4165f2.
//
// Solidity: function getRepThreshold() view returns(uint256)
func (_Reputation *ReputationCaller) GetRepThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "getRepThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRepThreshold is a free data retrieval call binding the contract method 0xef4165f2.
//
// Solidity: function getRepThreshold() view returns(uint256)
func (_Reputation *ReputationSession) GetRepThreshold() (*big.Int, error) {
	return _Reputation.Contract.GetRepThreshold(&_Reputation.CallOpts)
}

// GetRepThreshold is a free data retrieval call binding the contract method 0xef4165f2.
//
// Solidity: function getRepThreshold() view returns(uint256)
func (_Reputation *ReputationCallerSession) GetRepThreshold() (*big.Int, error) {
	return _Reputation.Contract.GetRepThreshold(&_Reputation.CallOpts)
}

// GetReputation is a free data retrieval call binding the contract method 0x9c89a0e2.
//
// Solidity: function getReputation(address _address) view returns(uint256)
func (_Reputation *ReputationCaller) GetReputation(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "getReputation", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReputation is a free data retrieval call binding the contract method 0x9c89a0e2.
//
// Solidity: function getReputation(address _address) view returns(uint256)
func (_Reputation *ReputationSession) GetReputation(_address common.Address) (*big.Int, error) {
	return _Reputation.Contract.GetReputation(&_Reputation.CallOpts, _address)
}

// GetReputation is a free data retrieval call binding the contract method 0x9c89a0e2.
//
// Solidity: function getReputation(address _address) view returns(uint256)
func (_Reputation *ReputationCallerSession) GetReputation(_address common.Address) (*big.Int, error) {
	return _Reputation.Contract.GetReputation(&_Reputation.CallOpts, _address)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0x6a08b511.
//
// Solidity: function updateReputation(address _address, int256 _score) returns()
func (_Reputation *ReputationTransactor) UpdateReputation(opts *bind.TransactOpts, _address common.Address, _score *big.Int) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "updateReputation", _address, _score)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0x6a08b511.
//
// Solidity: function updateReputation(address _address, int256 _score) returns()
func (_Reputation *ReputationSession) UpdateReputation(_address common.Address, _score *big.Int) (*types.Transaction, error) {
	return _Reputation.Contract.UpdateReputation(&_Reputation.TransactOpts, _address, _score)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0x6a08b511.
//
// Solidity: function updateReputation(address _address, int256 _score) returns()
func (_Reputation *ReputationTransactorSession) UpdateReputation(_address common.Address, _score *big.Int) (*types.Transaction, error) {
	return _Reputation.Contract.UpdateReputation(&_Reputation.TransactOpts, _address, _score)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x607b6037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220df10cb82dc6935b57f6eda10534b9d22720394adbcc0e91292808dd99700a23d64736f6c637827302e382e31392d646576656c6f702e323032332e332e372b636f6d6d69742e37646436643430340058",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
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
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// SignedMathMetaData contains all meta data concerning the SignedMath contract.
var SignedMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x607b6037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220179746e94a300c32e6cd574959d1b89a919c81fd811eb71abbe2213f84c5a13264736f6c637827302e382e31392d646576656c6f702e323032332e332e372b636f6d6d69742e37646436643430340058",
}

// SignedMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SignedMathMetaData.ABI instead.
var SignedMathABI = SignedMathMetaData.ABI

// SignedMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SignedMathMetaData.Bin instead.
var SignedMathBin = SignedMathMetaData.Bin

// DeploySignedMath deploys a new Ethereum contract, binding an instance of SignedMath to it.
func DeploySignedMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SignedMath, error) {
	parsed, err := SignedMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SignedMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignedMath{SignedMathCaller: SignedMathCaller{contract: contract}, SignedMathTransactor: SignedMathTransactor{contract: contract}, SignedMathFilterer: SignedMathFilterer{contract: contract}}, nil
}

// SignedMath is an auto generated Go binding around an Ethereum contract.
type SignedMath struct {
	SignedMathCaller     // Read-only binding to the contract
	SignedMathTransactor // Write-only binding to the contract
	SignedMathFilterer   // Log filterer for contract events
}

// SignedMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignedMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignedMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignedMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignedMathSession struct {
	Contract     *SignedMath       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignedMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignedMathCallerSession struct {
	Contract *SignedMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SignedMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignedMathTransactorSession struct {
	Contract     *SignedMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SignedMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignedMathRaw struct {
	Contract *SignedMath // Generic contract binding to access the raw methods on
}

// SignedMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignedMathCallerRaw struct {
	Contract *SignedMathCaller // Generic read-only contract binding to access the raw methods on
}

// SignedMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignedMathTransactorRaw struct {
	Contract *SignedMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignedMath creates a new instance of SignedMath, bound to a specific deployed contract.
func NewSignedMath(address common.Address, backend bind.ContractBackend) (*SignedMath, error) {
	contract, err := bindSignedMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignedMath{SignedMathCaller: SignedMathCaller{contract: contract}, SignedMathTransactor: SignedMathTransactor{contract: contract}, SignedMathFilterer: SignedMathFilterer{contract: contract}}, nil
}

// NewSignedMathCaller creates a new read-only instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathCaller(address common.Address, caller bind.ContractCaller) (*SignedMathCaller, error) {
	contract, err := bindSignedMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignedMathCaller{contract: contract}, nil
}

// NewSignedMathTransactor creates a new write-only instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SignedMathTransactor, error) {
	contract, err := bindSignedMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignedMathTransactor{contract: contract}, nil
}

// NewSignedMathFilterer creates a new log filterer instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SignedMathFilterer, error) {
	contract, err := bindSignedMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignedMathFilterer{contract: contract}, nil
}

// bindSignedMath binds a generic wrapper to an already deployed contract.
func bindSignedMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignedMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedMath *SignedMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedMath.Contract.SignedMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedMath *SignedMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedMath.Contract.SignedMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedMath *SignedMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedMath.Contract.SignedMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedMath *SignedMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedMath *SignedMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedMath *SignedMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedMath.Contract.contract.Transact(opts, method, params...)
}
