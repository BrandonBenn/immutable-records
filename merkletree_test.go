package main

import (
	"testing"
)

func TestGetFileNames(t *testing.T) {
	filenames, _ := getFileNames("docs/tests/articles")
	expected := 24
	result := len(filenames)

	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
}

func TestCreateLeaves(t *testing.T) {
	filenames := []string{"docs/tests/articles/day_1/post_1"}
	leaves, _ := createLeaves(filenames)

	result := leaves[0].Hash
	expected := "df14eda6e74f15dbaad5974e15d90c9bdca9ba601527a93b8d8b6d472e868d16"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestMerkleTreeBuild(t *testing.T) {
	expected := "895df7eb0a86dcb4aff1d3ebefe8b69659b3bb5a2faaeb5ce341c3b4cfd6ebd3"
	hashes := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	leaves := make([]*Node, 0)

	for _, hash := range hashes {
		leaves = append(leaves, NewNode(hash, nil, nil))
	}

	tree := buildMerkleTree(leaves)
	result := tree.Root.Hash

	if result != expected {
		t.Errorf("Got %q, expected %q", result, expected)
	}
}
