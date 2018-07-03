package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("Send 1 BTC to Johnny")
	bc.AddBlock("Send 2 more BTC to Johnny")

	for _, block := range bc.blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
