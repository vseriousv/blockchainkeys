package blockchain_keys

import (
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

type Waves struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

// GenerateKeyPair - for waves MainNetScheme
func (w *Waves) GenerateKeyPair() error {
	privateKey, publicKey, err := crypto.GenerateKeyPair(nil)
	if err != nil {
		return err
	}

	address, err := proto.NewAddressFromPublicKey(proto.MainNetScheme, publicKey)
	if err != nil {
		return err
	}

	w.PrivateKey = privateKey.String()
	w.PublicKey = publicKey.String()
	w.Address = address.String()
	return nil
}
