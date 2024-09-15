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

func (b *Blockchain) IsValid() (bool, error) {
	if len(b.blocks) == 0 {
		return false, fmt.Errorf("no blocks in blockchain")
	}

	genesis := b.GetGenesis()

	if valid, err := genesis.IsGenesisValid(); !valid {
		return false, err
	}

	for i := 1; i < len(b.blocks); i++ {
		if valid, err := b.blocks[i].IsValid(); !valid {
			return false, err
		}
	}

	//TODO: should also check the hash. The genesis block should be stored somewhere
	return true, nil
}

func (b *Blockchain) GetGenesis() *Block {
	return b.blocks[0]
}

func (b *Blockchain) GetWalletAmount(address []byte) float64 {
	amount := 0.0
	for _, block := range b.blocks {
		amount += block.GetWalletAmount(address)
	}
	return amount
}

func (b *Blockchain) GetUserTransactions(address []byte) []Transaction {
	tr := make([]Transaction, 0)

	for _, block := range b.blocks {
		tr = append(tr, block.GetUserTransactions(address)...)
	}
	return tr
}

func (b *Blockchain) GetLastHash() []byte {
	return b.blocks[len(b.blocks)-1].Hash
}

func NewBlockChain() *Blockchain {
	return &Blockchain{
		blocks: []*Block{NewGenesisBlock()},
	}
}
