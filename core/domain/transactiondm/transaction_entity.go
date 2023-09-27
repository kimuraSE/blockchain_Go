package transactiondm

type Transaction struct {
	senderBlockchianAddress    string
	recipientBlockchianAddress string
	value                      float32
}

func (t *Transaction) SenderBlockchianAddress() string {
	return t.senderBlockchianAddress
}

func (t *Transaction) RecipientBlockchianAddress() string {
	return t.recipientBlockchianAddress
}

func (t *Transaction) Value() float32 {
	return t.value
}
