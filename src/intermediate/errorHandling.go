type error interface {
    Error() string
}


package main

import (
    "fmt"
    "os"
)

func main() {
    // open a file
    f, err := os.Open("test.txt")

    // check for errors
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // do something with the file
    fmt.Println("File opened successfully")
    f.Close()
}


/*
In this example, we use the os.Open function to open a file called test.txt. This function returns two values: the file object and an error value.

We then check for errors by comparing the error value to nil. If an error occurred, we print out the error message using fmt.Println and return from the function.

If no error occurred, we print out a success message and close the file using the Close method.

This is a simple example of error handling in Go, but it demonstrates the basic approach. By convention, functions that can return an error value should always return the error as the last return value, and callers should check the error value and handle errors appropriately.
*/