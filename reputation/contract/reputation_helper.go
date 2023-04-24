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

// RepValueMetaData contains all meta data concerning the RepValue contract.
var RepValueMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getRepThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getReputation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"registerAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"reward\",\"type\":\"int256\"}],\"name\":\"updateReputation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ef4165f2": "getRepThreshold()",
		"9c89a0e2": "getReputation(address)",
		"94cf795e": "getSigners()",
		"97c414df": "registerAccount(address)",
		"6a08b511": "updateReputation(address,int256)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061076c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80636a08b5111461005c57806394cf795e1461007157806397c414df1461008f5780639c89a0e2146100a2578063ef4165f2146100c3575b600080fd5b61006f61006a3660046105dc565b6100cb565b005b610079610143565b6040516100869190610606565b60405180910390f35b61006f61009d366004610653565b6102db565b6100b56100b0366004610653565b610307565b604051908152602001610086565b6100b5610336565b6100d66000836103b2565b1561012b5760006100e781846103ce565b9050600082121561010c576100fb826103e3565b6101059082610684565b9050610119565b6101168282610697565b90505b610125600084836103fa565b50505050565b600080821315610119576101168282610697565b5050565b6060600061014f610336565b9050600061015d6003610418565b67ffffffffffffffff811115610175576101756106aa565b60405190808252806020026020018201604052801561019e578160200160208202803683370190505b5090506000805b6101af6003610418565b8110156102295760006101c3600383610423565b509050846101d26000836103ce565b10610216578084846101e3816106c0565b9550815181106101f5576101f56106d9565b60200260200101906001600160a01b031690816001600160a01b0316815250505b5080610221816106c0565b9150506101a5565b5060008167ffffffffffffffff811115610245576102456106aa565b60405190808252806020026020018201604052801561026e578160200160208202803683370190505b50905060005b828110156102d25783818151811061028e5761028e6106d9565b60200260200101518282815181106102a8576102a86106d9565b6001600160a01b0390921660209283029190910190910152806102ca816106c0565b915050610274565b50949350505050565b6102e7600382436103fa565b506102f36000826103b2565b6103045761013f60008260016103fa565b50565b600061031381836103b2565b15610329576103236000836103ce565b92915050565b506001919050565b919050565b60008080805b6103466000610418565b81101561038f5760006103598183610423565b915050801561037c5761036c8185610697565b935082610378816106c0565b9350505b5080610387816106c0565b91505061033c565b50806000036103a15760009250505090565b6103ab81836106ef565b9250505090565b60006103c7836001600160a01b03841661043f565b9392505050565b60006103c7836001600160a01b03841661044b565b6000808212156103f65781600003610323565b5090565b6000610410846001600160a01b038516846104bf565b949350505050565b6000610323826104dc565b600080808061043286866104e7565b9097909650945050505050565b60006103c78383610512565b60008181526002830160205260408120548015158061046f575061046f848461043f565b6103c75760405162461bcd60e51b815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b65790000604482015260640160405180910390fd5b60008281526002840160205260408120829055610410848461052a565b600061032382610536565b600080806104f58585610540565b600081815260029690960160205260409095205494959350505050565b600081815260018301602052604081205415156103c7565b60006103c7838361054c565b6000610323825490565b60006103c7838361059b565b600081815260018301602052604081205461059357508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610323565b506000610323565b60008260000182815481106105b2576105b26106d9565b9060005260206000200154905092915050565b80356001600160a01b038116811461033157600080fd5b600080604083850312156105ef57600080fd5b6105f8836105c5565b946020939093013593505050565b6020808252825182820181905260009190848201906040850190845b818110156106475783516001600160a01b031683529284019291840191600101610622565b50909695505050505050565b60006020828403121561066557600080fd5b6103c7826105c5565b634e487b7160e01b600052601160045260246000fd5b818103818111156103235761032361066e565b808201808211156103235761032361066e565b634e487b7160e01b600052604160045260246000fd5b6000600182016106d2576106d261066e565b5060010190565b634e487b7160e01b600052603260045260246000fd5b60008261070c57634e487b7160e01b600052601260045260246000fd5b50049056fea264697066735822122064ee5fc7d9e8f0bb746c295158f2f670e7219c1137b2e5809a64ac9a4728a3d164736f6c637827302e382e31392d646576656c6f702e323032332e332e372b636f6d6d69742e37646436643430340058",
}

