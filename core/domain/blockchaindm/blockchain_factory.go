package blockchaindm

import "udemy/core/domain/blockdm"

func NewBlockchain(ba string) *Blockchain {
	b := &blockdm.Block{}
	bc := &Blockchain{}
	bc.BlockchainAddress = ba
	bc.CreateBlock(0, b.Hash())
	return bc
}