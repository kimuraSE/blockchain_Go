package blockdm

import "udemy/core/domain/transactiondm"

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*transactiondm.Transaction
}
