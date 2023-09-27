package blockdm

import (
	"time"
	"udemy/core/domain/transactiondm"
)

func NewBlock(n int, ph [32]byte, t []*transactiondm.Transaction) *Block {
	return &Block{
		nonce:        n,
		previousHash: ph,
		timestamp:    time.Now().UnixNano(),
		transactions: t,
	}
}

func (b *Block) Nonce() int {
	return b.nonce
}

func (b *Block) PreviousHash() [32]byte {
	return b.previousHash
}

func (b *Block) Timestamp() int64 {
	return b.timestamp
}

func (b *Block) Transactions() []*transactiondm.Transaction {
	return b.transactions
}
