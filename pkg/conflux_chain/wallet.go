package conflux_chain

import (
	"crypto/ecdsa"
	"log"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func NewWallet() (ethAddress string, cfxMainBase32Address string, cfxTestBase32Address string, ethPrivateKey string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	//生成私钥

	ethPrivateKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	ethAddress = hexutil.Encode(hash.Sum(nil)[12:])

	cfxOldAddress := "0x1" + ethAddress[3:]
	cfxMainAddress := cfxaddress.MustNewFromHex(cfxOldAddress, 1029)
	cfxTestAddress := cfxaddress.MustNewFromHex(cfxOldAddress, 1)
	cfxMainBase32Address = cfxMainAddress.MustGetBase32Address()
	cfxTestBase32Address = cfxTestAddress.MustGetBase32Address()

	return ethAddress, cfxMainBase32Address, cfxTestBase32Address, ethPrivateKey

}
