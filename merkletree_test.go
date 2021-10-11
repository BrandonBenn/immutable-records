package main

import "testing"

func TestMerkleTreeBuild(t *testing.T) {
	expected := "12345678910910910910"
	hashes := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	leaves := make([]*Node, 0)

	for _, hash := range hashes {
		leaves = append(leaves, NewNode(hash, nil, nil))
	}

	tree := NewMerkleTree()
	tree.Build(leaves)
	result := tree.Root.Hash

	if result != expected {
		t.Errorf("Got %q, expected %q", result, expected)
	}
}

func TestGetFileNames(t *testing.T) {
	filenames, _ := getFileNames("docs/tests/articles")
	expected := 24
	result := len(filenames)

	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
}
