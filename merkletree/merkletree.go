package merkletree

import (
	"bytes"
	"crypto/sha256"
)

type Merkle struct {
	tree   [][]node
	hashes [][]byte
}

type node struct {
	hash []byte
}

type dir int

const (
	LEFT dir = iota
	RIGHT
)

type proofNode struct {
	hash      []byte
	direction dir
}

func joinHash(node1, node2 []byte) []byte {
	joined := bytes.Join([][]byte{node1, node2}, []byte{})
	hash := sha256.Sum256(joined)
	return hash[:]
}

func GenerateMerkleTree(hashes [][]byte) Merkle {
	merkleTree := Merkle{}
	merkleTree.hashes = hashes
	nodes := make([][]node, 1)
	nodes[0] = make([]node, len(hashes))
	for i, hash := range hashes {
		nodes[0][i] = node{hash}
	}
	current_level := nodes[0]
	for len(current_level) > 1 {
		dim := len(current_level) / 2
		if len(current_level)%2 != 0 {
			dim += 1
		}
		next_level := make([]node, dim)

		for i := 0; i < len(current_level); i += 2 {
			if i+1 < len(current_level) {
				next_level[i/2].hash = joinHash(current_level[i].hash, current_level[i+1].hash)
			} else {
				next_level[i/2].hash = current_level[i].hash
			}
		}
		nodes = append(nodes, next_level)
		current_level = next_level
	}
	merkleTree.tree = nodes

	return merkleTree
}

func (merkle *Merkle) GetRoot() []byte {
	return merkle.tree[len(merkle.tree)-1][0].hash
}

func (merkle *Merkle) generateProof(pos int) []proofNode {
	tree := merkle.tree
	proof := make([]proofNode, 0)
	for lvl := 0; lvl < len(tree)-1; lvl++ {
		if pos%2 == 0 {
			if pos+1 < len(tree[lvl]) {
				proof = append(proof, proofNode{tree[lvl][pos+1].hash, RIGHT})
			}
		} else {
			proof = append(proof, proofNode{tree[lvl][pos-1].hash, LEFT})
		}
		pos /= 2
	}
	return proof
}

func (merkle *Merkle) CheckValidity(hash []byte, pos int) bool {
	if pos >= len(merkle.tree[0]) || pos < 0 {
		return false
	}
	proof := merkle.generateProof(pos)
	total_hash := hash
	for _, proof_hash := range proof {
		if proof_hash.direction == LEFT {
			total_hash = joinHash(proof_hash.hash, total_hash)
		} else { //RIGHT
			total_hash = joinHash(total_hash, proof_hash.hash)
		}
	}

	return bytes.Equal(total_hash, merkle.GetRoot())
}
