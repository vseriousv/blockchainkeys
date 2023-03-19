package blockchain_keys

import "fmt"

type Blockchain interface {
	GenerateKeyPair() (err error)
}

func NewBlockchain(name string) (Blockchain, error) {
	switch name {
	case "Ethereum":
		return &Ethereum{}, nil
	case "Tron":
		return &Tron{}, nil
	case "Waves":
		return &Waves{}, nil
	default:
		return nil, fmt.Errorf("unknown blockchain: %s", name)
	}
}
