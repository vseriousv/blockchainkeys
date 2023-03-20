package blockchainkeys

import (
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

// GenerateKeyPair - for waves MainNetScheme
func (w *WavesNetwork) GenerateKeyPair() (privateKey string, publicKey string, address string, err error) {
	private, public, err := crypto.GenerateKeyPair(nil)
	if err != nil {
		return "", "", "", err
	}

	addrs, err := proto.NewAddressFromPublicKey(proto.MainNetScheme, public)
	if err != nil {
		return "", "", "", err
	}

	privateKey = private.String()
	publicKey = public.String()
	address = addrs.String()
	return privateKey, publicKey, address, nil
}
