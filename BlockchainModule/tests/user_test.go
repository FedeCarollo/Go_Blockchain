package tests

import (
	"os"
	"simple_blockchain/blockchain"
	"testing"
)

func TestGetUser(t *testing.T) {
	creating := false
	if _, err := os.Stat("private.pem"); os.IsNotExist(err) {
		creating = true
		blockchain.CreateUser()

	}
	blockchain.GetUserFromFile("private.pem")

	if creating {
		os.Remove("private.pem")
	}

}
