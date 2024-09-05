package tests

import (
	"crypto/sha256"
	"fmt"
	"simple_blockchain/merkletree"
	"testing"
)

func TestMerkleTree(t *testing.T) {
	// Test code here
	hashes := make([][]byte, 0)
	for i := 0; i < 11; i++ {
		hashes = append(hashes, hashInt(i))
	}
	merkleTree := merkletree.GenerateMerkleTree(hashes)
	// fmt.Printf("Root hash: %x\n", merkleTree.GetRoot())

	if merkleTree.CheckValidity(hashes[2], 3) {
		t.Errorf("Merkle tree validity check failed")
	}

	if !merkleTree.CheckValidity(hashes[2], 2) {
		t.Errorf("Merkle tree validity check failed")
	}

}

func hashInt(i int) []byte {
	bytes := []byte(fmt.Sprintf("%d", i))
	h := sha256.Sum256(bytes)
	return h[:]
}
