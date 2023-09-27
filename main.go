package main

import (
	"fmt"
	"udemy/core/domain/blockchaindm"
)

func main() {

	myBlockchainAddress := "my_blockchain_address"
	blockChain := blockchaindm.NewBlockchain(myBlockchainAddress)

	blockChain.AddTransaction("A", "B", 1.0)
	blockChain.Mining()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	blockChain.Mining()
	blockChain.Print()

	fmt.Printf("my_blockchain_address: %s, total_amount: %f\n", myBlockchainAddress, blockChain.CalculateTotalAmount(myBlockchainAddress))
	fmt.Printf("A: %f\n", blockChain.CalculateTotalAmount("A"))
	fmt.Printf("B: %f\n", blockChain.CalculateTotalAmount("B"))
	fmt.Printf("C: %f\n", blockChain.CalculateTotalAmount("C"))
	fmt.Printf("D: %f\n", blockChain.CalculateTotalAmount("D"))
	fmt.Printf("X: %f\n", blockChain.CalculateTotalAmount("X"))

}
