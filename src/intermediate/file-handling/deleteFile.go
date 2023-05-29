package main

import (
	"log"
	"os"
)

func main() {

	/*
	   Delete a file
	*/

	err := os.Remove("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
}