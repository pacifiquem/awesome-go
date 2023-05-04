package main

import "fmt"

func main() {

	age := 16
	name := "pacifiquem@github.com"

	fmt.Println("Println method add a new line")
	fmt.Print("Print method doesn't ")
	fmt.Print("add a new line")

	fmt.Println("My name is: ", name, " And my age is ", age, ".") // printing string and a variable to the console by separating them with comma.

	//Printf("Formatted string") %_ = format specifier
	fmt.Printf("My name is %v And my age is %v \n", name, age) //%v to tell Printf method that it will insert a variable there.
	fmt.Printf("My name is %q And my age is %v \n", name, age) //%q to tell Printf method that it will add quotes around a variable.
	fmt.Printf("Type of age is: %T \n", age) //%T to tell print type of a variable.
	fmt.Printf("Type of age is: %f \n", 12.12) //%f to tell print float values on console.
	fmt.Printf("Type of age is: %.2f \n", 12.12) //%f to tell print limit numbers after deciaml point of float values on console.

	//fmt.Sprintf() // Save formatted string to a variable
	bio := fmt.Sprintf("My name is %q And my age is %v \n", name, age)
	fmt.Println(bio) // printing my saved formatted string.

}

// this is a quick reminder that Printf() method doesn't add a new line so, you have to do it by your self.