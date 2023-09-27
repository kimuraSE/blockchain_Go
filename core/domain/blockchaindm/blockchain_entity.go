package blockchaindm

import (
	"udemy/core/domain/blockdm"
	"udemy/core/domain/transactiondm"
)

type Blockchain struct {
	transactionPool []*transactiondm.Transaction
	chain           []*blockdm.Block
	BlockchainAddress string
}

