package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pacifiquem/go/utils"
)

func main() {

	gameMode := ""
	passed := false
	chanceNumber := 10
	reader := bufio.NewReader(os.Stdin)

chooseGameMode:

	fmt.Print("\nChoose Game Mode (numbers, words): ")
	gameMode, _ = reader.ReadString('\n')

	gameMode = strings.TrimSpace(gameMode) // Remove leading/trailing whitespace

	if gameMode == "numbers" {

	reGuessNumbers:

		//check available chances and print rules if it's the first time.
		utils.CheckChances(chanceNumber, gameMode)
		passed = GuessingNumbers()
		chanceNumber-- //reduce chance numbers after he plays

		if passed {
			fmt.Println("\n🎉🎉 Congratulations You won !!! 🎉🎉\n")
		} else {
			fmt.Println("😭 Wrong Guess, Try Again 😭")
			goto reGuessNumbers
		}

	} else if gameMode == "words" {

	reGuessWords:

		//check available chances and print rules if it's the first time.
		utils.CheckChances(chanceNumber, gameMode)
		passed = GuessingWords()
		chanceNumber-- //reduce chance numbers after he plays

		if passed {
			fmt.Println("\n🎉🎉 Congratulations You won !!! 🎉🎉\n")
		} else {
			fmt.Println("😭 Wrong Guess, Try Again 😭")
			goto reGuessWords
		}

	} else {

		fmt.Println("Please enter a valid Game mode")
		goto chooseGameMode

	}

}
