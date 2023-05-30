package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"github.com/pacifiquem/Go/utils"
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

		reGuess:

		//check available chances and print rules if it's the first time.
		utils.CheckChances(chanceNumber, gameMode)
		passed = GuessingNumbers()
		chanceNumber-- //reduce chance numbers after he plays

		if passed {
			fmt.Println("\n🎉🎉 Congratulations You won !!! 🎉🎉\n")
		}else {
			fmt.Println("😭 Wrong Guess, Try Again 😭")
			goto reGuess
		}

	} else if gameMode == "words" {

		GuessingWords()

	} else {

		fmt.Println("Please enter a valid Game mode");
		goto chooseGameMode

	}

}