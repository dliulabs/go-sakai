# go-sakai

# Lab 1: How Blocks are generated

[NewBlock](https://github.com/ethereum/go-ethereum/blob/c503f98f6d5e80e079c1d8a3601d188af2a899da/core/types/block.go#L187)

```
type Block struct {
	timestamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
}
```

# Lab 2: How to create a Blockchain?

[Blockchain](https://github.com/ethereum/go-ethereum/blob/87bb5db675057d35ef5cbad4e4a64f50a7f06e7e/core/blockchain.go#L161)


```
type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
	port              uint16
	mux               sync.Mutex
}
```

# Lab 3: Creating the Hash of a Block

[sha256/Sum256](https://pkg.go.dev/crypto/sha256#example-Sum256)

`func Sum256(data []byte) [Size]byte`

```package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("%x", sum)
}
```