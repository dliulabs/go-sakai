package main

import (
	"fmt"
	"goblockchain/block"
)

func main() {
	myBlockchainAddress := "Miner-M"
	bc := block.NewBlockchain(myBlockchainAddress, 8000)
	bc.Print()

	bc.AddTransaction("A", "B", 1.0, nil, nil) // A -> B: 1.0
	/*
		preHash := bc.LastBlock().Hash()
		nonce := bc.ProofOfWork()
		bc.CreateBlock(nonce, preHash)
	*/
	bc.Mining()
	bc.Print()

	bc.AddTransaction("C", "D", 2.0, nil, nil) // C -> D: 2.0
	bc.AddTransaction("X", "Y", 3.0, nil, nil) // X -> Y: 3.0
	/*
		preHash = bc.LastBlock().Hash()
		nonce = bc.ProofOfWork()
		bc.CreateBlock(nonce, preHash)
	*/
	bc.Mining()
	bc.Print()

	fmt.Printf("my %.1f\n", bc.CalculateTotalAmount(myBlockchainAddress))
	fmt.Printf("C %.1f\n", bc.CalculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", bc.CalculateTotalAmount("D"))
}
