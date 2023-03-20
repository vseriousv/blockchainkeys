package blockchainkeys

import "fmt"

type BlockchainType string

const (
	Ethereum BlockchainType = "Ethereum"
	Tron     BlockchainType = "Tron"
	Waves    BlockchainType = "Waves"
)

type Blockchain interface {
	GenerateKeyPair() (privateKey string, publicKey string, address string, err error)
}

type EthereumNetwork struct{}
type TronNetwork struct{}
type WavesNetwork struct{}

func NewBlockchain(name BlockchainType) (Blockchain, error) {
	switch name {
	case Ethereum:
		return &EthereumNetwork{}, nil
	case Tron:
		return &TronNetwork{}, nil
	case Waves:
		return &WavesNetwork{}, nil
	default:
		return nil, fmt.Errorf("unknown blockchain: %s", name)
	}
}
