package main

import (
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
)

func incorrect_usage_error() {
	fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
	os.Exit(1)
}

func cmd_init() {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
		}
	}

	headFileContents := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
	}

	fmt.Println("Initialized git directory")
}

func cmd_cat_file(arguments []string) {
	if len(arguments) < 4 || arguments[2] != "-p" {
		incorrect_usage_error()
	}

	sha_id := arguments[3]
	folder_name := sha_id[0:2]
	file_name := sha_id[2:]

	curr_dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory")
		return
	}

	file, err := os.Open(fmt.Sprintf("%s/.git/objects/%s/%s", curr_dir, folder_name, file_name))
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}
	defer file.Close()

	zr, err := zlib.NewReader(file)
	if err != nil {
		fmt.Println("Error decoding the file")
		return
	}

	decoded_content, err := ioutil.ReadAll(zr)
	if err != nil {
		fmt.Println("Error reading the decoded content")
		return
	}

	fmt.Printf("%s", string(decoded_content))
}

// Usage: your_git.sh <command> <arg1> <arg2> ...
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage!

	arguments := os.Args
	if len(arguments) < 2 {
		incorrect_usage_error()
	}

	switch command := arguments[1]; command {
	case "init":
		cmd_init()
	case "cat-file":
		cmd_cat_file(arguments)

	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
