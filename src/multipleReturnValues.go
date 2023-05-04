package main

import "fmt" // import the fmt package for printing to the console

// Define a function that takes an integer as an argument and returns its square and cube as separate values
func squareAndCube(x int) (int, int) {
    return x*x, x*x*x
}

func main() {
    // Call the squareAndCube function with argument 5, and print the result
    square, cube := squareAndCube(5)
    fmt.Printf("The square of 5 is %d, and the cube of 5 is %d\n", square, cube)
}
