package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"github.com/pacifiquem/Go/utils"
)

func main() {

	choice := ""
	passed := false
	chanceNumber := 10
	reader := bufio.NewReader(os.Stdin)

	chooseGameMode:

	fmt.Print("\nChoose Game Mode (numbers, words): ")
	choice, _ = reader.ReadString('\n')

	choice = strings.TrimSpace(choice) // Remove leading/trailing whitespace

	if choice == "numbers" {

		reGuess:
		passed = GuessingNumbers()

		if passed {
			fmt.Println("🎉🎉 Congratulations You won !!! 🎉🎉")
		}else {
			fmt.Println(" You lost try again later ")
			goto reGuess
		}

	} else if choice == "words" {

		GuessingWords()

	} else {

		fmt.Println("Please enter a valid Game mode");
		goto chooseGameMode

	}

}