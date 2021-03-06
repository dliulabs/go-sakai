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

# Lab 4: Add a Transaction

- add transactions to the blockchain, new transactions will be stored in the "transaction pool"
- where creating a new block, transactions in the transaction pool are added to the new block.

[Transaction](https://github.com/ethereum/go-ethereum/blob/ae8ce7202244621d6e80eb69fcc31683fa0d4cea/core/types/transaction.go#L51)

[NewTransaction](https://github.com/ethereum/go-ethereum/blob/9244d5cd61f3ea5a7645fdf2a1a96d53421e412f/core/types/legacy_tx.go#L38)

```
type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}
```

# Lab 5: Mining

# Lab 6: Calculate Balance

# Lab 7: Private Key, Public Key, and ECDSA

ECDSA

- 

[Mastering Bitcoiin: Keys](https://www.oreilly.com/library/view/mastering-bitcoin/9781491902639/ch04.html)

[Key Generation](https://pkg.go.dev/crypto/ecdsa#example-package)

[Technical Background of Blockchain Address](https://wiki.bitcoinsv.io/index.php/Technical_background_of_Bitcoin_addresses)

- generate private key using `privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)`
- the PrivateKey has the following struct:

```
type PrivateKey struct {
	PublicKey
	D *big.Int
}
```

- to get the public key, all you need is to access the `.PublicKey` property.
- the public key has the following struct:

```
type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}
```

- when signing, the Sign() returns a pair of integers (a pair of keys, r & s):

```
func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)
```



```
import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	msg := "hello, world"
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", sig)

	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sig)
	fmt.Println("signature verified:", valid)
}
```

# Lab 8: Generating Signature for Transactions

A signature contains an R and an S part: this is consistent with the returned values of Sign():

`func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)`



```
type Transaction struct {
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}
```

```
type Signature struct {
	R *big.Int
	S *big.Int
}
```

# Lab 9: Verify transaction signature

- get the transaction bytes
- get the hash of the transaction
- use the `Verify()` to verify against sender's public key

```
func (bc *Blockchain) VerifyTransactionSignature(
	senderPublicKey *ecdsa.PublicKey, s *utils.Signature, t *Transaction) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}
```

- whenever adding a new transaction to the pool, we must verify signature is valid.