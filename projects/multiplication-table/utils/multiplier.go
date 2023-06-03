package utils

import "fmt"

func PrintTable (number int) {
	
	for i := 1; i <= number; i++ {
		fmt.Printf("\t \t %d X %d = %d\n", number, i, i*number)
	}

}