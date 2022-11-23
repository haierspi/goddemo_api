// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package conflux_chain

import (
	"math/big"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient/bulk"
	types "github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = ethBind.Bind
	_ = common.Big1
	_ = ethtypes.BloomLookup
	_ = event.NewSubscription
)

// MyTokenABI is the input ABI used to generate the binding from.
const MyTokenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter_\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelist\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_initialOwners\",\"type\":\"address[]\"}],\"name\":\"batchAddItemByAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"burnBatchByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burnByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter_\",\"type\":\"address\"}],\"name\":\"removeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelist\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"featureCode\",\"type\":\"uint256\"}],\"name\":\"setTokenFeatureCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newuri\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenFeatureCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"tokensOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MyToken is an auto generated Go binding around an Conflux contract.
type MyToken struct {
	MyTokenCaller     // Read-only binding to the contract
	MyTokenTransactor // Write-only binding to the contract
	MyTokenFilterer   // Log filterer for contract events
}

// MyTokenCaller is an auto generated read-only Go binding around an Conflux contract.
type MyTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type MyTokenBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenTransactor is an auto generated write-only Go binding around an Conflux contract.
type MyTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type MyTokenBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type MyTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type MyTokenSession struct {
	Contract     *MyToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyTokenCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type MyTokenCallerSession struct {
	Contract *MyTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MyTokenTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type MyTokenTransactorSession struct {
	Contract     *MyTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MyTokenRaw is an auto generated low-level Go binding around an Conflux contract.
type MyTokenRaw struct {
	Contract *MyToken // Generic contract binding to access the raw methods on
}

// MyTokenCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type MyTokenCallerRaw struct {
	Contract *MyTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MyTokenTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type MyTokenTransactorRaw struct {
	Contract *MyTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyToken creates a new instance of MyToken, bound to a specific deployed contract.
func NewMyToken(address types.Address, backend bind.ContractBackend) (*MyToken, error) {
	contract, err := bindMyToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyToken{MyTokenCaller: MyTokenCaller{contract: contract}, MyTokenTransactor: MyTokenTransactor{contract: contract}, MyTokenFilterer: MyTokenFilterer{contract: contract}}, nil
}

// NewMyTokenCaller creates a new read-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenCaller(address types.Address, caller bind.ContractCaller) (*MyTokenCaller, error) {
	contract, err := bindMyToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenCaller{contract: contract}, nil
}

// NewMyTokenTransactor creates a new write-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenTransactor(address types.Address, transactor bind.ContractTransactor) (*MyTokenTransactor, error) {
	contract, err := bindMyToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenTransactor{contract: contract}, nil
}

// NewMyTokenFilterer creates a new log filterer instance of MyToken, bound to a specific deployed contract.
func NewMyTokenFilterer(address types.Address, filterer bind.ContractFilterer) (*MyTokenFilterer, error) {
	contract, err := bindMyToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyTokenFilterer{contract: contract}, nil
}

// NewMyTokenCaller creates a new read-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenBulkCaller(address types.Address, caller bind.ContractCaller) (*MyTokenBulkCaller, error) {
	contract, err := bindMyToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenBulkCaller{contract: contract}, nil
}

// NewMyTokenBulkTransactor creates a new write-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*MyTokenBulkTransactor, error) {
	contract, err := bindMyToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenBulkTransactor{contract: contract}, nil
}

