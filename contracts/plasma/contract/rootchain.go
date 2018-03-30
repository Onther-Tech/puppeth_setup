// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ByteUtilsABI is the input ABI used to generate the binding from.
const ByteUtilsABI = "[]"

// ByteUtilsBin is the compiled bytecode used for deploying new contracts.
const ByteUtilsBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a723058201d885d5c137627d5e1aa5bae6319c90e3a74fab6d9e324f27fceb54c7066414f0029`

// DeployByteUtils deploys a new Ethereum contract, binding an instance of ByteUtils to it.
func DeployByteUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ByteUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ByteUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ByteUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ByteUtils{ByteUtilsCaller: ByteUtilsCaller{contract: contract}, ByteUtilsTransactor: ByteUtilsTransactor{contract: contract}, ByteUtilsFilterer: ByteUtilsFilterer{contract: contract}}, nil
}

// ByteUtils is an auto generated Go binding around an Ethereum contract.
type ByteUtils struct {
	ByteUtilsCaller     // Read-only binding to the contract
	ByteUtilsTransactor // Write-only binding to the contract
	ByteUtilsFilterer   // Log filterer for contract events
}

// ByteUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ByteUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ByteUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ByteUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ByteUtilsSession struct {
	Contract     *ByteUtils        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ByteUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ByteUtilsCallerSession struct {
	Contract *ByteUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ByteUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ByteUtilsTransactorSession struct {
	Contract     *ByteUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ByteUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ByteUtilsRaw struct {
	Contract *ByteUtils // Generic contract binding to access the raw methods on
}

// ByteUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ByteUtilsCallerRaw struct {
	Contract *ByteUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ByteUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ByteUtilsTransactorRaw struct {
	Contract *ByteUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewByteUtils creates a new instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtils(address common.Address, backend bind.ContractBackend) (*ByteUtils, error) {
	contract, err := bindByteUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ByteUtils{ByteUtilsCaller: ByteUtilsCaller{contract: contract}, ByteUtilsTransactor: ByteUtilsTransactor{contract: contract}, ByteUtilsFilterer: ByteUtilsFilterer{contract: contract}}, nil
}

// NewByteUtilsCaller creates a new read-only instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtilsCaller(address common.Address, caller bind.ContractCaller) (*ByteUtilsCaller, error) {
	contract, err := bindByteUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ByteUtilsCaller{contract: contract}, nil
}

// NewByteUtilsTransactor creates a new write-only instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ByteUtilsTransactor, error) {
	contract, err := bindByteUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ByteUtilsTransactor{contract: contract}, nil
}

// NewByteUtilsFilterer creates a new log filterer instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ByteUtilsFilterer, error) {
	contract, err := bindByteUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ByteUtilsFilterer{contract: contract}, nil
}

// bindByteUtils binds a generic wrapper to an already deployed contract.
func bindByteUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ByteUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteUtils *ByteUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ByteUtils.Contract.ByteUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteUtils *ByteUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteUtils.Contract.ByteUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteUtils *ByteUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteUtils.Contract.ByteUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteUtils *ByteUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ByteUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteUtils *ByteUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteUtils *ByteUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteUtils.Contract.contract.Transact(opts, method, params...)
}

// ECRecoveryABI is the input ABI used to generate the binding from.
const ECRecoveryABI = "[]"

// ECRecoveryBin is the compiled bytecode used for deploying new contracts.
const ECRecoveryBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a723058208ccfd4de23b9668ca78caea965e61fb8652eb898a64bae1f1fa13d6c7d6b29020029`

// DeployECRecovery deploys a new Ethereum contract, binding an instance of ECRecovery to it.
func DeployECRecovery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECRecovery, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECRecoveryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}, ECRecoveryFilterer: ECRecoveryFilterer{contract: contract}}, nil
}

// ECRecovery is an auto generated Go binding around an Ethereum contract.
type ECRecovery struct {
	ECRecoveryCaller     // Read-only binding to the contract
	ECRecoveryTransactor // Write-only binding to the contract
	ECRecoveryFilterer   // Log filterer for contract events
}

// ECRecoveryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECRecoveryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECRecoveryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECRecoveryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoverySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECRecoverySession struct {
	Contract     *ECRecovery       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECRecoveryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECRecoveryCallerSession struct {
	Contract *ECRecoveryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ECRecoveryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECRecoveryTransactorSession struct {
	Contract     *ECRecoveryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ECRecoveryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECRecoveryRaw struct {
	Contract *ECRecovery // Generic contract binding to access the raw methods on
}

// ECRecoveryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECRecoveryCallerRaw struct {
	Contract *ECRecoveryCaller // Generic read-only contract binding to access the raw methods on
}

// ECRecoveryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECRecoveryTransactorRaw struct {
	Contract *ECRecoveryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECRecovery creates a new instance of ECRecovery, bound to a specific deployed contract.
func NewECRecovery(address common.Address, backend bind.ContractBackend) (*ECRecovery, error) {
	contract, err := bindECRecovery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}, ECRecoveryFilterer: ECRecoveryFilterer{contract: contract}}, nil
}

// NewECRecoveryCaller creates a new read-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryCaller(address common.Address, caller bind.ContractCaller) (*ECRecoveryCaller, error) {
	contract, err := bindECRecovery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryCaller{contract: contract}, nil
}

// NewECRecoveryTransactor creates a new write-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryTransactor(address common.Address, transactor bind.ContractTransactor) (*ECRecoveryTransactor, error) {
	contract, err := bindECRecovery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryTransactor{contract: contract}, nil
}

// NewECRecoveryFilterer creates a new log filterer instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryFilterer(address common.Address, filterer bind.ContractFilterer) (*ECRecoveryFilterer, error) {
	contract, err := bindECRecovery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryFilterer{contract: contract}, nil
}

// bindECRecovery binds a generic wrapper to an already deployed contract.
func bindECRecovery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.ECRecoveryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transact(opts, method, params...)
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820367f201e767701ae8cc087577d281e9971f38d913cdd0cf08739237adc1872890029`

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// MerkleABI is the input ABI used to generate the binding from.
const MerkleABI = "[]"

// MerkleBin is the compiled bytecode used for deploying new contracts.
const MerkleBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820796381c51a183ff82ee0333c01b46faa39975bab605526bb136d2256c440d59b0029`

// DeployMerkle deploys a new Ethereum contract, binding an instance of Merkle to it.
func DeployMerkle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Merkle, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Merkle{MerkleCaller: MerkleCaller{contract: contract}, MerkleTransactor: MerkleTransactor{contract: contract}, MerkleFilterer: MerkleFilterer{contract: contract}}, nil
}

// Merkle is an auto generated Go binding around an Ethereum contract.
type Merkle struct {
	MerkleCaller     // Read-only binding to the contract
	MerkleTransactor // Write-only binding to the contract
	MerkleFilterer   // Log filterer for contract events
}

// MerkleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleSession struct {
	Contract     *Merkle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleCallerSession struct {
	Contract *MerkleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MerkleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTransactorSession struct {
	Contract     *MerkleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleRaw struct {
	Contract *Merkle // Generic contract binding to access the raw methods on
}

// MerkleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleCallerRaw struct {
	Contract *MerkleCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTransactorRaw struct {
	Contract *MerkleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkle creates a new instance of Merkle, bound to a specific deployed contract.
func NewMerkle(address common.Address, backend bind.ContractBackend) (*Merkle, error) {
	contract, err := bindMerkle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Merkle{MerkleCaller: MerkleCaller{contract: contract}, MerkleTransactor: MerkleTransactor{contract: contract}, MerkleFilterer: MerkleFilterer{contract: contract}}, nil
}

// NewMerkleCaller creates a new read-only instance of Merkle, bound to a specific deployed contract.
func NewMerkleCaller(address common.Address, caller bind.ContractCaller) (*MerkleCaller, error) {
	contract, err := bindMerkle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleCaller{contract: contract}, nil
}

// NewMerkleTransactor creates a new write-only instance of Merkle, bound to a specific deployed contract.
func NewMerkleTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTransactor, error) {
	contract, err := bindMerkle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTransactor{contract: contract}, nil
}

// NewMerkleFilterer creates a new log filterer instance of Merkle, bound to a specific deployed contract.
func NewMerkleFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleFilterer, error) {
	contract, err := bindMerkle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleFilterer{contract: contract}, nil
}

// bindMerkle binds a generic wrapper to an already deployed contract.
func bindMerkle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merkle *MerkleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Merkle.Contract.MerkleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merkle *MerkleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkle.Contract.MerkleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merkle *MerkleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merkle.Contract.MerkleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merkle *MerkleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Merkle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merkle *MerkleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merkle *MerkleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merkle.Contract.contract.Transact(opts, method, params...)
}

