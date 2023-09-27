package blockchaindm

import (
	"fmt"
	"log"
	"strings"
	"udemy/core/domain/blockdm"
	"udemy/core/domain/transactiondm"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

func (bc *Blockchain) CreateBlock(n int, ph [32]byte) *blockdm.Block {
	b := blockdm.NewBlock(n, ph, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*transactiondm.Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *blockdm.Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) AddTransaction(s, r string, v float32) {
	t := transactiondm.NewTransaction(s, r, v)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*transactiondm.Transaction {
	transactions := make([]*transactiondm.Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, 
			transactiondm.NewTransaction(t.RecipientBlockchianAddress(), t.SenderBlockchianAddress(), t.Value()),
		)
	}
	return transactions
}

func (bc *Blockchain) Print() {
	for _, block := range bc.chain {
		fmt.Println("====================================")
		block.Print()
	}
}

func(bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*transactiondm.Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := blockdm.NewBlock(nonce, previousHash, transactions)
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.BlockchainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

func(bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.Transactions() {
			value := float32(t.Value())
			if blockchainAddress == t.RecipientBlockchianAddress() {
				totalAmount += value
			}
			if blockchainAddress == t.SenderBlockchianAddress() {
				totalAmount -= value
			}
		}
	}
	return float32(totalAmount)
}