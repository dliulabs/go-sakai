package main

import (
	"goblockchain/block"
)

func main() {
	bc := block.NewBlockchain("http://localhost:5000", 8000)
	bc.Print()

	bc.AddTransaction("A", "B", 1.0, nil, nil) // A -> B: 1.0
	preHash := bc.LastBlock().Hash()
	nonce := bc.ProofOfWork()
	bc.CreateBlock(nonce, preHash)
	bc.Print()

	bc.AddTransaction("C", "D", 2.0, nil, nil) // C -> D: 2.0
	bc.AddTransaction("X", "Y", 3.0, nil, nil) // X -> Y: 3.0
	preHash = bc.LastBlock().Hash()
	nonce = bc.ProofOfWork()
	bc.CreateBlock(nonce, preHash)
	bc.Print()
}
