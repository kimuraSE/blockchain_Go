package blockdm

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"udemy/core/domain/transactiondm"
)

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	h := sha256.Sum256([]byte(m))
	return h
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Timestamp    int64          `json:"timestamp"`
		Transactions []*transactiondm.Transaction `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
}

func (b *Block) Print() {
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("previousHash: %x\n", b.previousHash)
	fmt.Printf("timestamp: %d\n", b.timestamp)
	for _, t := range b.transactions {
		t.Print()
	}

}