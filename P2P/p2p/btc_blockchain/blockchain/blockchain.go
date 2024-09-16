package blockchain

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Blockchain struct {
	Blocks []*Block
}

func (b *Blockchain) AddBlock(block *Block) error {
	if valid, err := block.IsValid(); !valid {
		return err
	}
	lastBlock := b.Blocks[len(b.Blocks)-1].Hash
	if !bytes.Equal(lastBlock, block.PrevBlockHash) {
		return fmt.Errorf("inconsistent block. Last hash in blockchain: %s, Last hash provided: %s",
			hex.EncodeToString(lastBlock),
			hex.EncodeToString(block.PrevBlockHash))
	}

	b.Blocks = append(b.Blocks, block)
	return nil
}

func (b *Blockchain) IsValid() (bool, error) {
	if len(b.Blocks) == 0 {
		return false, fmt.Errorf("no blocks in blockchain")
	}

	genesis := b.GetGenesis()

	if valid, err := genesis.IsGenesisValid(); !valid {
		return false, err
	}

	for i := 1; i < len(b.Blocks); i++ {
		if valid, err := b.Blocks[i].IsValid(); !valid {
			return false, err
		}
	}

	//TODO: should also check the hash. The genesis block should be stored somewhere
	return true, nil
}

func (b *Blockchain) GetGenesis() *Block {
	return b.Blocks[0]
}

func (b *Blockchain) GetWalletAmount(address []byte) float64 {
	amount := 0.0
	for _, block := range b.Blocks {
		amount += block.GetWalletAmount(address)
	}
	return amount
}

func (b *Blockchain) GetUserTransactions(address []byte) []Transaction {
	tr := make([]Transaction, 0)

	for _, block := range b.Blocks {
		tr = append(tr, block.GetUserTransactions(address)...)
	}
	return tr
}

func (b *Blockchain) GetLastHash() []byte {
	return b.Blocks[len(b.Blocks)-1].Hash
}

func NewBlockChain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{NewGenesisBlock()},
	}
}

func (b *Blockchain) GetLength() int {
	return len(b.Blocks)
}

func (b *Blockchain) Serialize() []byte {
	ser, _ := json.Marshal(b)

	return ser

}
