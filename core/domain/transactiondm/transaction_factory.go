package transactiondm

func NewTransaction(s, r string, v float32) *Transaction {
	return &Transaction{
		senderBlockchianAddress:    s,
		recipientBlockchianAddress: r,
		value:                      v,
	}
}