package main

import (
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {

	/*
	   Check if File Exists
	*/
	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
	}
	log.Println("File does exist. File information : ")
	log.Println(fileInfo)
}