// bindMyToken binds a generic wrapper to an already deployed contract.
func bindMyToken(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyToken *MyTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyToken.Contract.MyTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyToken *MyTokenRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.MyTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyToken *MyTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.MyTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyToken *MyTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyToken *MyTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyToken *MyTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.contract.Transact(opts, method, params...)
}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_MyToken *MyTokenCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "BURNER_ROLE")

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_MyToken *MyTokenBulkCaller) BURNERROLE(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "BURNER_ROLE")

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "BURNER_ROLE")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_MyToken *MyTokenSession) BURNERROLE() ([32]byte, error) {
	return _MyToken.Contract.BURNERROLE(&_MyToken.CallOpts)
}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_MyToken *MyTokenCallerSession) BURNERROLE() ([32]byte, error) {
	return _MyToken.Contract.BURNERROLE(&_MyToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MyToken *MyTokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MyToken *MyTokenBulkCaller) DEFAULTADMINROLE(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "DEFAULT_ADMIN_ROLE")

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "DEFAULT_ADMIN_ROLE")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MyToken *MyTokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MyToken.Contract.DEFAULTADMINROLE(&_MyToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MyToken *MyTokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MyToken.Contract.DEFAULTADMINROLE(&_MyToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_MyToken *MyTokenCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "MINTER_ROLE")

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_MyToken *MyTokenBulkCaller) MINTERROLE(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "MINTER_ROLE")

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "MINTER_ROLE")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_MyToken *MyTokenSession) MINTERROLE() ([32]byte, error) {
	return _MyToken.Contract.MINTERROLE(&_MyToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_MyToken *MyTokenCallerSession) MINTERROLE() ([32]byte, error) {
	return _MyToken.Contract.MINTERROLE(&_MyToken.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MyToken *MyTokenCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "PAUSER_ROLE")

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MyToken *MyTokenBulkCaller) PAUSERROLE(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "PAUSER_ROLE")

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "PAUSER_ROLE")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MyToken *MyTokenSession) PAUSERROLE() ([32]byte, error) {
	return _MyToken.Contract.PAUSERROLE(&_MyToken.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MyToken *MyTokenCallerSession) PAUSERROLE() ([32]byte, error) {
	return _MyToken.Contract.PAUSERROLE(&_MyToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyToken *MyTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "balanceOf", owner)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyToken *MyTokenBulkCaller) BalanceOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "balanceOf", owner)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "balanceOf")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyToken *MyTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyToken.Contract.BalanceOf(&_MyToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyToken *MyTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyToken.Contract.BalanceOf(&_MyToken.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyToken *MyTokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "getApproved", tokenId)

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyToken *MyTokenBulkCaller) GetApproved(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "getApproved", tokenId)

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "getApproved")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyToken *MyTokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MyToken.Contract.GetApproved(&_MyToken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyToken *MyTokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MyToken.Contract.GetApproved(&_MyToken.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MyToken *MyTokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "getRoleAdmin", role)

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MyToken *MyTokenBulkCaller) GetRoleAdmin(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, role [32]byte) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "getRoleAdmin", role)

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "getRoleAdmin")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MyToken *MyTokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MyToken.Contract.GetRoleAdmin(&_MyToken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MyToken *MyTokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MyToken.Contract.GetRoleAdmin(&_MyToken.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MyToken *MyTokenCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "getRoleMember", role, index)

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MyToken *MyTokenBulkCaller) GetRoleMember(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, role [32]byte, index *big.Int) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "getRoleMember", role, index)

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "getRoleMember")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MyToken *MyTokenSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MyToken.Contract.GetRoleMember(&_MyToken.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MyToken *MyTokenCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MyToken.Contract.GetRoleMember(&_MyToken.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MyToken *MyTokenCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "getRoleMemberCount", role)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MyToken *MyTokenBulkCaller) GetRoleMemberCount(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, role [32]byte) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "getRoleMemberCount", role)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "getRoleMemberCount")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MyToken *MyTokenSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MyToken.Contract.GetRoleMemberCount(&_MyToken.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MyToken *MyTokenCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MyToken.Contract.GetRoleMemberCount(&_MyToken.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MyToken *MyTokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "hasRole", role, account)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MyToken *MyTokenBulkCaller) HasRole(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, role [32]byte, account common.Address) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "hasRole", role, account)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "hasRole")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MyToken *MyTokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MyToken.Contract.HasRole(&_MyToken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MyToken *MyTokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MyToken.Contract.HasRole(&_MyToken.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MyToken *MyTokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MyToken *MyTokenBulkCaller) IsApprovedForAll(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address, operator common.Address) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "isApprovedForAll", owner, operator)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "isApprovedForAll")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MyToken *MyTokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MyToken.Contract.IsApprovedForAll(&_MyToken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MyToken *MyTokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MyToken.Contract.IsApprovedForAll(&_MyToken.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyToken *MyTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "name")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyToken *MyTokenBulkCaller) Name(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "name")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "name")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyToken *MyTokenSession) Name() (string, error) {
	return _MyToken.Contract.Name(&_MyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyToken *MyTokenCallerSession) Name() (string, error) {
	return _MyToken.Contract.Name(&_MyToken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyToken *MyTokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "ownerOf", tokenId)

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyToken *MyTokenBulkCaller) OwnerOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "ownerOf", tokenId)

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "ownerOf")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyToken *MyTokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MyToken.Contract.OwnerOf(&_MyToken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyToken *MyTokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MyToken.Contract.OwnerOf(&_MyToken.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyToken *MyTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyToken *MyTokenBulkCaller) SupportsInterface(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, interfaceId [4]byte) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "supportsInterface", interfaceId)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "supportsInterface")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyToken *MyTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MyToken.Contract.SupportsInterface(&_MyToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyToken *MyTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MyToken.Contract.SupportsInterface(&_MyToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyToken *MyTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "symbol")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyToken *MyTokenBulkCaller) Symbol(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "symbol")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "symbol")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyToken *MyTokenSession) Symbol() (string, error) {
	return _MyToken.Contract.Symbol(&_MyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyToken *MyTokenCallerSession) Symbol() (string, error) {
	return _MyToken.Contract.Symbol(&_MyToken.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_MyToken *MyTokenCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "tokenByIndex", index)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_MyToken *MyTokenBulkCaller) TokenByIndex(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, index *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "tokenByIndex", index)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "tokenByIndex")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_MyToken *MyTokenSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _MyToken.Contract.TokenByIndex(&_MyToken.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_MyToken *MyTokenCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _MyToken.Contract.TokenByIndex(&_MyToken.CallOpts, index)
}

