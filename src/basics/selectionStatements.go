package main

import "fmt"

func main() {

	var name string = "pac"
	
	//if
	if name == "pac" {
		fmt.Println("pacifiquem@github.com contributed to these codes")
	}

	//if-else
	num := 5
	if num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}


	//switch cases
	day := 3
	switch day {
		case 1:
			fmt.Println("Monday")
			break
		case 2:
			fmt.Println("Tuesday")
			break
		case 3:
			fmt.Println("Wednesday")
			break
		case 4:
			fmt.Println("Thursday")
			break
		case 5:
			fmt.Println("Friday")
			break
		default:
			fmt.Println("Weekend")
			break
	}

}