package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"github.com/pacifiquem/multiplication-table/utils"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number you want multiplication table for: ")
	scanner.Scan()
	input := scanner.Text()
	multiplierNumber, _ := strconv.Atoi(input)

	if multiplierNumber <= 0 {
		fmt.Println("Invalid input number")
	}else {
		utils.PrintTable(multiplierNumber)
	}

}