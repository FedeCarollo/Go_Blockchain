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

	genesis := blockchain.NewGenesisBlock()

	genesis.SetHash()

	block := blockchain.NewBlock("Test block", genesis.Hash)
	transaction := blockchain.NewTransaction(usr.GetUserId(), usr2.GetUserId(), 1)

	block.AddTransaction(*transaction)

	block.MineBlock(usr.GetPublicKey().SerializeUncompressed(), 1)

	// fmt.Printf("User ID: %x\n", usr.GetUserId())
}
