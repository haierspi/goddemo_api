package conflux_chain

import (
	"math/big"
	"os"
	"time"

	"starfission_go_api/global"
	"starfission_go_api/pkg/path"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/Conflux-Chain/go-conflux-sdk/types/unit"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type NetType string

// MainNet 树图主网树图主网
const MainNet = NetType("https://main.confluxrpc.com")

// TestNet 树图测试网
const TestNet = NetType("https://test.confluxrpc.com")

func (c *NetType) String() string {
	return string(*c)
}

type ConfluxChain struct {
	Client   *sdk.Client
	AM       *sdk.AccountManager
	ChainID  uint32
	Token    *MyToken
	KeyStore string
}

type Address string

func NewAddress(addr string) *Address {
	a := Address(addr)
	return &a
}

func (a *Address) String() string {
	return string(*a)
}
func (a *Address) CFXAddress() cfxaddress.Address {
	addr := cfxaddress.MustNew(a.String())
	return addr
}
func (a *Address) CFXPointerAddress() *cfxaddress.Address {
	x := a.CFXAddress()
	return &x
}

type AccountBalance struct {
	Balnce *hexutil.Big
}

func (a *AccountBalance) FormatCFX() decimal.Decimal {
	return unit.NewDrip(a.Balnce.ToInt()).FormatCFX()
}
func (a *AccountBalance) FormatCFXFloat64() (float64, bool) {
	return a.FormatCFX().Float64()
}

// TransactionResult 交易结果
type TransactionResult struct {
	Hash    string
	Receipt *types.TransactionReceipt
}

func (a *TransactionResult) NftTokenId() int64 {
	return a.Receipt.Logs[0].Topics[3].ToCommonHash().Big().Int64()
}

// NewClient 创建
func NewClient(net NetType, keystore string) (*ConfluxChain, error) {
	c := &ConfluxChain{}
	var err error
	c.Client, err = sdk.NewClient(net.String())
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NewClient error")
	}
	status, err := c.Client.GetStatus()
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NewClient GetStatus error")
	}
	c.ChainID = uint32(status.ChainID)

	c.InitAM(keystore)

	return c, nil
}

func GetKeyStorePath(add ...string) string {
	keystorePath := ""
	if global.CONFIG != "configs/" {
		if !path.Exists(global.CONFIG + "conflux_keystore/") {
			pwdDir, _ := os.Getwd()
			keystorePath = pwdDir + "/" + "configs/conflux_keystore/"
		} else {
			keystorePath = global.CONFIG + "conflux_keystore/"
		}
	} else {
		if !path.Exists(global.ROOT + global.CONFIG + "conflux_keystore/") {
			pwdDir, _ := os.Getwd()
			keystorePath = pwdDir + "/" + global.CONFIG + "conflux_keystore/"
		} else {
			keystorePath = global.ROOT + global.CONFIG + "conflux_keystore/"
		}
	}
	if add != nil {
		for _, s := range add {
			keystorePath += s + "/"
		}
	}
	return keystorePath
}

// InitAM 初始化账户管理器
func (c *ConfluxChain) InitAM(keystore string) {
	c.KeyStore = keystore
	c.AM = sdk.NewAccountManager(keystore, c.ChainID)
	c.Client.SetAccountManager(c.AM)
}

// Exists 检查账户管理器内账户是否存在
func (c *ConfluxChain) Exists(addr *Address) bool {
	return c.AM.Exists(addr.CFXAddress())
}

// Import 账户导入
func (c *ConfluxChain) Import(addr *Address, privateKeyString string, passphrase string) (err error) {
	if !c.Exists(addr) {
		address, err := c.AM.ImportKey(privateKeyString, passphrase)
		if err != nil {
			return errors.Wrap(err, "ConfluxChain AccountManager Import error")
		}
		if address.MustGetBase32Address() != addr.String() {
			c.AM.Delete(address, passphrase)
			return errors.New("ConfluxChain AccountManager Account inconsistency error")
		}
	}
	return nil
}

// Unlock 解锁账户
func (c *ConfluxChain) Unlock(addr *Address, passphrase string) error {
	return c.AM.Unlock(addr.CFXAddress(), passphrase)
}

// ImportAndUnlock 账户导入AND解锁
func (c *ConfluxChain) ImportAndUnlock(addr *Address, privateKeyString string, passphrase string) error {
	var err error
	if !c.Exists(addr) {
		err = c.Import(addr, privateKeyString, passphrase)
		if err != nil {
			return errors.Wrap(err, "ConfluxChain ImportAndUnlock Import error")
		}
	}
	err = c.Unlock(addr, passphrase)
	if err != nil {
		return errors.Wrap(err, "ConfluxChain ImportAndUnlock Unlock error")
	}
	return nil
}

// GetEpochNumber GetEpochNumber
func (c *ConfluxChain) GetEpochNumber() (*hexutil.Uint64, error) {
	epochNumber, err := c.Client.GetEpochNumber()
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain GetEpochNumber error")
	}
	return types.NewUint64(epochNumber.ToInt().Uint64()), nil
}

// GetGasPrice 燃气价格
func (c *ConfluxChain) GetGasPrice() (*hexutil.Big, error) {
	gasPrice, err := c.Client.GetGasPrice()
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain GetGasPrice error")
	}
	return gasPrice, nil
}

