package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"simple_blockchain/merkletree"
	"strconv"
	"time"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Transactions  []Transaction
	Hash          []byte
	MerkleHash    []byte
	Nonce         int64
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{
		b.PrevBlockHash,
		b.Data,
		timestamp,
		b.GenerateMerkleTree().GetRoot(),
		Int64ToBytes(b.Nonce)},
		[]byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func Int64ToBytes(n int64) []byte {
	// Create an 8-byte slice to hold the binary representation of the int64
	bytes := make([]byte, 8)

	// Use LittleEndian or BigEndian based on your preference
	binary.LittleEndian.PutUint64(bytes, uint64(n))

	return bytes
}

func (b *Block) GenerateMerkleTree() *merkletree.Merkle {
	return merkletree.GenerateMerkleTree(b.GetTransactionHashes())
}

func (b *Block) GetTransactionHashes() [][]byte {
	hashes := make([][]byte, len(b.Transactions))

	for i, transaction := range b.Transactions {
		hash := transaction.Hash()
		hashes[i] = hash
	}
	return hashes
}

func (b *Block) HashTransactions() []byte {
	hashes := make([][]byte, len(b.Transactions))

	for i, transaction := range b.Transactions {
		hash := transaction.Hash()
		hashes[i] = hash
	}

	hash := sha256.Sum256(bytes.Join(hashes, []byte{}))
	return hash[:]
}

func (b *Block) AddTransaction(t Transaction) {
	b.Transactions = append(b.Transactions, t)
}

func (b *Block) MineBlock(minerKey []byte, difficulty uint8) {

	pubKey, err := secp256k1.ParsePubKey(minerKey)

	if err != nil {
		log.Println("Error while parsing given public key")
	}
	transactions := make([]Transaction, 1+len(b.Transactions))
	transactions[0] = *NewTransaction(pubKey.SerializeCompressed(), []byte{}, 10.0)

	for i, transaction := range b.Transactions {
		transactions[i+1] = transaction
	}

	block := Block{
		Timestamp:     b.Timestamp,
		Data:          b.Data,
		Transactions:  transactions,
		PrevBlockHash: b.PrevBlockHash,
		Nonce:         0,
	}

	b.MerkleHash = block.GenerateMerkleTree().GetRoot()

	for {
		block.SetHash()

		if checkValidHashWithDifficulty(block.Hash, difficulty) {
			fmt.Printf("Found correct Nonce: %v with hash : %s", block.Nonce, block.Hash)
			b = &block
		}

		block.Nonce++
	}
}

func checkValidHashWithDifficulty(hash []byte, difficulty uint8) bool {
	for i := 0; i < int(difficulty); i++ {
		if hash[i] != 0 {
			return false
		}
	}
	return true
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Transactions:  []Transaction{},
		MerkleHash:    []byte{},
	}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis block", []byte{})
}
