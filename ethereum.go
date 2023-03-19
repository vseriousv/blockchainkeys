package blockchain_keys

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Ethereum struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func (e *Ethereum) GenerateKeyPair() error {
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

	e.PrivateKey = hexutil.Encode(crypto.FromECDSA(privateKey))
	e.PublicKey = hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))
	e.Address = address

	return nil
}
