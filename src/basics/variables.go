package main

import "fmt"

func main() {

	//Strings
	var nameOne string = "NameOne"
	var nameTwo = "NameTwo"
	var name3 string
	name3 = "NameThree"
	nameFour := "shorthand_of_`var nameFour string`" // this is always used inside a function only.

	fmt.Println(nameOne, nameTwo, name3, nameFour);

	nameOne = "newName"
	name3 = nameTwo
	nameTwo = nameOne
	nameFour = "AnotherName" // when you re-assign a variable you don't use :=

	fmt.Println(nameOne, nameTwo, name3, nameFour);

	//Integers
	var numOne int = 1
	var numTwo = 2
	var num3 int
	num3 = 3
	numFour := 4 // this is always used inside a function only.
	
	fmt.Println(numOne, numTwo, num3, numFour);
	
	numOne = 20
	num3 = numTwo
	numTwo = numOne
	numFour = 10 // when you re-assign a variable you don't use :=
	
	fmt.Println(numOne, numTwo, num3, numFour);

	var num5 int16 = 256 //interger of 16 bits, we can also have int8, int16, int32, int64
	var num6 uint = 122 //interger which can't hold a negative number, it can also be uint8, uint16, uint32, unint64
	fmt.Println(num5, num6)

	//Floats
	var scoreOne float32 = 32.23 // float of 32 bits, we always have to specify bits of our float always and another possible value is float64
	fmt.Println(scoreOne)

}

/*  
* Always remember that when you decrale variable and don't use it any where in the program it is going to give an error.
*/