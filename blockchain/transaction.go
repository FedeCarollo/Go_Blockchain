package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/schnorr"
)

const genesisAmount float64 = 1000.0

type Transaction struct {
	From      []byte
	To        []byte
	Amount    float64
	Gas       float64
	Signature SignatureCheck
}

type SignatureCheck struct {
	Signature       []byte
	SenderPublicKey []byte
}

func (t *Transaction) Sign(sender_private_key []byte) {
	hash := HashTransaction(*t)

	privKey := secp256k1.PrivKeyFromBytes(sender_private_key)

	signature, err := schnorr.Sign(privKey, hash)

	if err != nil {
		panic(err)
	}

	serialSign := signature.Serialize()

	t.Signature.Signature = serialSign
	t.Signature.SenderPublicKey = privKey.PubKey().SerializeUncompressed()
}

func (t *Transaction) Verify() bool {
	hash := HashTransaction(*t)

	if bytes.Equal([]byte{}, t.Signature.SenderPublicKey) {
		return false
	}

	pubKey, err := secp256k1.ParsePubKey(t.Signature.SenderPublicKey)

	if err != nil {
		panic(err)
	}

	signature, err := schnorr.ParseSignature(t.Signature.Signature)

	if err != nil {
		panic(err)
	}

	//hashes check if hashes match
	if !signature.Verify(hash, pubKey) {
		return false
	}

	// sender address must match the signer public key compressed
	return bytes.Equal(pubKey.SerializeCompressed(), t.From)
}

func (t *Transaction) VerifyMiner() bool {
	hash := HashTransaction(*t)

	if bytes.Equal([]byte{}, t.Signature.SenderPublicKey) {
		return false
	}

	pubKey, err := secp256k1.ParsePubKey(t.Signature.SenderPublicKey)

	if err != nil {
		panic(err)
	}

	signature, err := schnorr.ParseSignature(t.Signature.Signature)

	if err != nil {
		panic(err)
	}

	//hashes check if hashes match
	if !signature.Verify(hash, pubKey) {
		return false
	}

	// sender address must match the miner public key compressed
	return bytes.Equal(pubKey.SerializeCompressed(), t.To)
}

func (t *Transaction) IsValid() bool {
	//TODO: check for each sender if it has enough money and transaction > 0
	verified := t.Verify()
	return verified
}

// If the address is the miner of the block it should also take the gas fee
func (t *Transaction) GetWalletAmount(address []byte, miner bool) float64 {
	amount := 0.0
	if miner {
		amount += t.Gas
	}
	if bytes.Equal(t.From, address) {
		amount -= t.Amount
	} else if bytes.Equal(t.To, address) {
		amount += t.Amount
	}
	return amount
}

func (t *Transaction) IsGenesisValid() (bool, error) {
	//TODO: check also for other fields?
	creator := readPrivateKeyFromFile("private.key").PubKey().SerializeCompressed()
	if !bytes.Equal(creator, t.To) {
		return false, fmt.Errorf("invalid receiver for genesis block")
	}
	if !bytes.Equal([]byte{}, t.From) || t.Amount != genesisAmount {
		return false, fmt.Errorf("invalid format for genesis block")
	}
	return true, nil
}

func ValidateMinerTransaction(t Transaction) (bool, error) {
	if !bytes.Equal(t.From, []byte{}) {
		return false, errors.New("miner transaction should have no source")
	}

	minerFee := 10.0

	if t.Amount != minerFee { //TODO: dynamic miner fee to add
		return false, fmt.Errorf("miner transaction should have an amount of %v", minerFee)
	}

	// if !t.Verify() {
	// 	return false, fmt.Errorf("transaction %s is invalid", hex.EncodeToString(t.Hash()))
	// }

	return true, nil

}

func NewTransaction(from, to []byte, amount float64, gas float64) *Transaction {
	return &Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Gas:       gas,
		Signature: SignatureCheck{},
	}
}

func GenerateGenesisTransaction() *Transaction {
	creator := GetUserFromFile("private.key")
	priv := creator.PrivateKey
	return NewTransaction([]byte{}, priv.PubKey().SerializeCompressed(), genesisAmount, 0)
}

func HashTransaction(t Transaction) []byte {
	joined := bytes.Join([][]byte{t.From, t.To, convertFloatToByte(t.Amount), convertFloatToByte(t.Gas)}, []byte{})
	hash := sha256.Sum256(joined)
	return hash[:]
}

func (t *Transaction) Hash() []byte {
	return HashTransaction(*t)
}

func convertFloatToByte(f float64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, math.Float64bits(f))
	return bytes
}