// GetGasLimit 燃气价格限制
func (c *ConfluxChain) GetGasLimit(addr *Address) (uint64, error) {
	request := types.CallRequest{
		To:    addr.CFXPointerAddress(),
		Value: types.NewBigInt(1),
	}

	r, err := c.Client.EstimateGasAndCollateral(request)
	if err != nil {
		return 0, errors.Wrap(err, "ConfluxChain GetGasLimit error")
	}
	return r.GasLimit.ToInt().Uint64(), nil
}

// GetNextNonce 获取账户
func (c *ConfluxChain) GetNextNonce(addr *Address) (*hexutil.Big, error) {
	nextNonce, err := c.Client.GetNextNonce(addr.CFXAddress(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain GetGasPrice error")
	}
	return nextNonce, nil
}

// GetBalance 获取账户余额
func (c *ConfluxChain) GetBalance(addr *Address) (*AccountBalance, error) {
	receipt, err := c.Client.GetBalance(addr.CFXAddress())
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain GetBalance error")
	}
	a := &AccountBalance{Balnce: receipt}
	return a, nil
}

// SendTransaction 账户CFX转账
func (c *ConfluxChain) SendTransaction(from *Address, to *Address, value *hexutil.Big) (*TransactionResult, error) {

	utx, err := c.Client.CreateUnsignedTransaction(from.CFXAddress(), to.CFXAddress(), value, nil)
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain SendTransaction CreateUnsignedTransaction error")
	}

	hash, err := c.Client.SendTransaction(utx)
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain SendTransaction error")
	}
	t := &TransactionResult{}
	t.Hash = hash.String()

	receipt, err := c.Client.WaitForTransationReceipt(hash, time.Second)
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain SendTransaction WaitForTransationReceipt error")
	}
	t.Receipt = receipt

	return t, nil
}

// NftMint NFT MINT
func (c *ConfluxChain) NftMint(contractAddr *Address, from *Address, to *Address) (*TransactionResult, error) {

	EpochHeight, err := c.GetEpochNumber()
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NftTransfer GetEpochNumber error")
	}
	gasPrice, err := c.GetGasPrice()
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NftTransfer GetGasPrice error")
	}
	fromNonce, err := c.GetNextNonce(from)
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NftTransfer GetNextNonce error")
	}
	gasLimit, err := c.GetGasLimit(to)
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NftTransfer GetGasLimit error")
	}

	token, err := NewMyToken(contractAddr.CFXAddress(), c.Client)

	opts := &bind.TransactOpts{
		//Value:      types.NewBigInt(2e18),
		Nonce:        fromNonce,
		From:         from.CFXPointerAddress(),
		GasPrice:     gasPrice,
		Gas:          types.NewBigInt(gasLimit * 5),
		StorageLimit: types.NewUint64(500),
		//StorageLimit: types.NewUint64(x.StorageCollateralized.ToInt().Uint64()),
		EpochHeight: EpochHeight,
	}

	_, hash, err := token.Mint(opts, to.CFXPointerAddress().MustGetCommonAddress())

	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NftMint Mint error")
	}
	t := &TransactionResult{}
	t.Hash = hash.String()

	receipt, err := c.Client.WaitForTransationReceipt(*hash, time.Second)
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NftMint WaitForTransationReceipt error")
	}
	t.Receipt = receipt

	return t, nil
}

// NftTransfer 账户nft资产转移
func (c *ConfluxChain) NftTransfer(contractAddr *Address, from *Address, to *Address, tokenId int64) (*TransactionResult, error) {

	EpochHeight, err := c.GetEpochNumber()
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NftTransfer GetEpochNumber error")
	}
	gasPrice, err := c.GetGasPrice()
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NftTransfer GetGasPrice error")
	}
	fromNonce, err := c.GetNextNonce(from)
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NftTransfer GetNextNonce error")
	}
	gasLimit, err := c.GetGasLimit(to)
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NftTransfer GetGasLimit error")
	}

	token, err := NewMyToken(contractAddr.CFXAddress(), c.Client)

	opts := &bind.TransactOpts{
		//Value:      types.NewBigInt(2e18),
		Nonce:        fromNonce,
		From:         from.CFXPointerAddress(),
		GasPrice:     gasPrice,
		Gas:          types.NewBigInt(gasLimit * 5),
		StorageLimit: types.NewUint64(500),
		//StorageLimit: types.NewUint64(x.StorageCollateralized.ToInt().Uint64()),
		EpochHeight: EpochHeight,
	}

	_, hash, err := token.SafeTransferFrom(opts, from.CFXPointerAddress().MustGetCommonAddress(), to.CFXPointerAddress().MustGetCommonAddress(), big.NewInt(tokenId))

	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NftTransfer SafeTransferFrom error")
	}
	t := &TransactionResult{}
	t.Hash = hash.String()

	receipt, err := c.Client.WaitForTransationReceipt(*hash, time.Second)
	if err != nil {
		return nil, errors.Wrap(err, "ConfluxChain NftTransfer WaitForTransationReceipt error")
	}
	t.Receipt = receipt

	return t, nil
}
