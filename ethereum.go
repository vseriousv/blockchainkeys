package blockchainkeys

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (e *EthereumNetwork) GenerateKeyPair() (privateKey string, publicKey string, address string, err error) {
	// Generation privateKey
	private, err := crypto.GenerateKey()
	if err != nil {
		fmt.Printf("Error generation privateKey: %v", err)
		return "", "", "", err
	}

	// Get publicKey from privateKey
	public := private.Public()

	// PublicKey to ECDSA format
	publicKeyECDSA, ok := public.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Error format publicKey to ECDSA")
		return "", "", "", err
	}

	privateKey = hexutil.Encode(crypto.FromECDSA(private))
	publicKey = hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex() // Get Address from PublicKey

	return privateKey, publicKey, address, nil
}
