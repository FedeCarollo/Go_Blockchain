package main

import (
	"crypto/sha256"
	"fmt"
	"simple_blockchain/cryptography"
)

func main() {
	// bc := NewBlockChain()

	// bc.AddBlock("Sent 1 BTC to Ivan")
	// bc.AddBlock("Sent 2 BTC to Ivan")

	// for _, block := range bc.blocks {
	// 	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Println()
	// }
	// hashes := make([][]byte, 5)
	// for i := 0; i < 5; i++ {
	// 	hashes[i] = hashInt(i)
	// }
	// merkle := merkletree.GenerateMerkleTree(hashes)
	// // fmt.Println(merkle)

	// fmt.Println(merkle.CheckValidity(hashInt(0), 1))
	cryptography.GenerateKeyPair()

	privateKey := cryptography.ReadPrivateKeyFromFile("private.pem")
	publicKey := privateKey.PubKey()

	cryptography.PrintKeys(*privateKey, *publicKey)
}

func hashInt(i int) []byte {
	bytes := []byte(fmt.Sprintf("%d", i))
	h := sha256.Sum256(bytes)
	return h[:]
}
