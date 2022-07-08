package main

import (
	"fmt"
	"goblockchain/block"
	"goblockchain/wallet"
)

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// First, sender has to create a new transaction using his/her wallet
	// Wallet
	// A -> B: 2.0
	var value float32 = 2.0
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), value)

	// Second, miner adds transaction to the blockchain
	// Blockchain
	blockchain := block.NewBlockchain(walletM.BlockchainAddress(), 8000)
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), value,
		walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added? ", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))
}