// TokenFeatureCode is a free data retrieval call binding the contract method 0xf65a60cb.
//
// Solidity: function tokenFeatureCode(uint256 ) view returns(uint256)
func (_MyToken *MyTokenCaller) TokenFeatureCode(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "tokenFeatureCode", arg0)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TokenFeatureCode is a free data retrieval call binding the contract method 0xf65a60cb.
//
// Solidity: function tokenFeatureCode(uint256 ) view returns(uint256)
func (_MyToken *MyTokenBulkCaller) TokenFeatureCode(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, arg0 *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "tokenFeatureCode", arg0)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "tokenFeatureCode")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokenFeatureCode is a free data retrieval call binding the contract method 0xf65a60cb.
//
// Solidity: function tokenFeatureCode(uint256 ) view returns(uint256)
func (_MyToken *MyTokenSession) TokenFeatureCode(arg0 *big.Int) (*big.Int, error) {
	return _MyToken.Contract.TokenFeatureCode(&_MyToken.CallOpts, arg0)
}

// TokenFeatureCode is a free data retrieval call binding the contract method 0xf65a60cb.
//
// Solidity: function tokenFeatureCode(uint256 ) view returns(uint256)
func (_MyToken *MyTokenCallerSession) TokenFeatureCode(arg0 *big.Int) (*big.Int, error) {
	return _MyToken.Contract.TokenFeatureCode(&_MyToken.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_MyToken *MyTokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_MyToken *MyTokenBulkCaller) TokenOfOwnerByIndex(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address, index *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "tokenOfOwnerByIndex", owner, index)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "tokenOfOwnerByIndex")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_MyToken *MyTokenSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _MyToken.Contract.TokenOfOwnerByIndex(&_MyToken.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_MyToken *MyTokenCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _MyToken.Contract.TokenOfOwnerByIndex(&_MyToken.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyToken *MyTokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "tokenURI", tokenId)

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyToken *MyTokenBulkCaller) TokenURI(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "tokenURI", tokenId)

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "tokenURI")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyToken *MyTokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MyToken.Contract.TokenURI(&_MyToken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyToken *MyTokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MyToken.Contract.TokenURI(&_MyToken.CallOpts, tokenId)
}

// Tokens is a free data retrieval call binding the contract method 0x8b4864d6.
//
// Solidity: function tokens(uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_MyToken *MyTokenCaller) Tokens(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "tokens", offset, limit)

	outstruct := new(struct {
		Total    *big.Int
		TokenIds []*big.Int
	})
	if __err != nil {
		return *outstruct, __err
	}

	outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, __err

}

// Tokens is a free data retrieval call binding the contract method 0x8b4864d6.
//
// Solidity: function tokens(uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_MyToken *MyTokenBulkCaller) Tokens(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, offset *big.Int, limit *big.Int) (*struct {
	Total    *big.Int
	TokenIds []*big.Int
}, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "tokens", offset, limit)

	outstruct := new(struct {
		Total    *big.Int
		TokenIds []*big.Int
	})

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "tokens")
		if err != nil {
			return err
		}

		outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		outstruct.TokenIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return outstruct, __err

}

// Tokens is a free data retrieval call binding the contract method 0x8b4864d6.
//
// Solidity: function tokens(uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_MyToken *MyTokenSession) Tokens(offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	return _MyToken.Contract.Tokens(&_MyToken.CallOpts, offset, limit)
}

