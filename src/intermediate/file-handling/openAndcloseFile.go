package main

import (
	"log"
	"os"
)

func main() {

	/*
	   Open and Close Files
	*/

	file, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()	file, err = os.OpenFile("demo.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	/*

		   OpenFile() second parameters types;

		   os.O_RDONLY   : Only read
		   os.O_WRONLY   : Only write
		   os.O_RDWR     : Both read and write
		   os.O_APPEND   : Append to end of file
		   os.O_CREATE   : Create a file if it doesn't exist
		   os.O_TRUNC    : Truncate file

		   // Can also be used more than once

		   	-> os.O_CREATE|os.O_APPEND
			-> os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	*/
}