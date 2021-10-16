package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
)

type Node struct {
	Hash  string
	Left  *Node
	Right *Node
}

func NewNode(hash string, left, right *Node) *Node {
	return &Node{
		hash,
		left,
		right,
	}
}

func (t Node) String() string {
	return fmt.Sprintf("%v", t.Hash)
}

type MerkleTree struct{ Root *Node }

func NewMerkleTree(head *Node) *MerkleTree {
	return &MerkleTree{Root: head}
}

func (t MerkleTree) String() string {
	return fmt.Sprintf("%v", t.Root.Hash)
}

// In-place algorithm for generating Merkle Tree structure. Build takes an
// array of leaf nodes that contain the hashes of each document in the
// directory.
func BuildMerkleTree(leaves []*Node) *MerkleTree {
	// Hash values are then repeatedly pair-wise merge-hashed till a single
	// root hash value is produced.
	size := len(leaves)
	for size != 1 {
		for i, left, right := 0, 0, 1; i < size && left < size; i, left, right = i+1, left+2, left+3 {
			if right >= size {
				// When the  # of hash values  is not even, the last hash value is
				// reused to create a pair.
				right = left
			}
			left_child, right_child := leaves[left], leaves[right]

			hash := mergeHash(right_child.Hash, left_child.Hash)
			leaves[i] = NewNode(hash, left_child, right_child)
		}
		size = (size + 2 - 1) / 2
	}

	return NewMerkleTree(leaves[0])
}

func mergeHash(hash1, hash2 string) string {
	return generateHash([]byte(hash1 + hash2))
}

// Takes an array of bytes, returns SHA256 string representation of content.
func generateHash(content []byte) string {
	hash := sha256.Sum256(content)
	return fmt.Sprintf("%x", hash[:])
}

// creates leaf nodes containing the hashes corresponding to each file in
// specified directory.
func createLeaves(filenames []string) []*Node {
	leaves := make([]*Node, 0)

	for _, file := range filenames {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}

		hash := generateHash(content)
		leaves = append(leaves, NewNode(hash, nil, nil))
	}

	return leaves
}

// Returns a list of paths for all the files in the directory, ignoring
// directories.
func getFileNames(directory string) ([]string, error) {
	filenames := []string{}

	err := filepath.Walk(directory,
		func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() && f.Name() != ChecksumFile {
				filenames = append(filenames, path)
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return filenames, nil
}
