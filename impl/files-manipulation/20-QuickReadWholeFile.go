package main

import (
	"io"
	"log"
)

func main() {

	/*
	   Dosyayı Hızlı Okuma (Quick Read Whole File to Memory)
	*/

	// Dosyadan byte dilimine okuma yapma
	data, err := io.ReadFile("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)
}