// Tokens is a free data retrieval call binding the contract method 0x8b4864d6.
//
// Solidity: function tokens(uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_MyToken *MyTokenCallerSession) Tokens(offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	return _MyToken.Contract.Tokens(&_MyToken.CallOpts, offset, limit)
}

// TokensOf is a free data retrieval call binding the contract method 0x23185dc9.
//
// Solidity: function tokensOf(address owner, uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_MyToken *MyTokenCaller) TokensOf(opts *bind.CallOpts, owner common.Address, offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "tokensOf", owner, offset, limit)

	outstruct := new(struct {
		Total    *big.Int
		TokenIds []*big.Int
	})
	if __err != nil {
		return *outstruct, __err
	}

	outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, __err

}

// TokensOf is a free data retrieval call binding the contract method 0x23185dc9.
//
// Solidity: function tokensOf(address owner, uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_MyToken *MyTokenBulkCaller) TokensOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address, offset *big.Int, limit *big.Int) (*struct {
	Total    *big.Int
	TokenIds []*big.Int
}, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "tokensOf", owner, offset, limit)

	outstruct := new(struct {
		Total    *big.Int
		TokenIds []*big.Int
	})

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "tokensOf")
		if err != nil {
			return err
		}

		outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		outstruct.TokenIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return outstruct, __err

}

// TokensOf is a free data retrieval call binding the contract method 0x23185dc9.
//
// Solidity: function tokensOf(address owner, uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_MyToken *MyTokenSession) TokensOf(owner common.Address, offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	return _MyToken.Contract.TokensOf(&_MyToken.CallOpts, owner, offset, limit)
}

// TokensOf is a free data retrieval call binding the contract method 0x23185dc9.
//
// Solidity: function tokensOf(address owner, uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_MyToken *MyTokenCallerSession) TokensOf(owner common.Address, offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	return _MyToken.Contract.TokensOf(&_MyToken.CallOpts, owner, offset, limit)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyToken *MyTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _MyToken.contract.Call(opts, &out, "totalSupply")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyToken *MyTokenBulkCaller) TotalSupply(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyToken.contract.GenRequest(opts, "totalSupply")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "totalSupply")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyToken *MyTokenSession) TotalSupply() (*big.Int, error) {
	return _MyToken.Contract.TotalSupply(&_MyToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyToken *MyTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MyToken.Contract.TotalSupply(&_MyToken.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter_) returns()
func (_MyToken *MyTokenTransactor) AddMinter(opts *bind.TransactOpts, minter_ common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "addMinter", minter_)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter_) returns()
func (_MyToken *MyTokenBulkTransactor) AddMinter(opts *bind.TransactOpts, minter_ common.Address) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "addMinter", minter_)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter_) returns()
func (_MyToken *MyTokenSession) AddMinter(minter_ common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.AddMinter(&_MyToken.TransactOpts, minter_)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter_) returns()
func (_MyToken *MyTokenTransactorSession) AddMinter(minter_ common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.AddMinter(&_MyToken.TransactOpts, minter_)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address account) payable returns()
func (_MyToken *MyTokenTransactor) AddWhitelist(opts *bind.TransactOpts, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "addWhitelist", account)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address account) payable returns()
func (_MyToken *MyTokenBulkTransactor) AddWhitelist(opts *bind.TransactOpts, account common.Address) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "addWhitelist", account)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address account) payable returns()
func (_MyToken *MyTokenSession) AddWhitelist(account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.AddWhitelist(&_MyToken.TransactOpts, account)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address account) payable returns()
func (_MyToken *MyTokenTransactorSession) AddWhitelist(account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.AddWhitelist(&_MyToken.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyToken *MyTokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyToken *MyTokenBulkTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyToken *MyTokenSession) Approve(to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.Approve(&_MyToken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyToken *MyTokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.Approve(&_MyToken.TransactOpts, to, tokenId)
}

// BatchAddItemByAddress is a paid mutator transaction binding the contract method 0x54777f41.
//
// Solidity: function batchAddItemByAddress(address[] _initialOwners) returns()
func (_MyToken *MyTokenTransactor) BatchAddItemByAddress(opts *bind.TransactOpts, _initialOwners []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "batchAddItemByAddress", _initialOwners)
}

// BatchAddItemByAddress is a paid mutator transaction binding the contract method 0x54777f41.
//
// Solidity: function batchAddItemByAddress(address[] _initialOwners) returns()
func (_MyToken *MyTokenBulkTransactor) BatchAddItemByAddress(opts *bind.TransactOpts, _initialOwners []common.Address) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "batchAddItemByAddress", _initialOwners)
}

