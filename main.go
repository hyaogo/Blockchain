package main

import (
	"BlockchainProject/bolck"
	"fmt"
)

func main() {
	block := bolck.NewBlock("123", []byte("fjd0432kkfdsa"))
	fmt.Println(block)
}