// PriorityQueueABI is the input ABI used to generate the binding from.
const PriorityQueueABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"minChild\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"insert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"delMin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PriorityQueueBin is the compiled bytecode used for deploying new contracts.
const PriorityQueueBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a03191633600160a060020a03161790556020604051908101604052600081526100469060019081610051565b5060006002556100be565b828054828255906000526020600020908101928215610091579160200282015b82811115610091578251829060ff16905591602001919060010190610071565b5061009d9291506100a1565b5090565b6100bb91905b8082111561009d57600081556001016100a7565b90565b6105db806100cd6000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632dcdcd0c811461007157806390b5561d14610099578063b07576ac146100b1578063bda1504b146100c4578063d6362e97146100d7575b600080fd5b341561007c57600080fd5b6100876004356100ea565b60405190815260200160405180910390f35b34156100a457600080fd5b6100af6004356101ba565b005b34156100bc57600080fd5b610087610227565b34156100cf57600080fd5b6100876102ec565b34156100e257600080fd5b6100876102f2565b6000600254610114600161010860028661031590919063ffffffff16565b9063ffffffff61034b16565b11156101325761012b82600263ffffffff61031516565b90506101b5565b60016101498161010885600263ffffffff61031516565b8154811061015357fe5b600091825260209091200154600161017284600263ffffffff61031516565b8154811061017c57fe5b906000526020600020900154101561019f5761012b82600263ffffffff61031516565b61012b600161010884600263ffffffff61031516565b919050565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146101e257600080fd5b600180548082016101f38382610576565b50600091825260209091200181905560025461021690600163ffffffff61034b16565b60028190556102249061035a565b50565b6000805481903373ffffffffffffffffffffffffffffffffffffffff90811691161461025257600080fd5b600180548190811061026057fe5b9060005260206000209001549050600160025481548110151561027f57fe5b90600052602060002090015460018081548110151561029a57fe5b6000918252602090912001556002546001805490919081106102b857fe5b60009182526020822001556002546102d790600163ffffffff61046916565b6002556102e4600161047b565b8091505b5090565b60025481565b600060018081548110151561030357fe5b90600052602060002090015490505b90565b6000808315156103285760009150610344565b5082820282848281151561033857fe5b041461034057fe5b8091505b5092915050565b60008282018381101561034057fe5b60005b600061037083600263ffffffff61055f16565b111561046557600161038983600263ffffffff61055f16565b8154811061039357fe5b9060005260206000209001546001838154811015156103ae57fe5b906000526020600020900154101561044d5760016103d383600263ffffffff61055f16565b815481106103dd57fe5b90600052602060002090015490506001828154811015156103fa57fe5b600091825260209091200154600161041984600263ffffffff61055f16565b8154811061042357fe5b600091825260209091200155600180548291908490811061044057fe5b6000918252602090912001555b61045e82600263ffffffff61055f16565b915061035d565b5050565b60008282111561047557fe5b50900390565b6000805b60025461049660028561031590919063ffffffff16565b1161055a576104a4836100ea565b91506001828154811015156104b557fe5b9060005260206000209001546001848154811015156104d057fe5b90600052602060002090015411156105525760018054849081106104f057fe5b906000526020600020900154905060018281548110151561050d57fe5b90600052602060002090015460018481548110151561052857fe5b600091825260209091200155600180548291908490811061054557fe5b6000918252602090912001555b81925061047f565b505050565b600080828481151561056d57fe5b04949350505050565b81548183558181151161055a5760008381526020902061055a91810190830161031291905b808211156102e8576000815560010161059b5600a165627a7a723058200b3aa9ad9f7f87895de2430edcbcf092bd5df0e72129ace68494aa62d2ff91520029`

// DeployPriorityQueue deploys a new Ethereum contract, binding an instance of PriorityQueue to it.
func DeployPriorityQueue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriorityQueue, error) {
	parsed, err := abi.JSON(strings.NewReader(PriorityQueueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PriorityQueueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriorityQueue{PriorityQueueCaller: PriorityQueueCaller{contract: contract}, PriorityQueueTransactor: PriorityQueueTransactor{contract: contract}, PriorityQueueFilterer: PriorityQueueFilterer{contract: contract}}, nil
}

// PriorityQueue is an auto generated Go binding around an Ethereum contract.
type PriorityQueue struct {
	PriorityQueueCaller     // Read-only binding to the contract
	PriorityQueueTransactor // Write-only binding to the contract
	PriorityQueueFilterer   // Log filterer for contract events
}

// PriorityQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriorityQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriorityQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriorityQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriorityQueueSession struct {
	Contract     *PriorityQueue    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriorityQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriorityQueueCallerSession struct {
	Contract *PriorityQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PriorityQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriorityQueueTransactorSession struct {
	Contract     *PriorityQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PriorityQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriorityQueueRaw struct {
	Contract *PriorityQueue // Generic contract binding to access the raw methods on
}

// PriorityQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriorityQueueCallerRaw struct {
	Contract *PriorityQueueCaller // Generic read-only contract binding to access the raw methods on
}

// PriorityQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriorityQueueTransactorRaw struct {
	Contract *PriorityQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriorityQueue creates a new instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueue(address common.Address, backend bind.ContractBackend) (*PriorityQueue, error) {
	contract, err := bindPriorityQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriorityQueue{PriorityQueueCaller: PriorityQueueCaller{contract: contract}, PriorityQueueTransactor: PriorityQueueTransactor{contract: contract}, PriorityQueueFilterer: PriorityQueueFilterer{contract: contract}}, nil
}

// NewPriorityQueueCaller creates a new read-only instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueCaller(address common.Address, caller bind.ContractCaller) (*PriorityQueueCaller, error) {
	contract, err := bindPriorityQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueCaller{contract: contract}, nil
}

// NewPriorityQueueTransactor creates a new write-only instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*PriorityQueueTransactor, error) {
	contract, err := bindPriorityQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueTransactor{contract: contract}, nil
}

// NewPriorityQueueFilterer creates a new log filterer instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*PriorityQueueFilterer, error) {
	contract, err := bindPriorityQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueFilterer{contract: contract}, nil
}

// bindPriorityQueue binds a generic wrapper to an already deployed contract.
func bindPriorityQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriorityQueueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriorityQueue *PriorityQueueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriorityQueue.Contract.PriorityQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriorityQueue *PriorityQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriorityQueue.Contract.PriorityQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriorityQueue *PriorityQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriorityQueue.Contract.PriorityQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriorityQueue *PriorityQueueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriorityQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriorityQueue *PriorityQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriorityQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriorityQueue *PriorityQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriorityQueue.Contract.contract.Transact(opts, method, params...)
}

// CurrentSize is a free data retrieval call binding the contract method 0xbda1504b.
//
// Solidity: function currentSize() constant returns(uint256)
func (_PriorityQueue *PriorityQueueCaller) CurrentSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriorityQueue.contract.Call(opts, out, "currentSize")
	return *ret0, err
}

// CurrentSize is a free data retrieval call binding the contract method 0xbda1504b.
//
// Solidity: function currentSize() constant returns(uint256)
func (_PriorityQueue *PriorityQueueSession) CurrentSize() (*big.Int, error) {
	return _PriorityQueue.Contract.CurrentSize(&_PriorityQueue.CallOpts)
}

// CurrentSize is a free data retrieval call binding the contract method 0xbda1504b.
//
// Solidity: function currentSize() constant returns(uint256)
func (_PriorityQueue *PriorityQueueCallerSession) CurrentSize() (*big.Int, error) {
	return _PriorityQueue.Contract.CurrentSize(&_PriorityQueue.CallOpts)
}

// GetMin is a free data retrieval call binding the contract method 0xd6362e97.
//
// Solidity: function getMin() constant returns(uint256)
func (_PriorityQueue *PriorityQueueCaller) GetMin(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriorityQueue.contract.Call(opts, out, "getMin")
	return *ret0, err
}

// GetMin is a free data retrieval call binding the contract method 0xd6362e97.
//
// Solidity: function getMin() constant returns(uint256)
func (_PriorityQueue *PriorityQueueSession) GetMin() (*big.Int, error) {
	return _PriorityQueue.Contract.GetMin(&_PriorityQueue.CallOpts)
}

// GetMin is a free data retrieval call binding the contract method 0xd6362e97.
//
// Solidity: function getMin() constant returns(uint256)
func (_PriorityQueue *PriorityQueueCallerSession) GetMin() (*big.Int, error) {
	return _PriorityQueue.Contract.GetMin(&_PriorityQueue.CallOpts)
}

// MinChild is a free data retrieval call binding the contract method 0x2dcdcd0c.
//
// Solidity: function minChild(i uint256) constant returns(uint256)
func (_PriorityQueue *PriorityQueueCaller) MinChild(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriorityQueue.contract.Call(opts, out, "minChild", i)
	return *ret0, err
}

// MinChild is a free data retrieval call binding the contract method 0x2dcdcd0c.
//
// Solidity: function minChild(i uint256) constant returns(uint256)
func (_PriorityQueue *PriorityQueueSession) MinChild(i *big.Int) (*big.Int, error) {
	return _PriorityQueue.Contract.MinChild(&_PriorityQueue.CallOpts, i)
}

// MinChild is a free data retrieval call binding the contract method 0x2dcdcd0c.
//
// Solidity: function minChild(i uint256) constant returns(uint256)
func (_PriorityQueue *PriorityQueueCallerSession) MinChild(i *big.Int) (*big.Int, error) {
	return _PriorityQueue.Contract.MinChild(&_PriorityQueue.CallOpts, i)
}

// DelMin is a paid mutator transaction binding the contract method 0xb07576ac.
//
// Solidity: function delMin() returns(uint256)
func (_PriorityQueue *PriorityQueueTransactor) DelMin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriorityQueue.contract.Transact(opts, "delMin")
}

// DelMin is a paid mutator transaction binding the contract method 0xb07576ac.
//
// Solidity: function delMin() returns(uint256)
func (_PriorityQueue *PriorityQueueSession) DelMin() (*types.Transaction, error) {
	return _PriorityQueue.Contract.DelMin(&_PriorityQueue.TransactOpts)
}

// DelMin is a paid mutator transaction binding the contract method 0xb07576ac.
//
// Solidity: function delMin() returns(uint256)
func (_PriorityQueue *PriorityQueueTransactorSession) DelMin() (*types.Transaction, error) {
	return _PriorityQueue.Contract.DelMin(&_PriorityQueue.TransactOpts)
}

// Insert is a paid mutator transaction binding the contract method 0x90b5561d.
//
// Solidity: function insert(k uint256) returns()
func (_PriorityQueue *PriorityQueueTransactor) Insert(opts *bind.TransactOpts, k *big.Int) (*types.Transaction, error) {
	return _PriorityQueue.contract.Transact(opts, "insert", k)
}

// Insert is a paid mutator transaction binding the contract method 0x90b5561d.
//
// Solidity: function insert(k uint256) returns()
func (_PriorityQueue *PriorityQueueSession) Insert(k *big.Int) (*types.Transaction, error) {
	return _PriorityQueue.Contract.Insert(&_PriorityQueue.TransactOpts, k)
}

// Insert is a paid mutator transaction binding the contract method 0x90b5561d.
//
// Solidity: function insert(k uint256) returns()
func (_PriorityQueue *PriorityQueueTransactorSession) Insert(k *big.Int) (*types.Transaction, error) {
	return _PriorityQueue.Contract.Insert(&_PriorityQueue.TransactOpts, k)
}

// RLPABI is the input ABI used to generate the binding from.
const RLPABI = "[]"

// RLPBin is the compiled bytecode used for deploying new contracts.
const RLPBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a723058204627ae830bf4c4e699b1a8a0776d7edc590ba3326c4452f48f3f1b6f9fd619560029`

