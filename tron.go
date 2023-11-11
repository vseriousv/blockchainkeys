package blockchainkeys

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)

func (t *TronNetwork) GenerateKeyPair() (privateKey string, publicKey string, address string, err error) {
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

	// Get Address from PublicKey
	addrs := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	addrs = "41" + addrs[2:]
	addb, _ := hex.DecodeString(addrs)
	hash1 := s256(s256(addb))
	secret := hash1[:4]
	for _, v := range secret {
		addb = append(addb, v)
	}

	privateKey = hexutil.Encode(crypto.FromECDSA(private))
	publicKey = hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))
	address = base58.Encode(addb)

	return privateKey, publicKey, address, nil
}

func (t *TronNetwork) GetPublicKeyAndAddressByPrivateKey(privateKey string) (publicKey string, address string, err error) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", "", err
	}

	pb := pk.Public()
	publicKeyECDSA, ok := pb.(*ecdsa.PublicKey)
	if !ok {
		return "", "", fmt.Errorf("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyHex := hexutil.Encode(publicKeyBytes)[4:]

	addrs := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	addrs = "41" + addrs[2:]
	addb, _ := hex.DecodeString(addrs)
	hash1 := s256(s256(addb))
	secret := hash1[:4]
	for _, v := range secret {
		addb = append(addb, v)
	}

	publicKey = hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))
	address = base58.Encode(addb)
	return publicKeyHex, address, nil
}

func s256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	bs := h.Sum(nil)
	return bs
}
