package main

import (
	"goblockchain/block"
)

func main() {
	bc := block.NewBlockchain("http://localhost:5000", 8000)
	bc.Print()

	preHash := bc.LastBlock().Hash()
	bc.CreateBlock(5, preHash)
	bc.Print()

	preHash = bc.LastBlock().Hash()
	bc.CreateBlock(2, preHash)
	bc.Print()
}
