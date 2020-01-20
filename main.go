package main

import (
	"blockchain/blockchain"
	"fmt"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First")
	chain.AddBlock("Second")

	for _, block := range chain.Blocks {

		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
}
