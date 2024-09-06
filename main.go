package main

import (
	"simple_blockchain/blockchain"
	"simple_blockchain/user"
)

func main() {
	// user.CreateUser()
	usr := user.GetUserFromFile("private.key")
	// usr.PrintKeys()

	usr2 := user.CreateUserWithParams(false, "")

	bchain := blockchain.NewBlockChain()

	block := blockchain.NewBlock("Test block", bchain.GetLastHash())
	transaction := blockchain.NewTransaction(usr.GetUserId(), usr2.GetUserId(), 1, 0)

	block.AddTransaction(*transaction)

	block.MineBlock(usr.GetPublicKey().SerializeUncompressed(), 4, bchain)

	// fmt.Printf("User ID: %x\n", usr.GetUserId())
}