// BatchAddItemByAddress is a paid mutator transaction binding the contract method 0x54777f41.
//
// Solidity: function batchAddItemByAddress(address[] _initialOwners) returns()
func (_MyToken *MyTokenSession) BatchAddItemByAddress(_initialOwners []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.BatchAddItemByAddress(&_MyToken.TransactOpts, _initialOwners)
}

// BatchAddItemByAddress is a paid mutator transaction binding the contract method 0x54777f41.
//
// Solidity: function batchAddItemByAddress(address[] _initialOwners) returns()
func (_MyToken *MyTokenTransactorSession) BatchAddItemByAddress(_initialOwners []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.BatchAddItemByAddress(&_MyToken.TransactOpts, _initialOwners)
}

// BurnBatchByAdmin is a paid mutator transaction binding the contract method 0xd57caa8e.
//
// Solidity: function burnBatchByAdmin(uint256[] tokenIds) returns()
func (_MyToken *MyTokenTransactor) BurnBatchByAdmin(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "burnBatchByAdmin", tokenIds)
}

// BurnBatchByAdmin is a paid mutator transaction binding the contract method 0xd57caa8e.
//
// Solidity: function burnBatchByAdmin(uint256[] tokenIds) returns()
func (_MyToken *MyTokenBulkTransactor) BurnBatchByAdmin(opts *bind.TransactOpts, tokenIds []*big.Int) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "burnBatchByAdmin", tokenIds)
}

// BurnBatchByAdmin is a paid mutator transaction binding the contract method 0xd57caa8e.
//
// Solidity: function burnBatchByAdmin(uint256[] tokenIds) returns()
func (_MyToken *MyTokenSession) BurnBatchByAdmin(tokenIds []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.BurnBatchByAdmin(&_MyToken.TransactOpts, tokenIds)
}

// BurnBatchByAdmin is a paid mutator transaction binding the contract method 0xd57caa8e.
//
// Solidity: function burnBatchByAdmin(uint256[] tokenIds) returns()
func (_MyToken *MyTokenTransactorSession) BurnBatchByAdmin(tokenIds []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.BurnBatchByAdmin(&_MyToken.TransactOpts, tokenIds)
}

// BurnByAdmin is a paid mutator transaction binding the contract method 0x21ea5308.
//
// Solidity: function burnByAdmin(uint256 tokenId) returns()
func (_MyToken *MyTokenTransactor) BurnByAdmin(opts *bind.TransactOpts, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "burnByAdmin", tokenId)
}

// BurnByAdmin is a paid mutator transaction binding the contract method 0x21ea5308.
//
// Solidity: function burnByAdmin(uint256 tokenId) returns()
func (_MyToken *MyTokenBulkTransactor) BurnByAdmin(opts *bind.TransactOpts, tokenId *big.Int) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "burnByAdmin", tokenId)
}

// BurnByAdmin is a paid mutator transaction binding the contract method 0x21ea5308.
//
// Solidity: function burnByAdmin(uint256 tokenId) returns()
func (_MyToken *MyTokenSession) BurnByAdmin(tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.BurnByAdmin(&_MyToken.TransactOpts, tokenId)
}

// BurnByAdmin is a paid mutator transaction binding the contract method 0x21ea5308.
//
// Solidity: function burnByAdmin(uint256 tokenId) returns()
func (_MyToken *MyTokenTransactorSession) BurnByAdmin(tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.BurnByAdmin(&_MyToken.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MyToken *MyTokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MyToken *MyTokenBulkTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MyToken *MyTokenSession) GrantRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.GrantRole(&_MyToken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MyToken *MyTokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.GrantRole(&_MyToken.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_MyToken *MyTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_MyToken *MyTokenBulkTransactor) Mint(opts *bind.TransactOpts, to common.Address) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_MyToken *MyTokenSession) Mint(to common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.Mint(&_MyToken.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_MyToken *MyTokenTransactorSession) Mint(to common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.Mint(&_MyToken.TransactOpts, to)
}

// MintBatch is a paid mutator transaction binding the contract method 0x248b71fc.
//
// Solidity: function mintBatch(address to, uint256 amount) returns()
func (_MyToken *MyTokenTransactor) MintBatch(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "mintBatch", to, amount)
}

// MintBatch is a paid mutator transaction binding the contract method 0x248b71fc.
//
// Solidity: function mintBatch(address to, uint256 amount) returns()
func (_MyToken *MyTokenBulkTransactor) MintBatch(opts *bind.TransactOpts, to common.Address, amount *big.Int) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "mintBatch", to, amount)
}

