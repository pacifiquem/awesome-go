package main

import (
	"log"
	"os"
)

func main() {

	/*
	   Rename and Move a File
	*/

	originalPath := "demo.txt"
	newPath := "test.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	// Try to do the migration with Rename()
}