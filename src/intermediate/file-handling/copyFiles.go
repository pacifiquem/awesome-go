package main

import (
	"io"
	"log"
	"os"
)

func main() {

	/*
	   Copy a file
	*/

	originalFile, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	newFile, err := os.Create("demo_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	// Process file content(commit)
	// Free memory to disk(flush)
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}