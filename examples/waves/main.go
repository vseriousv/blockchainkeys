package main

import (
	"fmt"
	"github.com/vseriousv/blockchainkeys"
)

func main() {
	bc, err := blockchainkeys.NewBlockchain(blockchainkeys.Waves)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	privateKey, publicKey, address, err := bc.GenerateKeyPair()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Private Key:", privateKey)
	fmt.Println("Public Key:", publicKey)
	fmt.Println("Address:", address)
}
