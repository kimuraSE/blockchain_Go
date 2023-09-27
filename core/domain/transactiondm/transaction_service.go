package transactiondm

import (
	"encoding/json"
	"fmt"
)

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchianAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchianAddress string  `json:"recipient_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchianAddress:    t.senderBlockchianAddress,
		RecipientBlockchianAddress: t.recipientBlockchianAddress,
		Value:                      t.value,
	})
}

func (t *Transaction) Print() {
	fmt.Println("*** Transaction ***")
	fmt.Printf("sender_blockchain_address: %s\n", t.senderBlockchianAddress)
	fmt.Printf("recipient_blockchain_address: %s\n", t.recipientBlockchianAddress)
	fmt.Printf("value: %f\n", t.value)
}
