package blockchainkeys

import (
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

// GenerateKeyPair - for waves MainNetScheme
func (w *WavesNetwork) GenerateKeyPair() (privateKey string, publicKey string, address string, err error) {
	seed, err := hdkeychain.GenerateSeed(hdkeychain.MinSeedBytes)
	if err != nil {
		return "", "", "", err
	}
	private, public, err := crypto.GenerateKeyPair(seed)
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
