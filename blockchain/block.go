package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
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

func (b *Block) MineBlock(minerPrivKey []byte, difficulty uint8, blockchain *Blockchain) error {
	privKey := secp256k1.PrivKeyFromBytes(minerPrivKey)
	pubKey := privKey.PubKey()

	// if err != nil {
	// 	log.Println("Error while parsing given private key")
	// }

	// pubKey, err := secp256k1.ParsePubKey(minerKey)

	transactions := make([]Transaction, 1+len(b.Transactions))
	transactions[0] = *NewTransaction(pubKey.SerializeCompressed(), []byte{}, 10.0, 0.0)
	transactions[0].Sign(privKey.Serialize())

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

	//TODO: Check block validity also withing blockchain
	if valid, err := block.IsValid(); !valid {
		return err
	}

	if len(blockchain.blocks) > 0 {
		if !bytes.Equal(blockchain.blocks[len(blockchain.blocks)-1].Hash, block.PrevBlockHash) {
			return fmt.Errorf("inconsistent block in blockchain")
		}
	}

	for {
		block.SetHash()

		if checkValidHashWithDifficulty(block.Hash, difficulty) {
			fmt.Printf("Found correct Nonce: %v with hash : %s\n", block.Nonce, hex.EncodeToString(block.Hash))
			b = &block
			break
		}

		block.Nonce++
	}

	blockchain.blocks = append(blockchain.blocks, &block)
	return nil
}

func (b *Block) IsValid() (bool, error) {
	if len(b.Transactions) < 2 { //At least 1 transaction + miner fee
		return false, errors.New("required at least one transaction in block")
	}

	if valid, err := b.validateMiner(); !valid {
		return false, err
	}

	for i := 1; i < len(b.Transactions); i++ { //Check for invalid transactions
		if !b.Transactions[i].IsValid() {
			return false, fmt.Errorf("transaction %s is invalid",
				hex.EncodeToString(b.Transactions[i].Hash()))
		}
	}
	difficulty := 1 //TODO: difficulty should be taken based on the last blocks mine time
	b.SetHash()

	hash := b.Hash

	if !checkValidHashWithDifficulty(hash, uint8(difficulty)) {
		return false, errors.New("Hash does not satisfy difficulty level")
	}

	return true, nil
}

func (b *Block) validateMiner() (bool, error) {
	if len(b.Transactions) == 0 {
		return false, errors.New("no transactions for provided block")
	}
	minerTransaction := b.Transactions[0]

	//TODO: other checks are needed?
	return ValidateMinerTransaction(minerTransaction)
}

func (b *Block) GetWalletAmount(address []byte) float64 {
	isMiner := bytes.Equal(b.Transactions[0].From, address) //first transaction is miner transaction
	amount := 0.0
	for _, tr := range b.Transactions {
		amount += tr.GetWalletAmount(address, isMiner)
	}
	return amount
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
