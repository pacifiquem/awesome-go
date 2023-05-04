package main

import "fmt" // import the fmt package for printing to the console

// Define a function that takes two integers as arguments and returns their sum
func add(a, b int) int {
    return a + b
}

func main() {
    // Call the add function with arguments 3 and 4, and print the result
    sum := add(3, 4)
    fmt.Printf("The sum of 3 and 4 is %d\n", sum)
}
