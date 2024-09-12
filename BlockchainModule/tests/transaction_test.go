package tests

import (
	"simple_blockchain/blockchain"
	"testing"
)

func TestTransactionSigning(t *testing.T) {
	usr1 := blockchain.CreateUserWithParams(false, "")
	usr2 := blockchain.CreateUserWithParams(false, "")

	transaction := blockchain.Transaction{
		From:      usr1.GetUserId(),
		To:        usr2.GetUserId(),
		Amount:    10,
		Gas:       0,
		Signature: blockchain.SignatureCheck{},
	}

	transaction.Sign(usr1.PrivateKey.Serialize())

	if !transaction.Verify() {
		t.Errorf("Transaction should have been verified")
	}

	transaction.Sign(usr2.PrivateKey.Serialize())

	if transaction.Verify() {
		t.Errorf("Transaction should have not been verified")
	}
}