// DeployRLP deploys a new Ethereum contract, binding an instance of RLP to it.
func DeployRLP(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLP, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLP{RLPCaller: RLPCaller{contract: contract}, RLPTransactor: RLPTransactor{contract: contract}, RLPFilterer: RLPFilterer{contract: contract}}, nil
}

// RLP is an auto generated Go binding around an Ethereum contract.
type RLP struct {
	RLPCaller     // Read-only binding to the contract
	RLPTransactor // Write-only binding to the contract
	RLPFilterer   // Log filterer for contract events
}

// RLPCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPSession struct {
	Contract     *RLP              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPCallerSession struct {
	Contract *RLPCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RLPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPTransactorSession struct {
	Contract     *RLPTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPRaw struct {
	Contract *RLP // Generic contract binding to access the raw methods on
}

// RLPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPCallerRaw struct {
	Contract *RLPCaller // Generic read-only contract binding to access the raw methods on
}

// RLPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPTransactorRaw struct {
	Contract *RLPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLP creates a new instance of RLP, bound to a specific deployed contract.
func NewRLP(address common.Address, backend bind.ContractBackend) (*RLP, error) {
	contract, err := bindRLP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLP{RLPCaller: RLPCaller{contract: contract}, RLPTransactor: RLPTransactor{contract: contract}, RLPFilterer: RLPFilterer{contract: contract}}, nil
}

// NewRLPCaller creates a new read-only instance of RLP, bound to a specific deployed contract.
func NewRLPCaller(address common.Address, caller bind.ContractCaller) (*RLPCaller, error) {
	contract, err := bindRLP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPCaller{contract: contract}, nil
}

// NewRLPTransactor creates a new write-only instance of RLP, bound to a specific deployed contract.
func NewRLPTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPTransactor, error) {
	contract, err := bindRLP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPTransactor{contract: contract}, nil
}

// NewRLPFilterer creates a new log filterer instance of RLP, bound to a specific deployed contract.
func NewRLPFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPFilterer, error) {
	contract, err := bindRLP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPFilterer{contract: contract}, nil
}

// bindRLP binds a generic wrapper to an already deployed contract.
func bindRLP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLP *RLPRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLP.Contract.RLPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLP *RLPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLP.Contract.RLPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLP *RLPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLP.Contract.RLPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLP *RLPCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLP *RLPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLP *RLPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLP.Contract.contract.Transact(opts, method, params...)
}

// RLPEncodeABI is the input ABI used to generate the binding from.
const RLPEncodeABI = "[]"

