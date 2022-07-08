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

