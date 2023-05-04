package main

import "fmt"

func main() {

	//for loop
	for i := 1; i <= 10; i++ { // initizalizer; condition; increment/decrement
		fmt.Println("for loop: value of i: ",i)
	}

	//while looop
	j := 1
	for j <= 10 { //only condition
		fmt.Println("while loop: value of j: ", j)
		j++
	}

	// Range loop
	arr := []string{"apple", "banana", "cherry"}
	for index, value := range arr {
		fmt.Printf("Range loop iteration #%d: %s\n", index, value)
	}

}