// RLPEncodeBin is the compiled bytecode used for deploying new contracts.
const RLPEncodeBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820f207c06f9e3522b6a0e2b3fe55f7e3fb21e43a1d81d989fcb87bff53f5d566000029`

// DeployRLPEncode deploys a new Ethereum contract, binding an instance of RLPEncode to it.
func DeployRLPEncode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPEncode, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPEncodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPEncodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// RLPEncode is an auto generated Go binding around an Ethereum contract.
type RLPEncode struct {
	RLPEncodeCaller     // Read-only binding to the contract
	RLPEncodeTransactor // Write-only binding to the contract
	RLPEncodeFilterer   // Log filterer for contract events
}

// RLPEncodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPEncodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPEncodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPEncodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPEncodeSession struct {
	Contract     *RLPEncode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPEncodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPEncodeCallerSession struct {
	Contract *RLPEncodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPEncodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPEncodeTransactorSession struct {
	Contract     *RLPEncodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPEncodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPEncodeRaw struct {
	Contract *RLPEncode // Generic contract binding to access the raw methods on
}

// RLPEncodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPEncodeCallerRaw struct {
	Contract *RLPEncodeCaller // Generic read-only contract binding to access the raw methods on
}

// RLPEncodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPEncodeTransactorRaw struct {
	Contract *RLPEncodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPEncode creates a new instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncode(address common.Address, backend bind.ContractBackend) (*RLPEncode, error) {
	contract, err := bindRLPEncode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// NewRLPEncodeCaller creates a new read-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeCaller(address common.Address, caller bind.ContractCaller) (*RLPEncodeCaller, error) {
	contract, err := bindRLPEncode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeCaller{contract: contract}, nil
}

// NewRLPEncodeTransactor creates a new write-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPEncodeTransactor, error) {
	contract, err := bindRLPEncode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeTransactor{contract: contract}, nil
}

// NewRLPEncodeFilterer creates a new log filterer instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPEncodeFilterer, error) {
	contract, err := bindRLPEncode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeFilterer{contract: contract}, nil
}

// bindRLPEncode binds a generic wrapper to an already deployed contract.
func bindRLPEncode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPEncodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.RLPEncodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transact(opts, method, params...)
}

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"nextWeekOldChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exitIds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"utxoPos\",\"type\":\"uint256\"},{\"name\":\"txBytes\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"sigs\",\"type\":\"bytes\"}],\"name\":\"startExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cUtxoPos\",\"type\":\"uint256\"},{\"name\":\"eUtxoPos\",\"type\":\"uint256\"},{\"name\":\"txBytes\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"sigs\",\"type\":\"bytes\"},{\"name\":\"confirmationSig\",\"type\":\"bytes\"}],\"name\":\"challengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exits\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"utxoPos\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"childBlockInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weekOldBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"recentBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getChildChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentDepositBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"submitBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeExits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"priority\",\"type\":\"uint256\"}],\"name\":\"getExit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childChain\",\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"created_at\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"exitor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"utxoPos\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `0x6060604052341561000f57600080fd5b60048054600160a060020a03191633600160a060020a03161790556103e8600981905560055560016006819055600855610047610082565b604051809103906000f080151561005d57600080fd5b60038054600160a060020a031916600160a060020a0392909216919091179055610093565b6040516106a880620021d483390190565b61213180620000a36000396000f3006060604052600436106100e25763ffffffff60e060020a600035041663107e58f181146100e757806316b388401461010f5780631c91a6b91461012557806332773ba314610201578063342de1791461032357806338a9e0bc1461036f5780634237b5f3146103825780636f84b695146103955780637a95f1e8146103a857806385444de3146103bb578063a98c7f2c146103e9578063baa47694146103fc578063bcd5926114610412578063bf7e214f14610425578063c6ab44cd14610454578063d0e30db014610467578063e60f1ff11461046f578063f95643b114610485575b600080fd5b34156100f257600080fd5b6100fd60043561049b565b60405190815260200160405180910390f35b341561011a57600080fd5b6100fd6004356104d8565b341561013057600080fd5b6101ff600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506104ea95505050505050565b005b341561020c57600080fd5b6101ff600480359060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061091595505050505050565b341561032e57600080fd5b610339600435610ad7565b6040518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390f35b341561037a57600080fd5b6100fd610b03565b341561038d57600080fd5b6100fd610b09565b34156103a057600080fd5b6100fd610b0f565b34156103b357600080fd5b6100fd610b15565b34156103c657600080fd5b6103d1600435610b1b565b60405191825260208201526040908101905180910390f35b34156103f457600080fd5b6100fd610b35565b341561040757600080fd5b6101ff600435610b3b565b341561041d57600080fd5b6100fd610c52565b341561043057600080fd5b610438610c76565b604051600160a060020a03909116815260200160405180910390f35b341561045f57600080fd5b6100fd610c85565b6101ff61104f565b341561047a57600080fd5b610339600435611248565b341561049057600080fd5b6103d1600435611276565b6009546000906104d2906104c660016104ba868463ffffffff61128f16565b9063ffffffff6112ab16565b9063ffffffff6112c116565b92915050565b60026020526000908152604090205481565b6104f261209b565b60008060008060008060005b6105114262093a8063ffffffff6112ec16565b60085460009081526020819052604090206001015410156105a45760085460009081526020819052604090206001015415156105875760008061055560085461049b565b81526020019081526020016000206001015460001415610574576105a4565b61057f60085461049b565b60085561059f565b60085461059b90600163ffffffff6112ab16565b6008555b6104fe565b6105be600b6105b28d6112fe565b9063ffffffff61135216565b9750633b9aca00808d049750612710908d0660008981526020819052604090205491900496506127108702633b9aca0089028e03039550935061061b88600660028802018151811061060c57fe5b90602001906020020151611408565b600160a060020a031633600160a060020a031614151561063a57600080fd5b8a6040518082805190602001908083835b6020831061066a5780518252601f19909201916020918201910161064b565b6001836020036101000a038019825116818451161790925250505091909101925060409150505180910390209250826106a68a60006082611455565b6040518281526020810182805190602001908083835b602083106106db5780518252601f1990920191602091820191016106bc565b6001836020036101000a03801982511681845116179092525050509190910193506040925050505180910390209150610744838561072e8b60008151811061071f57fe5b906020019060200201516114ab565b61073e8c60038151811061071f57fe5b8d6114f2565b151561074f57600080fd5b6107618287868d63ffffffff61165b16565b151561076c57600080fd5b60085487101561079b57610794600854888e81151561078757fe5b049063ffffffff6112c116565b905061079e565b508a5b60008c815260026020526040902054156107b757600080fd5b60008c81526002602052604090819020829055600354600160a060020a0316906390b5561d9083905160e060020a63ffffffff84160281526004810191909152602401600060405180830381600087803b151561081357600080fd5b6102c65a03f1151561082457600080fd5b5050506060604051908101604052806108488a886002026006018151811061060c57fe5b600160a060020a0316815260200161086b8a886002026007018151811061071f57fe5b815260209081018e905260008381526001909152604090208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0391909116178155602082015181600101556040820151600290910155507f22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631338d604051600160a060020a03909216825260208201526040908101905180910390a1505050505050505050505050565b633b9aca0080870460009081526020818152604080832054898452600290925280832054612710948b0694909404939192908190819081908b90518082805190602001908083835b6020831061097c5780518252601f19909201916020918201910161095d565b6001836020036101000a03801982511681845116179092525050509190910192506040915050518091039020935083866040519182526020820152604090810190518091039020925083896040518281526020810182805190602001908083835b602083106109fc5780518252601f1990920191602091820191016109dd565b6001836020036101000a038019825116818451161790925250505091909101935060409250505051908190039020600086815260016020526040902054909250600160a060020a03169050610a5183896116e8565b600160a060020a03828116911614610a6857600080fd5b610a7a8288888d63ffffffff61165b16565b1515610a8557600080fd5b505050600091825250600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff191681559283018490556002928301849055998352529687209690965550505050505050565b6001602081905260009182526040909120805491810154600290910154600160a060020a039092169183565b60095481565b60085481565b60075481565b60055481565b600090815260208190526040902080546001909101549091565b60065481565b60045433600160a060020a03908116911614610b5657600080fd5b610b694262093a8063ffffffff6112ec16565b6008546000908152602081905260409020600101541015610bfc576008546000908152602081905260409020600101541515610bdf57600080610bad60085461049b565b81526020019081526020016000206001015460001415610bcc57610bfc565b610bd760085461049b565b600855610bf7565b600854610bf390600163ffffffff6112ab16565b6008555b610b56565b604080519081016040908152828252426020808401919091526005546000908152908190522081518155602082015160019091015550600954600554610c479163ffffffff6112ab16565b600555506001600655565b6000610c716006546104ba6009546005546112ec90919063ffffffff16565b905090565b600454600160a060020a031681565b600080610c906120ad565b6000805b610ca74262093a8063ffffffff6112ec16565b6008546000908152602081905260409020600101541015610d3a576008546000908152602081905260409020600101541515610d1d57600080610ceb60085461049b565b81526020019081526020016000206001015460001415610d0a57610d3a565b610d1560085461049b565b600855610d35565b600854610d3190600163ffffffff6112ab16565b6008555b610c94565b610d4d426212750063ffffffff6112ec16565b600354909450600190600090600160a060020a031663d6362e9782604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610d9d57600080fd5b6102c65a03f11515610dae57600080fd5b505050604051805190508152602001908152602001600020606060405190810160409081528254600160a060020a03168252600183015460208301526002909201549181019182529350610e0f90633b9aca0090519063ffffffff61128f16565b91505b6000828152602081905260409020600101548490108015610e955750600354600090600160a060020a031663bda1504b82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610e7857600080fd5b6102c65a03f11515610e8957600080fd5b50505060405180519050115b15611048578251600160a060020a03166108fc84602001519081150290604051600060405180830381858888f193505050501515610ed257600080fd5b600354600160a060020a031663b07576ac6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610f1a57600080fd5b6102c65a03f11515610f2b57600080fd5b50505060405180516000818152600160208190526040808320805473ffffffffffffffffffffffffffffffffffffffff19168155918201839055600291820183905592945092509085015181526020019081526020016000206000905560016000600360009054906101000a9004600160a060020a0316600160a060020a031663d6362e976000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610fe757600080fd5b6102c65a03f11515610ff857600080fd5b505050604051805190508152602001908152602001600020606060405190810160409081528254600160a060020a0316825260018301546020830152600290920154918101919091529250610e12565b5050505090565b61105761209b565b600080600060095460065410151561106e57600080fd5b6110766117c8565b9350836040518082805190602001908083835b602083106110a85780518252601f199092019160209182019101611089565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051809103902060826040518059106110e35750595b818152601f19601f830116810160200160405290506040518281526020810182805190602001908083835b6020831061112d5780518252601f19909201916020918201910161110e565b6001836020036101000a03801982511681845116179092525050509190910193506040925050505180910390209150600090505b60108110156111a95781836040519182526020820152604090810190518091039020915082836040519182526020820152604090810190519081900390209250600101611161565b604080519081016040528281524260208201526000806111c7610c52565b81526020019081526020016000206000820151815560208201516001918201556006546111fb92509063ffffffff6112ab16565b6006557fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3334604051600160a060020a03909216825260208201526040908101905180910390a150505050565b6000818152600160208190526040909120805491810154600290910154600160a060020a0390921693909250565b6000602081905290815260409020805460019091015482565b600080828481151561129d57fe5b0490508091505b5092915050565b6000828201838110156112ba57fe5b9392505050565b6000808315156112d457600091506112a4565b508282028284828115156112e457fe5b04146112ba57fe5b6000828211156112f857fe5b50900390565b6113066120cd565b6000808351915081151561132f576040805190810160405260008082526020820152925061134b565b5060208301604080519081016040528181526020810183905292505b5050919050565b61135a61209b565b6113626120e4565b600061136d856119ae565b151561137857600080fd5b836040518059106113865750595b9080825280602002602001820160405280156113bc57816020015b6113a96120cd565b8152602001906001900390816113a15790505b5092506113c8856119db565b91505b6113d482611a14565b15611400576113e282611a37565b8382815181106113ee57fe5b602090810290910101526001016113cb565b505092915050565b600080600061141684611a79565b151561142157600080fd5b61142a84611aa3565b90925090506014811461143c57600080fd5b6c01000000000000000000000000825104949350505050565b61145d61209b565b61146561209b565b6040519050601f831680820184810186838901015b8183101561149257805183526020928301920161147a565b5050848352601f01601f19166040525090509392505050565b60008060006114b984611a79565b15156114c457600080fd5b6114cd84611aa3565b9150915060208111156114df57600080fd5b806020036101000a825104949350505050565b60006114fc61209b565b61150461209b565b61150c61209b565b600080600061151961209b565b6041895181151561152657fe5b061580156115375750610104895111155b151561154257600080fd5b61154f8960006041611455565b965061155d89604180611455565b955061156c8960826041611455565b94508c8c604051918252602082015260409081019051809103902093508a6000148015611597575089155b156115c1576115a684866116e8565b600160a060020a031633600160a060020a031614975061164b565b600192506001915060008b11156115fc576115dc84866116e8565b600160a060020a03166115ef8e896116e8565b600160a060020a03161492505b60008a111561163e576116128960c36041611455565b905061161e84826116e8565b600160a060020a03166116318e886116e8565b600160a060020a03161491505b8280156116485750815b97505b5050505050505095945050505050565b60008060008084516102001461167057600080fd5b5086905060205b61020081116116da578085015192506002870615156116b0578183604051918252602082015260409081019051809103902091506116cc565b8282604051918252602082015260409081019051809103902091505b600287049650602001611677565b509390931495945050505050565b600080600080845160411461170057600093506117bf565b6020850151925060408501519150606085015160001a9050601b8160ff16101561172857601b015b8060ff16601b1415801561174057508060ff16601c14155b1561174e57600093506117bf565b6001868285856040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f115156117b357600080fd5b50506020604051035193505b50505092915050565b6117d061209b565b60006117da61209b565b6117e261209b565b6117ea61209b565b600b6040518059106117f95750595b90808252806020026020018201604052801561182f57816020015b61181c61209b565b8152602001906001900390816118145790505b50925060146040518059106118415750595b818152601f19601f830116810160200160405290509150611884600060405180591061186a5750595b818152601f19601f83011681016020016040529050611b20565b9050600093505b81518410156118c55760008285815181106118a257fe5b906020010190600160f860020a031916908160001a90535060019093019261188b565b6118ce82611b20565b9150600093505b600684101561190057808385815181106118eb57fe5b602090810290910101526001909301926118d5565b61191a611915600160a060020a033316611ba1565b611b20565b8360068151811061192757fe5b6020908102909101015261193d61191534611ba1565b8360078151811061194a57fe5b60209081029091010152818360088151811061196257fe5b60209081029091010152808360098151811061197a57fe5b602090810290910101528083600a8151811061199257fe5b602090810290910101526119a583611c5b565b94505050505090565b600080826020015115156119c557600091506119d5565b8251905060c0815160001a101591505b50919050565b6119e36120e4565b60006119ee836119ae565b15156119f957600080fd5b611a0283611c93565b83519383529092016020820152919050565b6000611a1e6120cd565b8251905080602001518151018360200151109392505050565b611a3f6120cd565b600080611a4b84611a14565b156100e25783602001519150611a6082611d12565b828452602080850182905283820190860152905061134b565b60008082602001511515611a9057600091506119d5565b8251905060c0815160001a109392505050565b6000806000806000611ab486611a79565b1515611abf57600080fd5b85519150815160001a92506080831015611adf5781945060019350611b18565b60b8831015611afd5760018660200151039350816001019450611b18565b5060b619820180600160208801510303935080820160010194505b505050915091565b611b2861209b565b611b3061209b565b82516001148015611b665750608083600081518110611b4b57fe5b016020015160f860020a900460f860020a0260f860020a9004105b15611b94576001604051805910611b7a5750595b8181526020601f909201601f1916010160405250816104d2565b6112ba83608060b7611da4565b611ba961209b565b6000805b602082108015611bd65750838260208110611bc457fe5b1a60f860020a0260f860020a90046000145b15611be657816001019150611bad565b81602003604051805910611bf75750595b818152601f19601f8301168101602001604052905092505b602082101561134b57838260208110611c2457fe5b1a60f860020a02838281518110611c3757fe5b906020010190600160f860020a031916908160001a90535060019182019101611c0f565b611c6361209b565b611c6b61209b565b611c7361209b565b611c7c84611f3b565b9150611c8b8260c060f7611da4565b949350505050565b600080600083602001511515611cac576000925061134b565b83519050805160001a91506080821015611cc9576000925061134b565b60b8821080611ce4575060c08210158015611ce4575060f882105b15611cf2576001925061134b565b60c0821015611d075760b5198201925061134b565b5060f5190192915050565b600080825160001a90506080811015611d2e57600191506119d5565b60b8811015611d4357607e19810191506119d5565b60c0811015611d6d5760b78103806020036101000a600185015104808201600101935050506119d5565b60f8811015611d825760be19810191506119d5565b60f78103806020036101000a6001850151048082016001019350505050919050565b611dac61209b565b6000611db661209b565b60008060008060208a01955089519250600190505b8083811515611dd657fe5b0415611deb5760019091019061010002611dcb565b60378311611e625782600101604051805910611e045750595b818152601f19601f830116810160200160405290509450828960ff160160f860020a0285600081518110611e3457fe5b906020010190600160f860020a031916908160001a905350602185019350611e5d848785612056565b611f2d565b828260010101604051805910611e755750595b818152601f19601f830116810160200160405290509450818860ff160160f860020a0285600081518110611ea557fe5b906020010190600160f860020a031916908160001a905350600190505b818111611f1a576101008183036101000a84811515611edd57fe5b04811515611ee757fe5b0660f860020a02858281518110611efa57fe5b906020010190600160f860020a031916908160001a905350600101611ec2565b8160218601019350611f2d848785612056565b509298975050505050505050565b611f4361209b565b600080611f4e61209b565b6000611f5861209b565b600087511515611f8d576000604051805910611f715750595b818152601f19601f83011681016020016040529050965061204b565b600094505b8751851015611fc157878581518110611fa757fe5b906020019060200201515190950194600190940193611f92565b85604051805910611fcf5750595b8181526020601f909201601f191681018201604052600096509450840192505b87518510156120475787858151811061200457fe5b906020019060200201519150506020810161202183828451612056565b87858151811061202d57fe5b906020019060200201515160019095019490920191611fef565b8396505b505050505050919050565b60005b6020821061207c5782518452602084019350602083019250602082039150612059565b6001826020036101000a03905080198351168185511617909352505050565b60206040519081016040526000815290565b606060405190810160409081526000808352602083018190529082015290565b604080519081016040526000808252602082015290565b6060604051908101604052806120f86120cd565b81526020016000815250905600a165627a7a723058201d2eda03250d8d3379040fdc970f7db6fe3791aa02b3f627eca701dcd51bcf9100296060604052341561000f57600080fd5b60008054600160a060020a03191633600160a060020a03161790556020604051908101604052600081526100469060019081610051565b5060006002556100be565b828054828255906000526020600020908101928215610091579160200282015b82811115610091578251829060ff16905591602001919060010190610071565b5061009d9291506100a1565b5090565b6100bb91905b8082111561009d57600081556001016100a7565b90565b6105db806100cd6000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632dcdcd0c811461007157806390b5561d14610099578063b07576ac146100b1578063bda1504b146100c4578063d6362e97146100d7575b600080fd5b341561007c57600080fd5b6100876004356100ea565b60405190815260200160405180910390f35b34156100a457600080fd5b6100af6004356101ba565b005b34156100bc57600080fd5b610087610227565b34156100cf57600080fd5b6100876102ec565b34156100e257600080fd5b6100876102f2565b6000600254610114600161010860028661031590919063ffffffff16565b9063ffffffff61034b16565b11156101325761012b82600263ffffffff61031516565b90506101b5565b60016101498161010885600263ffffffff61031516565b8154811061015357fe5b600091825260209091200154600161017284600263ffffffff61031516565b8154811061017c57fe5b906000526020600020900154101561019f5761012b82600263ffffffff61031516565b61012b600161010884600263ffffffff61031516565b919050565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146101e257600080fd5b600180548082016101f38382610576565b50600091825260209091200181905560025461021690600163ffffffff61034b16565b60028190556102249061035a565b50565b6000805481903373ffffffffffffffffffffffffffffffffffffffff90811691161461025257600080fd5b600180548190811061026057fe5b9060005260206000209001549050600160025481548110151561027f57fe5b90600052602060002090015460018081548110151561029a57fe5b6000918252602090912001556002546001805490919081106102b857fe5b60009182526020822001556002546102d790600163ffffffff61046916565b6002556102e4600161047b565b8091505b5090565b60025481565b600060018081548110151561030357fe5b90600052602060002090015490505b90565b6000808315156103285760009150610344565b5082820282848281151561033857fe5b041461034057fe5b8091505b5092915050565b60008282018381101561034057fe5b60005b600061037083600263ffffffff61055f16565b111561046557600161038983600263ffffffff61055f16565b8154811061039357fe5b9060005260206000209001546001838154811015156103ae57fe5b906000526020600020900154101561044d5760016103d383600263ffffffff61055f16565b815481106103dd57fe5b90600052602060002090015490506001828154811015156103fa57fe5b600091825260209091200154600161041984600263ffffffff61055f16565b8154811061042357fe5b600091825260209091200155600180548291908490811061044057fe5b6000918252602090912001555b61045e82600263ffffffff61055f16565b915061035d565b5050565b60008282111561047557fe5b50900390565b6000805b60025461049660028561031590919063ffffffff16565b1161055a576104a4836100ea565b91506001828154811015156104b557fe5b9060005260206000209001546001848154811015156104d057fe5b90600052602060002090015411156105525760018054849081106104f057fe5b906000526020600020900154905060018281548110151561050d57fe5b90600052602060002090015460018481548110151561052857fe5b600091825260209091200155600180548291908490811061054557fe5b6000918252602090912001555b81925061047f565b505050565b600080828481151561056d57fe5b04949350505050565b81548183558181151161055a5760008381526020902061055a91810190830161031291905b808211156102e8576000815560010161059b5600a165627a7a723058200b3aa9ad9f7f87895de2430edcbcf092bd5df0e72129ace68494aa62d2ff91520029`

// DeployRootChain deploys a new Ethereum contract, binding an instance of RootChain to it.
func DeployRootChain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RootChain, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// RootChain is an auto generated Go binding around an Ethereum contract.
type RootChain struct {
	RootChainCaller     // Read-only binding to the contract
	RootChainTransactor // Write-only binding to the contract
	RootChainFilterer   // Log filterer for contract events
}

// RootChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainSession struct {
	Contract     *RootChain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainCallerSession struct {
	Contract *RootChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RootChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainTransactorSession struct {
	Contract     *RootChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RootChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainRaw struct {
	Contract *RootChain // Generic contract binding to access the raw methods on
}

// RootChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainCallerRaw struct {
	Contract *RootChainCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainTransactorRaw struct {
	Contract *RootChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChain creates a new instance of RootChain, bound to a specific deployed contract.
func NewRootChain(address common.Address, backend bind.ContractBackend) (*RootChain, error) {
	contract, err := bindRootChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// NewRootChainCaller creates a new read-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainCaller(address common.Address, caller bind.ContractCaller) (*RootChainCaller, error) {
	contract, err := bindRootChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainCaller{contract: contract}, nil
}

// NewRootChainTransactor creates a new write-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainTransactor, error) {
	contract, err := bindRootChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainTransactor{contract: contract}, nil
}

// NewRootChainFilterer creates a new log filterer instance of RootChain, bound to a specific deployed contract.
func NewRootChainFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainFilterer, error) {
	contract, err := bindRootChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainFilterer{contract: contract}, nil
}

// bindRootChain binds a generic wrapper to an already deployed contract.
func bindRootChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.RootChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_RootChain *RootChainCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_RootChain *RootChainSession) Authority() (common.Address, error) {
	return _RootChain.Contract.Authority(&_RootChain.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_RootChain *RootChainCallerSession) Authority() (common.Address, error) {
	return _RootChain.Contract.Authority(&_RootChain.CallOpts)
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_RootChain *RootChainCaller) ChildBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "childBlockInterval")
	return *ret0, err
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_RootChain *RootChainSession) ChildBlockInterval() (*big.Int, error) {
	return _RootChain.Contract.ChildBlockInterval(&_RootChain.CallOpts)
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_RootChain *RootChainCallerSession) ChildBlockInterval() (*big.Int, error) {
	return _RootChain.Contract.ChildBlockInterval(&_RootChain.CallOpts)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, created_at uint256)
func (_RootChain *RootChainCaller) ChildChain(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root      [32]byte
	CreatedAt *big.Int
}, error) {
	ret := new(struct {
		Root      [32]byte
		CreatedAt *big.Int
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "childChain", arg0)
	return *ret, err
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, created_at uint256)
func (_RootChain *RootChainSession) ChildChain(arg0 *big.Int) (struct {
	Root      [32]byte
	CreatedAt *big.Int
}, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, created_at uint256)
func (_RootChain *RootChainCallerSession) ChildChain(arg0 *big.Int) (struct {
	Root      [32]byte
	CreatedAt *big.Int
}, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentChildBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentChildBlock")
	return *ret0, err
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentChildBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentChildBlock(&_RootChain.CallOpts)
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentChildBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentChildBlock(&_RootChain.CallOpts)
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentDepositBlock")
	return *ret0, err
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentDepositBlock(&_RootChain.CallOpts)
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentDepositBlock(&_RootChain.CallOpts)
}

// ExitIds is a free data retrieval call binding the contract method 0x16b38840.
//
// Solidity: function exitIds( uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) ExitIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "exitIds", arg0)
	return *ret0, err
}

// ExitIds is a free data retrieval call binding the contract method 0x16b38840.
//
// Solidity: function exitIds( uint256) constant returns(uint256)
func (_RootChain *RootChainSession) ExitIds(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.ExitIds(&_RootChain.CallOpts, arg0)
}

// ExitIds is a free data retrieval call binding the contract method 0x16b38840.
//
// Solidity: function exitIds( uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) ExitIds(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.ExitIds(&_RootChain.CallOpts, arg0)
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(owner address, amount uint256, utxoPos uint256)
func (_RootChain *RootChainCaller) Exits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner   common.Address
	Amount  *big.Int
	UtxoPos *big.Int
}, error) {
	ret := new(struct {
		Owner   common.Address
		Amount  *big.Int
		UtxoPos *big.Int
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "exits", arg0)
	return *ret, err
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(owner address, amount uint256, utxoPos uint256)
func (_RootChain *RootChainSession) Exits(arg0 *big.Int) (struct {
	Owner   common.Address
	Amount  *big.Int
	UtxoPos *big.Int
}, error) {
	return _RootChain.Contract.Exits(&_RootChain.CallOpts, arg0)
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(owner address, amount uint256, utxoPos uint256)
func (_RootChain *RootChainCallerSession) Exits(arg0 *big.Int) (struct {
	Owner   common.Address
	Amount  *big.Int
	UtxoPos *big.Int
}, error) {
	return _RootChain.Contract.Exits(&_RootChain.CallOpts, arg0)
}

// GetChildChain is a free data retrieval call binding the contract method 0x85444de3.
//
// Solidity: function getChildChain(blockNumber uint256) constant returns(bytes32, uint256)
func (_RootChain *RootChainCaller) GetChildChain(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RootChain.contract.Call(opts, out, "getChildChain", blockNumber)
	return *ret0, *ret1, err
}

// GetChildChain is a free data retrieval call binding the contract method 0x85444de3.
//
// Solidity: function getChildChain(blockNumber uint256) constant returns(bytes32, uint256)
func (_RootChain *RootChainSession) GetChildChain(blockNumber *big.Int) ([32]byte, *big.Int, error) {
	return _RootChain.Contract.GetChildChain(&_RootChain.CallOpts, blockNumber)
}

// GetChildChain is a free data retrieval call binding the contract method 0x85444de3.
//
// Solidity: function getChildChain(blockNumber uint256) constant returns(bytes32, uint256)
func (_RootChain *RootChainCallerSession) GetChildChain(blockNumber *big.Int) ([32]byte, *big.Int, error) {
	return _RootChain.Contract.GetChildChain(&_RootChain.CallOpts, blockNumber)
}

// GetDepositBlock is a free data retrieval call binding the contract method 0xbcd59261.
//
// Solidity: function getDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) GetDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getDepositBlock")
	return *ret0, err
}

// GetDepositBlock is a free data retrieval call binding the contract method 0xbcd59261.
//
// Solidity: function getDepositBlock() constant returns(uint256)
func (_RootChain *RootChainSession) GetDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.GetDepositBlock(&_RootChain.CallOpts)
}

// GetDepositBlock is a free data retrieval call binding the contract method 0xbcd59261.
//
// Solidity: function getDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) GetDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.GetDepositBlock(&_RootChain.CallOpts)
}

// GetExit is a free data retrieval call binding the contract method 0xe60f1ff1.
//
// Solidity: function getExit(priority uint256) constant returns(address, uint256, uint256)
func (_RootChain *RootChainCaller) GetExit(opts *bind.CallOpts, priority *big.Int) (common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _RootChain.contract.Call(opts, out, "getExit", priority)
	return *ret0, *ret1, *ret2, err
}

// GetExit is a free data retrieval call binding the contract method 0xe60f1ff1.
//
// Solidity: function getExit(priority uint256) constant returns(address, uint256, uint256)
func (_RootChain *RootChainSession) GetExit(priority *big.Int) (common.Address, *big.Int, *big.Int, error) {
	return _RootChain.Contract.GetExit(&_RootChain.CallOpts, priority)
}

// GetExit is a free data retrieval call binding the contract method 0xe60f1ff1.
//
// Solidity: function getExit(priority uint256) constant returns(address, uint256, uint256)
func (_RootChain *RootChainCallerSession) GetExit(priority *big.Int) (common.Address, *big.Int, *big.Int, error) {
	return _RootChain.Contract.GetExit(&_RootChain.CallOpts, priority)
}

// NextWeekOldChildBlock is a free data retrieval call binding the contract method 0x107e58f1.
//
// Solidity: function nextWeekOldChildBlock(value uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) NextWeekOldChildBlock(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "nextWeekOldChildBlock", value)
	return *ret0, err
}

// NextWeekOldChildBlock is a free data retrieval call binding the contract method 0x107e58f1.
//
// Solidity: function nextWeekOldChildBlock(value uint256) constant returns(uint256)
func (_RootChain *RootChainSession) NextWeekOldChildBlock(value *big.Int) (*big.Int, error) {
	return _RootChain.Contract.NextWeekOldChildBlock(&_RootChain.CallOpts, value)
}

// NextWeekOldChildBlock is a free data retrieval call binding the contract method 0x107e58f1.
//
// Solidity: function nextWeekOldChildBlock(value uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) NextWeekOldChildBlock(value *big.Int) (*big.Int, error) {
	return _RootChain.Contract.NextWeekOldChildBlock(&_RootChain.CallOpts, value)
}

// RecentBlock is a free data retrieval call binding the contract method 0x6f84b695.
//
// Solidity: function recentBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) RecentBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "recentBlock")
	return *ret0, err
}

// RecentBlock is a free data retrieval call binding the contract method 0x6f84b695.
//
// Solidity: function recentBlock() constant returns(uint256)
func (_RootChain *RootChainSession) RecentBlock() (*big.Int, error) {
	return _RootChain.Contract.RecentBlock(&_RootChain.CallOpts)
}

// RecentBlock is a free data retrieval call binding the contract method 0x6f84b695.
//
// Solidity: function recentBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) RecentBlock() (*big.Int, error) {
	return _RootChain.Contract.RecentBlock(&_RootChain.CallOpts)
}

// WeekOldBlock is a free data retrieval call binding the contract method 0x4237b5f3.
//
// Solidity: function weekOldBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) WeekOldBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "weekOldBlock")
	return *ret0, err
}

// WeekOldBlock is a free data retrieval call binding the contract method 0x4237b5f3.
//
// Solidity: function weekOldBlock() constant returns(uint256)
func (_RootChain *RootChainSession) WeekOldBlock() (*big.Int, error) {
	return _RootChain.Contract.WeekOldBlock(&_RootChain.CallOpts)
}

// WeekOldBlock is a free data retrieval call binding the contract method 0x4237b5f3.
//
// Solidity: function weekOldBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) WeekOldBlock() (*big.Int, error) {
	return _RootChain.Contract.WeekOldBlock(&_RootChain.CallOpts)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x32773ba3.
//
// Solidity: function challengeExit(cUtxoPos uint256, eUtxoPos uint256, txBytes bytes, proof bytes, sigs bytes, confirmationSig bytes) returns()
func (_RootChain *RootChainTransactor) ChallengeExit(opts *bind.TransactOpts, cUtxoPos *big.Int, eUtxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte, confirmationSig []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeExit", cUtxoPos, eUtxoPos, txBytes, proof, sigs, confirmationSig)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x32773ba3.
//
// Solidity: function challengeExit(cUtxoPos uint256, eUtxoPos uint256, txBytes bytes, proof bytes, sigs bytes, confirmationSig bytes) returns()
func (_RootChain *RootChainSession) ChallengeExit(cUtxoPos *big.Int, eUtxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte, confirmationSig []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, cUtxoPos, eUtxoPos, txBytes, proof, sigs, confirmationSig)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x32773ba3.
//
// Solidity: function challengeExit(cUtxoPos uint256, eUtxoPos uint256, txBytes bytes, proof bytes, sigs bytes, confirmationSig bytes) returns()
func (_RootChain *RootChainTransactorSession) ChallengeExit(cUtxoPos *big.Int, eUtxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte, confirmationSig []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, cUtxoPos, eUtxoPos, txBytes, proof, sigs, confirmationSig)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_RootChain *RootChainTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_RootChain *RootChainSession) Deposit() (*types.Transaction, error) {
	return _RootChain.Contract.Deposit(&_RootChain.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_RootChain *RootChainTransactorSession) Deposit() (*types.Transaction, error) {
	return _RootChain.Contract.Deposit(&_RootChain.TransactOpts)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0xc6ab44cd.
//
// Solidity: function finalizeExits() returns(uint256)
func (_RootChain *RootChainTransactor) FinalizeExits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeExits")
}

// FinalizeExits is a paid mutator transaction binding the contract method 0xc6ab44cd.
//
// Solidity: function finalizeExits() returns(uint256)
func (_RootChain *RootChainSession) FinalizeExits() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeExits(&_RootChain.TransactOpts)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0xc6ab44cd.
//
// Solidity: function finalizeExits() returns(uint256)
func (_RootChain *RootChainTransactorSession) FinalizeExits() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeExits(&_RootChain.TransactOpts)
}

// StartExit is a paid mutator transaction binding the contract method 0x1c91a6b9.
//
// Solidity: function startExit(utxoPos uint256, txBytes bytes, proof bytes, sigs bytes) returns()
func (_RootChain *RootChainTransactor) StartExit(opts *bind.TransactOpts, utxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startExit", utxoPos, txBytes, proof, sigs)
}

// StartExit is a paid mutator transaction binding the contract method 0x1c91a6b9.
//
// Solidity: function startExit(utxoPos uint256, txBytes bytes, proof bytes, sigs bytes) returns()
func (_RootChain *RootChainSession) StartExit(utxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, utxoPos, txBytes, proof, sigs)
}

// StartExit is a paid mutator transaction binding the contract method 0x1c91a6b9.
//
// Solidity: function startExit(utxoPos uint256, txBytes bytes, proof bytes, sigs bytes) returns()
func (_RootChain *RootChainTransactorSession) StartExit(utxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, utxoPos, txBytes, proof, sigs)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_RootChain *RootChainTransactor) SubmitBlock(opts *bind.TransactOpts, root [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitBlock", root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_RootChain *RootChainSession) SubmitBlock(root [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitBlock(&_RootChain.TransactOpts, root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_RootChain *RootChainTransactorSession) SubmitBlock(root [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitBlock(&_RootChain.TransactOpts, root)
}

// RootChainDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the RootChain contract.
type RootChainDepositIterator struct {
	Event *RootChainDeposit // Event containing the contract specifics and raw log

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
func (it *RootChainDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainDeposit)
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
		it.Event = new(RootChainDeposit)
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
func (it *RootChainDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainDeposit represents a Deposit event raised by the RootChain contract.
type RootChainDeposit struct {
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(depositor address, amount uint256)
func (_RootChain *RootChainFilterer) FilterDeposit(opts *bind.FilterOpts) (*RootChainDepositIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &RootChainDepositIterator{contract: _RootChain.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(depositor address, amount uint256)
func (_RootChain *RootChainFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *RootChainDeposit) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainDeposit)
				if err := _RootChain.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// RootChainExitIterator is returned from FilterExit and is used to iterate over the raw logs and unpacked data for Exit events raised by the RootChain contract.
type RootChainExitIterator struct {
	Event *RootChainExit // Event containing the contract specifics and raw log

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
func (it *RootChainExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainExit)
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
		it.Event = new(RootChainExit)
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
func (it *RootChainExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainExit represents a Exit event raised by the RootChain contract.
type RootChainExit struct {
	Exitor  common.Address
	UtxoPos *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExit is a free log retrieval operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(exitor address, utxoPos uint256)
func (_RootChain *RootChainFilterer) FilterExit(opts *bind.FilterOpts) (*RootChainExitIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Exit")
	if err != nil {
		return nil, err
	}
	return &RootChainExitIterator{contract: _RootChain.contract, event: "Exit", logs: logs, sub: sub}, nil
}

// WatchExit is a free log subscription operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(exitor address, utxoPos uint256)
func (_RootChain *RootChainFilterer) WatchExit(opts *bind.WatchOpts, sink chan<- *RootChainExit) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Exit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainExit)
				if err := _RootChain.contract.UnpackLog(event, "Exit", log); err != nil {
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
const SafeMathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a723058200f0319e262a6b3b5bac16e7d380b3e2215adee7b0a26c4589e18aeffe74f58ec0029`

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

