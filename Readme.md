# Blockchain Keys Library

The Blockchain Keys Library is a Go library designed to generate private keys, public keys, and addresses for different blockchains, such as Ethereum, Tron, and Waves.

## Installation

To install the library, run the following command:

```shell
$ go get github.com/vseriousv/blockchainkeys
```

## Usage

To use the library, import it in your Go project and create a new instance of the desired blockchain:

```go
package main

import (
	"fmt"
	"github.com/vseriousv/blockchainkeys"
)

func main() {
	bc, err := blockchainkeys.NewBlockchain("Ethereum")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	privateKey, publicKey, address, err := bc.GenerateKeyPair()
	if err != nil {
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Private Key:", privateKey)
	fmt.Println("Public Key:", publicKey)
	fmt.Println("Address:", address)
}

```

You can replace "Ethereum" with "Tron" or "Waves" to generate keys and addresses for those blockchains.

## Supported Blockchains
Currently, the library supports the following blockchains:

- Ethereum
- Tron
- Waves

You can extend the library to support more blockchains by implementing the Blockchain interface and adding a new case to the NewBlockchain function.