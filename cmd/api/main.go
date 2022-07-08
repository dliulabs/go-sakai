package main

import (
	"goblockchain/block"
)

func main() {
	b := block.NewBlock(0, [32]byte{}, []*block.Transaction{})
	b.Print()
}
