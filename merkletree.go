package main

import (
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

type MerkleTree struct {
	Root *Node
}

func NewMerkleTree() *MerkleTree {
	return &MerkleTree{}
}

func (t MerkleTree) String() string {
	return fmt.Sprintf("%v", t.Root.Hash)
}

// In-place algorithm for generating Merkle Tree structure. Build takes an
// array of leaf nodes that contain the hashes of each document in the
// directory.
func (t *MerkleTree) Build(leaves []*Node) {
	size := len(leaves)
	for size != 1 {
		for i, left, right := 0, 0, 1; i < size && left < size; i, left, right = i+1, left+2, left+3 {
			if right >= size {
				right = left
			}
			hash := leaves[left].Hash + leaves[right].Hash
			left_child, right_child := leaves[left], leaves[right]
			leaves[i] = NewNode(hash, left_child, right_child)
		}
		size = (size + 2 - 1) / 2
	}

	t.Root = leaves[0]
}

// Returns a list of paths for all the files in the directory, ignoring
// directories.
func getFileNames(directory string) ([]string, error) {
	filenames := []string{}

	err := filepath.Walk(directory,
		func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				filenames = append(filenames, path)
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return filenames, nil
}
