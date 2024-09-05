package main

import (
	"fmt"
	"simple_blockchain/user"
)

func main() {
	user.CreateUser()
	usr := user.GetUserFromFile("private.pem")
	usr.PrintKeys()

	fmt.Printf("User ID: %x\n", usr.GetUserId())
}
