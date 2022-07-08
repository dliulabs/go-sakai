# go-sakai

# Lab 1: How Blocks are generated

```
type Block struct {
	timestamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
}
```

# Lab 2: How to create a Blockchain?

```
type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
	port              uint16
	mux               sync.Mutex
}
```
