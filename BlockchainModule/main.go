package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"simple_blockchain/blockchain"
)

func main() {
	// user.CreateUser()
	usr := blockchain.GetUserFromFile("private.key")
	// usr.PrintKeys()

	usr2 := blockchain.CreateUserWithParams(false, "")

	bchain := blockchain.NewBlockChain(filepath.Join(GetWd(), "private.key"))

	fmt.Println(bchain.IsValid())

	block := blockchain.NewBlock(bchain.GetLastHash())
	transaction := blockchain.NewTransaction(usr.GetUserId(), usr2.GetUserId(), 1, 0)
	transaction.Sign(usr.GetPrivateKey().Serialize())

	block.AddTransaction(*transaction)

	err := block.MineBlock(usr.GetPublicKey().SerializeUncompressed(), 1, bchain)

	if err != nil {
		fmt.Printf("Block was not added. Err: %v", err)
		return
	}

	fmt.Println(bchain.IsValid())

	// fmt.Printf("User ID: %x\n", usr.GetUserId())
}

func GetWd() string {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatalf("Error while getting working directory: %v", err)
	}
	return wd
}
