package main

import (
	"crypto/sha256"
	"fmt"
	"simple_blockchain/merkletree"
	"simple_blockchain/user"
)

func main() {
	hashes := make([][]byte, 0)
	for i := 0; i < 11; i++ {
		hashes = append(hashes, hashInt(i))
	}
	merkleTree := merkletree.GenerateMerkleTree(hashes)
	fmt.Printf("Root hash: %x\n", merkleTree.GetRoot())

	fmt.Println(merkleTree.CheckValidity(hashes[2], 3))

	user.CreateUser()
	usr := user.GetUserFromFile("private.pem")
	usr.PrintKeys()

	fmt.Printf("User ID: %x\n", usr.GetUserId())
}

func hashInt(i int) []byte {
	bytes := []byte(fmt.Sprintf("%d", i))
	h := sha256.Sum256(bytes)
	return h[:]
}
