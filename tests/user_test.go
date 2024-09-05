package tests

import (
	"os"
	"simple_blockchain/user"
	"testing"
)

func TestGetUser(t *testing.T) {
	creating := false
	if _, err := os.Stat("private.pem"); os.IsNotExist(err) {
		creating = true
		user.CreateUser()

	}
	user.GetUserFromFile("private.pem")

	if creating {
		os.Remove("private.pem")
	}

}
