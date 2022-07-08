package main

import (
	"goblockchain/block"
)

func main() {
	bc := block.NewBlockchain("http://localhost:5000", 8000)
	hash := bc.LastBlock().Hash()
	bc.CreateBlock(5, hash)
	bc.Print()

}
