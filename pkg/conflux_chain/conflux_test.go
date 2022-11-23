/**
  @author: HaierSpi
  @since: 2022/10/25
  @desc: //TODO
**/

package conflux_chain

import (
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/gookit/goutil/dump"
)

var testKeyStore = "../../configs/keystore"

func TestConfluxChain_Exists(t *testing.T) {
	c, err := NewClient(MainNet, testKeyStore)

	if err != nil {
		t.Errorf("ImportAndUnlock() NewClient error = %v", err)
	}

	c.Exists(NewAddress("cfx:"))

}

func TestConfluxChain_Import(t *testing.T) {
	c, err := NewClient(MainNet, testKeyStore)

	if err != nil {
		t.Errorf("Import() NewClient error = %v", err)
	}

	err = c.Import(NewAddress("cfx:"), "", "hello")

	//err = c.Import(NewAddress("cfx:"), "", "hello")

	if err != nil {
		t.Errorf("Import() error = %v", err)
	}

}

// unlock 解锁测试用例
func TestConfluxChain_Unlock(t *testing.T) {
	c, err := NewClient(MainNet, testKeyStore)

	if err != nil {
		t.Errorf("Unlock() NewClient error = %v", err)
	}

	err = c.Unlock(NewAddress("cfx:"), "hello")

	if err != nil {
		t.Errorf("Unlock() error = %v", err)
	}

	dump.P(c.AM)
}

func TestConfluxChain_ImportAndUnlock(t *testing.T) {

	c, err := NewClient(MainNet, testKeyStore)

	if err != nil {
		t.Errorf("ImportAndUnlock() NewClient error = %v", err)
	}

	err = c.ImportAndUnlock(NewAddress("cfx:"), "", "hello")

	if err != nil {
		t.Errorf("ImportAndUnlock() error = %v", err)
	}
}

func TestConfluxChain_GetBalance(t *testing.T) {
	c, err := NewClient(MainNet, testKeyStore)

	if err != nil {
		t.Errorf("GetBalance() NewClient error = %v", err)
	}

	r, err := c.GetBalance(NewAddress("cfx:"))

	if err != nil {
		t.Errorf("GetBalance() error = %v", err)
	}
	//x, _ := r.FormatCFX().Float64()
	dump.P(r.FormatCFX())
	dump.P(r.FormatCFXFloat64())
	t.Logf("GetBalance() Value  = %v", r)

}

func TestConfluxChain_SendTransaction(t *testing.T) {
	c, err := NewClient(MainNet, testKeyStore)

	if err != nil {
		t.Errorf("SendTransaction() NewClient error = %v", err)
	}

	from := NewAddress("cfx:")
	to := NewAddress("cfx:")
	num := types.NewBigInt(10000000000000)

	c.Unlock(from, "hello")

	r, err := c.SendTransaction(from, to, num)

	if err != nil {
		t.Errorf("SendTransaction() error = %v", err)
	}

	t.Logf("SendTransaction() Hash  = %v", r.Hash)
}

func TestConfluxChain_NftTransfer(t *testing.T) {
	c, err := NewClient(MainNet, testKeyStore)

	if err != nil {
		t.Errorf("NftTransfer() NewClient error = %v", err)
	}
	contract := NewAddress("cfx:")
	from := NewAddress("cfx:")
	to := NewAddress("cfx:")
	var tokenId int64 = 17

	c.Unlock(from, "hello")

	r, err := c.NftTransfer(contract, from, to, tokenId)

	if err != nil {
		t.Errorf("NftTransfer() error = %v", err)
	}
	dump.P(r)
}

func TestConfluxChain_NftMint(t *testing.T) {
	c, err := NewClient(MainNet, testKeyStore)

	if err != nil {
		t.Errorf("NftTransfer() NewClient error = %v", err)
	}
	contract := NewAddress("cfx:")
	from := NewAddress("cfx:")
	to := NewAddress("cfx:")

	c.Unlock(from, "hello")

	r, err := c.NftMint(contract, from, to)

	if err != nil {
		t.Errorf("NftMint() error = %v", err)
	}
	dump.P(r.NftTokenId())
}
