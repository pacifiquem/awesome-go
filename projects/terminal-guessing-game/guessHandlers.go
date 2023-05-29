package main

import (
	"github.com/pacifiquem/Go/utils"
	"os"
	"bufio"
	"fmt"
	"strconv"
	// "error"
	// "github.com/tjarratt/babble"
)

func GuessingWords() bool {
	utils.PrintRules("words")
	return true
}

func GuessingNumbers() bool {

	//get random number in range
	numToGuess := utils.GetRandomNumber()

	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your guess (hint: number between 10 - 20): ")

	// Scan the next token (in this case, a number)
	scanner.Scan()

	// Retrieve the scanned text
	input := scanner.Text()
	
	//convert input string to an integer
	guess, _ := strconv.Atoi(input)

	return guess == numToGuess

}