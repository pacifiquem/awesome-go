package main

import "fmt" // import the fmt package for printing to the console

// Declare a package-level variable
var count int

// Declare a package-level function
func increment() {
    count++
}

func main() {
    // Call the increment function 5 times
    for i := 0; i < 5; i++ {
        increment()
    }

    // Print the value of the count variable
    fmt.Printf("The count is %d\n", count)
}


// Note that package-level scope in Golang means that the variable or function is accessible from any function within the same package. If you declare a variable or function with package-level scope, you do not need to pass it as an argument or import it from another package.