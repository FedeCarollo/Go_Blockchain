package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/schnorr"
)

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

func NewTransaction(from, to []byte, amount float64) *Transaction {
	return &Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Signature: SignatureCheck{},
	}
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
