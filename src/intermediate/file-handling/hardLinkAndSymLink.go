package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	/*
	   Hard Links and Symlinks
	*/

	err := os.Link("demo.txt", "demo_also.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("creating a sym")

	// Create Symlink
	err = os.Symlink("demo.txt", "demo_also.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Lstat returns file information.
	// Symlink doesn't work on Windows.
	fileInfo, err := os.Lstat("demo_sym.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connection Info: %+v", fileInfo)

	// Just change ownership of a symlink.
	// But it does not change the ownership of the file it points to.
	err = os.Lchown("demo_sym.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}