// RepValueABI is the input ABI used to generate the binding from.
// Deprecated: Use RepValueMetaData.ABI instead.
var RepValueABI = RepValueMetaData.ABI

// Deprecated: Use RepValueMetaData.Sigs instead.
// RepValueFuncSigs maps the 4-byte function signature to its string representation.
var RepValueFuncSigs = RepValueMetaData.Sigs

// RepValueBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RepValueMetaData.Bin instead.
var RepValueBin = RepValueMetaData.Bin

// DeployRepValue deploys a new Ethereum contract, binding an instance of RepValue to it.
func DeployRepValue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RepValue, error) {
	parsed, err := RepValueMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RepValueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RepValue{RepValueCaller: RepValueCaller{contract: contract}, RepValueTransactor: RepValueTransactor{contract: contract}, RepValueFilterer: RepValueFilterer{contract: contract}}, nil
}

// RepValue is an auto generated Go binding around an Ethereum contract.
type RepValue struct {
	RepValueCaller     // Read-only binding to the contract
	RepValueTransactor // Write-only binding to the contract
	RepValueFilterer   // Log filterer for contract events
}

// RepValueCaller is an auto generated read-only Go binding around an Ethereum contract.
type RepValueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepValueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RepValueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepValueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RepValueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepValueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RepValueSession struct {
	Contract     *RepValue         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RepValueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RepValueCallerSession struct {
	Contract *RepValueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RepValueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RepValueTransactorSession struct {
	Contract     *RepValueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RepValueRaw is an auto generated low-level Go binding around an Ethereum contract.
type RepValueRaw struct {
	Contract *RepValue // Generic contract binding to access the raw methods on
}

// RepValueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RepValueCallerRaw struct {
	Contract *RepValueCaller // Generic read-only contract binding to access the raw methods on
}

// RepValueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RepValueTransactorRaw struct {
	Contract *RepValueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRepValue creates a new instance of RepValue, bound to a specific deployed contract.
func NewRepValue(address common.Address, backend bind.ContractBackend) (*RepValue, error) {
	contract, err := bindRepValue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RepValue{RepValueCaller: RepValueCaller{contract: contract}, RepValueTransactor: RepValueTransactor{contract: contract}, RepValueFilterer: RepValueFilterer{contract: contract}}, nil
}

// NewRepValueCaller creates a new read-only instance of RepValue, bound to a specific deployed contract.
func NewRepValueCaller(address common.Address, caller bind.ContractCaller) (*RepValueCaller, error) {
	contract, err := bindRepValue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RepValueCaller{contract: contract}, nil
}

// NewRepValueTransactor creates a new write-only instance of RepValue, bound to a specific deployed contract.
func NewRepValueTransactor(address common.Address, transactor bind.ContractTransactor) (*RepValueTransactor, error) {
	contract, err := bindRepValue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RepValueTransactor{contract: contract}, nil
}

// NewRepValueFilterer creates a new log filterer instance of RepValue, bound to a specific deployed contract.
func NewRepValueFilterer(address common.Address, filterer bind.ContractFilterer) (*RepValueFilterer, error) {
	contract, err := bindRepValue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RepValueFilterer{contract: contract}, nil
}

// bindRepValue binds a generic wrapper to an already deployed contract.
func bindRepValue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RepValueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RepValue *RepValueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RepValue.Contract.RepValueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RepValue *RepValueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepValue.Contract.RepValueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RepValue *RepValueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RepValue.Contract.RepValueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RepValue *RepValueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RepValue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RepValue *RepValueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepValue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RepValue *RepValueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RepValue.Contract.contract.Transact(opts, method, params...)
}

// GetRepThreshold is a free data retrieval call binding the contract method 0xef4165f2.
//
// Solidity: function getRepThreshold() view returns(uint256)
func (_RepValue *RepValueCaller) GetRepThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RepValue.contract.Call(opts, &out, "getRepThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRepThreshold is a free data retrieval call binding the contract method 0xef4165f2.
//
// Solidity: function getRepThreshold() view returns(uint256)
func (_RepValue *RepValueSession) GetRepThreshold() (*big.Int, error) {
	return _RepValue.Contract.GetRepThreshold(&_RepValue.CallOpts)
}

// GetRepThreshold is a free data retrieval call binding the contract method 0xef4165f2.
//
// Solidity: function getRepThreshold() view returns(uint256)
func (_RepValue *RepValueCallerSession) GetRepThreshold() (*big.Int, error) {
	return _RepValue.Contract.GetRepThreshold(&_RepValue.CallOpts)
}

// GetReputation is a free data retrieval call binding the contract method 0x9c89a0e2.
//
// Solidity: function getReputation(address account) view returns(uint256)
func (_RepValue *RepValueCaller) GetReputation(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RepValue.contract.Call(opts, &out, "getReputation", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReputation is a free data retrieval call binding the contract method 0x9c89a0e2.
//
// Solidity: function getReputation(address account) view returns(uint256)
func (_RepValue *RepValueSession) GetReputation(account common.Address) (*big.Int, error) {
	return _RepValue.Contract.GetReputation(&_RepValue.CallOpts, account)
}

// GetReputation is a free data retrieval call binding the contract method 0x9c89a0e2.
//
// Solidity: function getReputation(address account) view returns(uint256)
func (_RepValue *RepValueCallerSession) GetReputation(account common.Address) (*big.Int, error) {
	return _RepValue.Contract.GetReputation(&_RepValue.CallOpts, account)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_RepValue *RepValueCaller) GetSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RepValue.contract.Call(opts, &out, "getSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_RepValue *RepValueSession) GetSigners() ([]common.Address, error) {
	return _RepValue.Contract.GetSigners(&_RepValue.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_RepValue *RepValueCallerSession) GetSigners() ([]common.Address, error) {
	return _RepValue.Contract.GetSigners(&_RepValue.CallOpts)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x97c414df.
//
// Solidity: function registerAccount(address account) returns()
func (_RepValue *RepValueTransactor) RegisterAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _RepValue.contract.Transact(opts, "registerAccount", account)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x97c414df.
//
// Solidity: function registerAccount(address account) returns()
func (_RepValue *RepValueSession) RegisterAccount(account common.Address) (*types.Transaction, error) {
	return _RepValue.Contract.RegisterAccount(&_RepValue.TransactOpts, account)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x97c414df.
//
// Solidity: function registerAccount(address account) returns()
func (_RepValue *RepValueTransactorSession) RegisterAccount(account common.Address) (*types.Transaction, error) {
	return _RepValue.Contract.RegisterAccount(&_RepValue.TransactOpts, account)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0x6a08b511.
//
// Solidity: function updateReputation(address account, int256 reward) returns()
func (_RepValue *RepValueTransactor) UpdateReputation(opts *bind.TransactOpts, account common.Address, reward *big.Int) (*types.Transaction, error) {
	return _RepValue.contract.Transact(opts, "updateReputation", account, reward)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0x6a08b511.
//
// Solidity: function updateReputation(address account, int256 reward) returns()
func (_RepValue *RepValueSession) UpdateReputation(account common.Address, reward *big.Int) (*types.Transaction, error) {
	return _RepValue.Contract.UpdateReputation(&_RepValue.TransactOpts, account, reward)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0x6a08b511.
//
// Solidity: function updateReputation(address account, int256 reward) returns()
func (_RepValue *RepValueTransactorSession) UpdateReputation(account common.Address, reward *big.Int) (*types.Transaction, error) {
	return _RepValue.Contract.UpdateReputation(&_RepValue.TransactOpts, account, reward)
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