// ValidateABI is the input ABI used to generate the binding from.
const ValidateABI = "[]"

// ValidateBin is the compiled bytecode used for deploying new contracts.
const ValidateBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a723058209d710440c3c7c67c1b879b4ca9661fcab4578d9aab953df091956a98b3141ca00029`

// DeployValidate deploys a new Ethereum contract, binding an instance of Validate to it.
func DeployValidate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validate, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validate{ValidateCaller: ValidateCaller{contract: contract}, ValidateTransactor: ValidateTransactor{contract: contract}, ValidateFilterer: ValidateFilterer{contract: contract}}, nil
}

// Validate is an auto generated Go binding around an Ethereum contract.
type Validate struct {
	ValidateCaller     // Read-only binding to the contract
	ValidateTransactor // Write-only binding to the contract
	ValidateFilterer   // Log filterer for contract events
}

// ValidateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidateSession struct {
	Contract     *Validate         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidateCallerSession struct {
	Contract *ValidateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ValidateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidateTransactorSession struct {
	Contract     *ValidateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ValidateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidateRaw struct {
	Contract *Validate // Generic contract binding to access the raw methods on
}

// ValidateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidateCallerRaw struct {
	Contract *ValidateCaller // Generic read-only contract binding to access the raw methods on
}

// ValidateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidateTransactorRaw struct {
	Contract *ValidateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidate creates a new instance of Validate, bound to a specific deployed contract.
func NewValidate(address common.Address, backend bind.ContractBackend) (*Validate, error) {
	contract, err := bindValidate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validate{ValidateCaller: ValidateCaller{contract: contract}, ValidateTransactor: ValidateTransactor{contract: contract}, ValidateFilterer: ValidateFilterer{contract: contract}}, nil
}

// NewValidateCaller creates a new read-only instance of Validate, bound to a specific deployed contract.
func NewValidateCaller(address common.Address, caller bind.ContractCaller) (*ValidateCaller, error) {
	contract, err := bindValidate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidateCaller{contract: contract}, nil
}

// NewValidateTransactor creates a new write-only instance of Validate, bound to a specific deployed contract.
func NewValidateTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidateTransactor, error) {
	contract, err := bindValidate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidateTransactor{contract: contract}, nil
}

// NewValidateFilterer creates a new log filterer instance of Validate, bound to a specific deployed contract.
func NewValidateFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidateFilterer, error) {
	contract, err := bindValidate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidateFilterer{contract: contract}, nil
}

// bindValidate binds a generic wrapper to an already deployed contract.
func bindValidate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validate *ValidateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validate.Contract.ValidateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validate *ValidateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validate.Contract.ValidateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validate *ValidateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validate.Contract.ValidateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validate *ValidateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validate *ValidateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validate *ValidateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validate.Contract.contract.Transact(opts, method, params...)
}
