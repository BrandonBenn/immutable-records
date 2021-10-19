package main

import (
	"fmt"
	"os"
)

const (
	ChecksumFile = "CHECKSUM"
	HELP         = `
NAME
irow - Immutable Records using One-Way Hash

SYNOPSIS
irow generate [directory]
irow verify   [directory]

DESCRIPTION
Given a directory of files, a snapshot of its contents is built using a Merkle
hash tree. The hash of root the is stored in a file called CHECKSUM. If any of
the files in the directory, or the CHECKSUM was tampered with, the verification
will give a warning.`
)

func check(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "\033[31;1;4mERROR\033[0m: "+e.Error())
		os.Exit(1)
	}
}

func getArguments() []string {
	return os.Args[1:]
}

func generate(directory string) {
	filenames, err := getFileNames(directory)
	check(err)

	leaves, err := createLeaves(filenames)
	check(err)

	tree := BuildMerkleTree(leaves)
	os.Remove(directory + "/" + ChecksumFile)
	err = os.WriteFile(directory+"/"+ChecksumFile, []byte(tree.Root.Hash), 0644)
	check(err)
}

func verify(directory string) {
	filenames, err := getFileNames(directory)
	check(err)

	leaves, err := createLeaves(filenames)
	check(err)

	tree := BuildMerkleTree(leaves)
	hash, err := os.ReadFile(directory + "/" + ChecksumFile)
	check(err)

	if string(hash) != tree.Root.Hash {
		fmt.Fprintln(os.Stderr, "\033[31;1;4mWARNING\033[0m: computed checksum did NOT match")
		os.Exit(1)
	}
	fmt.Printf("%s: \033[32;1;4mOK\033[0m\n", directory)
}

func main() {
	args := getArguments()
	if len(args) < 1 {
		fmt.Println(HELP)
		os.Exit(1)
	}

	switch letter := args[0][0]; letter {
	case 'g':
		generate(args[1])
	case 'v':
		verify(args[1])
	default:
		fmt.Println(HELP)
	}
}
