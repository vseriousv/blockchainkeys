package blockchain_keys

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)

type Tron struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func (t *Tron) GenerateKeyPair() error {
	// Generation privateKey
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Printf("Error generation privateKey: %v", err)
		return err
	}

	// Get publicKey from privateKey
	publicKey := privateKey.Public()

	// PublicKey to ECDSA format
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Error format publicKey to ECDSA")
		return err
	}

	// Get Address from PublicKey
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	address = "41" + address[2:]
	addb, _ := hex.DecodeString(address)
	hash1 := s256(s256(addb))
	secret := hash1[:4]
	for _, v := range secret {
		addb = append(addb, v)
	}

	t.PrivateKey = hexutil.Encode(crypto.FromECDSA(privateKey))
	t.PublicKey = hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))
	t.Address = base58.Encode(addb)
	return nil
}

func s256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	bs := h.Sum(nil)
	return bs
}
