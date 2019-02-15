package main

import (
	"BlockchainProject/bolck"
	"fmt"
)

func main() {
	bc := bolck.NewBlockChain()
	bc.AddBlock("happy new year")
	bc.AddBlock("i love you")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PreBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