// MintBatch is a paid mutator transaction binding the contract method 0x248b71fc.
//
// Solidity: function mintBatch(address to, uint256 amount) returns()
func (_MyToken *MyTokenSession) MintBatch(to common.Address, amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.MintBatch(&_MyToken.TransactOpts, to, amount)
}

// MintBatch is a paid mutator transaction binding the contract method 0x248b71fc.
//
// Solidity: function mintBatch(address to, uint256 amount) returns()
func (_MyToken *MyTokenTransactorSession) MintBatch(to common.Address, amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.MintBatch(&_MyToken.TransactOpts, to, amount)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter_) returns()
func (_MyToken *MyTokenTransactor) RemoveMinter(opts *bind.TransactOpts, minter_ common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "removeMinter", minter_)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter_) returns()
func (_MyToken *MyTokenBulkTransactor) RemoveMinter(opts *bind.TransactOpts, minter_ common.Address) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "removeMinter", minter_)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter_) returns()
func (_MyToken *MyTokenSession) RemoveMinter(minter_ common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.RemoveMinter(&_MyToken.TransactOpts, minter_)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter_) returns()
func (_MyToken *MyTokenTransactorSession) RemoveMinter(minter_ common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.RemoveMinter(&_MyToken.TransactOpts, minter_)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x78c8cda7.
//
// Solidity: function removeWhitelist(address account) payable returns()
func (_MyToken *MyTokenTransactor) RemoveWhitelist(opts *bind.TransactOpts, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "removeWhitelist", account)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x78c8cda7.
//
// Solidity: function removeWhitelist(address account) payable returns()
func (_MyToken *MyTokenBulkTransactor) RemoveWhitelist(opts *bind.TransactOpts, account common.Address) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "removeWhitelist", account)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x78c8cda7.
//
// Solidity: function removeWhitelist(address account) payable returns()
func (_MyToken *MyTokenSession) RemoveWhitelist(account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.RemoveWhitelist(&_MyToken.TransactOpts, account)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x78c8cda7.
//
// Solidity: function removeWhitelist(address account) payable returns()
func (_MyToken *MyTokenTransactorSession) RemoveWhitelist(account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.RemoveWhitelist(&_MyToken.TransactOpts, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MyToken *MyTokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MyToken *MyTokenBulkTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MyToken *MyTokenSession) RenounceRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.RenounceRole(&_MyToken.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MyToken *MyTokenTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.RenounceRole(&_MyToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MyToken *MyTokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MyToken *MyTokenBulkTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MyToken *MyTokenSession) RevokeRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.RevokeRole(&_MyToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MyToken *MyTokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.RevokeRole(&_MyToken.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MyToken *MyTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MyToken *MyTokenBulkTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MyToken *MyTokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.SafeTransferFrom(&_MyToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MyToken *MyTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.SafeTransferFrom(&_MyToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MyToken *MyTokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MyToken *MyTokenBulkTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MyToken *MyTokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.SafeTransferFrom0(&_MyToken.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MyToken *MyTokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.SafeTransferFrom0(&_MyToken.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyToken *MyTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyToken *MyTokenBulkTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyToken *MyTokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.SetApprovalForAll(&_MyToken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyToken *MyTokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.SetApprovalForAll(&_MyToken.TransactOpts, operator, approved)
}

// SetTokenFeatureCode is a paid mutator transaction binding the contract method 0x1d61c7e1.
//
// Solidity: function setTokenFeatureCode(uint256 tokenId, uint256 featureCode) returns()
func (_MyToken *MyTokenTransactor) SetTokenFeatureCode(opts *bind.TransactOpts, tokenId *big.Int, featureCode *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "setTokenFeatureCode", tokenId, featureCode)
}

// SetTokenFeatureCode is a paid mutator transaction binding the contract method 0x1d61c7e1.
//
// Solidity: function setTokenFeatureCode(uint256 tokenId, uint256 featureCode) returns()
func (_MyToken *MyTokenBulkTransactor) SetTokenFeatureCode(opts *bind.TransactOpts, tokenId *big.Int, featureCode *big.Int) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "setTokenFeatureCode", tokenId, featureCode)
}

// SetTokenFeatureCode is a paid mutator transaction binding the contract method 0x1d61c7e1.
//
// Solidity: function setTokenFeatureCode(uint256 tokenId, uint256 featureCode) returns()
func (_MyToken *MyTokenSession) SetTokenFeatureCode(tokenId *big.Int, featureCode *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.SetTokenFeatureCode(&_MyToken.TransactOpts, tokenId, featureCode)
}

// SetTokenFeatureCode is a paid mutator transaction binding the contract method 0x1d61c7e1.
//
// Solidity: function setTokenFeatureCode(uint256 tokenId, uint256 featureCode) returns()
func (_MyToken *MyTokenTransactorSession) SetTokenFeatureCode(tokenId *big.Int, featureCode *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.SetTokenFeatureCode(&_MyToken.TransactOpts, tokenId, featureCode)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_MyToken *MyTokenTransactor) SetURI(opts *bind.TransactOpts, newuri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "setURI", newuri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_MyToken *MyTokenBulkTransactor) SetURI(opts *bind.TransactOpts, newuri string) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "setURI", newuri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_MyToken *MyTokenSession) SetURI(newuri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.SetURI(&_MyToken.TransactOpts, newuri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_MyToken *MyTokenTransactorSession) SetURI(newuri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.SetURI(&_MyToken.TransactOpts, newuri)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyToken *MyTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyToken *MyTokenBulkTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyToken *MyTokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.TransferFrom(&_MyToken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyToken *MyTokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.TransferFrom(&_MyToken.TransactOpts, from, to, tokenId)
}

// MyTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MyToken contract.
type MyTokenApprovalIterator struct {
	Event *MyTokenApproval // Event containing the contract specifics and raw log

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
func (it *MyTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenApproval)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyTokenApproval)
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
func (it *MyTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenApproval represents a Approval event raised by the MyToken contract.
type MyTokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// MyTokenApprovalOrChainReorg represents a Approval subscription event raised by the MyToken contract.
type MyTokenApprovalOrChainReorg struct {
	Event      *MyTokenApproval
	ChainReorg *types.ChainReorg
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MyToken *MyTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MyTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, err := _MyToken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenApprovalIterator{contract: _MyToken.contract, event: "Approval", logs: logs}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MyToken *MyTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MyTokenApprovalOrChainReorg, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyToken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenApprovalOrChainReorg)
				event.Event = new(MyTokenApproval)

				if log.ChainReorg == nil {
					if err := _MyToken.contract.UnpackLog(event.Event, "Approval", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MyToken *MyTokenFilterer) ParseApproval(log types.Log) (*MyTokenApproval, error) {
	event := new(MyTokenApproval)
	if err := _MyToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MyToken contract.
type MyTokenApprovalForAllIterator struct {
	Event *MyTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MyTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenApprovalForAll)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyTokenApprovalForAll)
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
func (it *MyTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenApprovalForAll represents a ApprovalForAll event raised by the MyToken contract.
type MyTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// MyTokenApprovalForAllOrChainReorg represents a ApprovalForAll subscription event raised by the MyToken contract.
type MyTokenApprovalForAllOrChainReorg struct {
	Event      *MyTokenApprovalForAll
	ChainReorg *types.ChainReorg
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MyToken *MyTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MyTokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, err := _MyToken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenApprovalForAllIterator{contract: _MyToken.contract, event: "ApprovalForAll", logs: logs}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MyToken *MyTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MyTokenApprovalForAllOrChainReorg, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MyToken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenApprovalForAllOrChainReorg)
				event.Event = new(MyTokenApprovalForAll)

				if log.ChainReorg == nil {
					if err := _MyToken.contract.UnpackLog(event.Event, "ApprovalForAll", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MyToken *MyTokenFilterer) ParseApprovalForAll(log types.Log) (*MyTokenApprovalForAll, error) {
	event := new(MyTokenApprovalForAll)
	if err := _MyToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyTokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MyToken contract.
type MyTokenRoleAdminChangedIterator struct {
	Event *MyTokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MyTokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenRoleAdminChanged)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyTokenRoleAdminChanged)
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
func (it *MyTokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenRoleAdminChanged represents a RoleAdminChanged event raised by the MyToken contract.
type MyTokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// MyTokenRoleAdminChangedOrChainReorg represents a RoleAdminChanged subscription event raised by the MyToken contract.
type MyTokenRoleAdminChangedOrChainReorg struct {
	Event      *MyTokenRoleAdminChanged
	ChainReorg *types.ChainReorg
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MyToken *MyTokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MyTokenRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, err := _MyToken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenRoleAdminChangedIterator{contract: _MyToken.contract, event: "RoleAdminChanged", logs: logs}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MyToken *MyTokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MyTokenRoleAdminChangedOrChainReorg, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MyToken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenRoleAdminChangedOrChainReorg)
				event.Event = new(MyTokenRoleAdminChanged)

				if log.ChainReorg == nil {
					if err := _MyToken.contract.UnpackLog(event.Event, "RoleAdminChanged", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MyToken *MyTokenFilterer) ParseRoleAdminChanged(log types.Log) (*MyTokenRoleAdminChanged, error) {
	event := new(MyTokenRoleAdminChanged)
	if err := _MyToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyTokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MyToken contract.
type MyTokenRoleGrantedIterator struct {
	Event *MyTokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *MyTokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenRoleGranted)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyTokenRoleGranted)
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
func (it *MyTokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenRoleGranted represents a RoleGranted event raised by the MyToken contract.
type MyTokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// MyTokenRoleGrantedOrChainReorg represents a RoleGranted subscription event raised by the MyToken contract.
type MyTokenRoleGrantedOrChainReorg struct {
	Event      *MyTokenRoleGranted
	ChainReorg *types.ChainReorg
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MyToken *MyTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MyTokenRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, err := _MyToken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenRoleGrantedIterator{contract: _MyToken.contract, event: "RoleGranted", logs: logs}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MyToken *MyTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MyTokenRoleGrantedOrChainReorg, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MyToken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenRoleGrantedOrChainReorg)
				event.Event = new(MyTokenRoleGranted)

				if log.ChainReorg == nil {
					if err := _MyToken.contract.UnpackLog(event.Event, "RoleGranted", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MyToken *MyTokenFilterer) ParseRoleGranted(log types.Log) (*MyTokenRoleGranted, error) {
	event := new(MyTokenRoleGranted)
	if err := _MyToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyTokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MyToken contract.
type MyTokenRoleRevokedIterator struct {
	Event *MyTokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MyTokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenRoleRevoked)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyTokenRoleRevoked)
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
func (it *MyTokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenRoleRevoked represents a RoleRevoked event raised by the MyToken contract.
type MyTokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// MyTokenRoleRevokedOrChainReorg represents a RoleRevoked subscription event raised by the MyToken contract.
type MyTokenRoleRevokedOrChainReorg struct {
	Event      *MyTokenRoleRevoked
	ChainReorg *types.ChainReorg
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MyToken *MyTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MyTokenRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, err := _MyToken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenRoleRevokedIterator{contract: _MyToken.contract, event: "RoleRevoked", logs: logs}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MyToken *MyTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MyTokenRoleRevokedOrChainReorg, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MyToken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenRoleRevokedOrChainReorg)
				event.Event = new(MyTokenRoleRevoked)

				if log.ChainReorg == nil {
					if err := _MyToken.contract.UnpackLog(event.Event, "RoleRevoked", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MyToken *MyTokenFilterer) ParseRoleRevoked(log types.Log) (*MyTokenRoleRevoked, error) {
	event := new(MyTokenRoleRevoked)
	if err := _MyToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyToken contract.
type MyTokenTransferIterator struct {
	Event *MyTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MyTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenTransfer)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyTokenTransfer)
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
func (it *MyTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenTransfer represents a Transfer event raised by the MyToken contract.
type MyTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// MyTokenTransferOrChainReorg represents a Transfer subscription event raised by the MyToken contract.
type MyTokenTransferOrChainReorg struct {
	Event      *MyTokenTransfer
	ChainReorg *types.ChainReorg
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MyToken *MyTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MyTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, err := _MyToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenTransferIterator{contract: _MyToken.contract, event: "Transfer", logs: logs}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MyToken *MyTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyTokenTransferOrChainReorg, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenTransferOrChainReorg)
				event.Event = new(MyTokenTransfer)

				if log.ChainReorg == nil {
					if err := _MyToken.contract.UnpackLog(event.Event, "Transfer", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MyToken *MyTokenFilterer) ParseTransfer(log types.Log) (*MyTokenTransfer, error) {
	event := new(MyTokenTransfer)
	if err := _MyToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
