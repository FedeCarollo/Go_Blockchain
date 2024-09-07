package blockchain

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

type Blockchain struct {
	blocks []*Block
}

func (b *Blockchain) AddBlock(block *Block) error {
	if valid, err := block.IsValid(); !valid {
		return err
	}
	lastBlock := b.blocks[len(b.blocks)-1].Hash
	if !bytes.Equal(lastBlock, block.PrevBlockHash) {
		return fmt.Errorf("inconsistent block. Last hash in blockchain: %s, Last hash provided: %s",
			hex.EncodeToString(lastBlock),
			hex.EncodeToString(block.PrevBlockHash))
	}

	b.blocks = append(b.blocks, block)
	return nil
}

func (b *Blockchain) GetWalletAmount(address []byte) float64 {
	amount := 0.0
	for _, block := range b.blocks {
		amount += block.GetWalletAmount(address)
	}
	return amount
}

func (b *Blockchain) GetLastHash() []byte {
	return b.blocks[len(b.blocks)-1].Hash
}

func NewBlockChain() *Blockchain {
	return &Blockchain{
		blocks: []*Block{NewGenesisBlock()},
	}
}
