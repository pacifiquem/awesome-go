package main

import (
    "bufio"     // import the bufio package for reading user inputs
    "fmt"       // import the fmt package for printing to the console
    "strconv"   // import the strconv package for converting string inputs to other types
)

func main() {
    // Create a new bufio reader for reading user inputs
    reader := bufio.NewReaderSize(os.Stdin, 1024)

    // Get a string input from the user
    fmt.Print("Enter a string: ")
    strInput, _ := reader.ReadString('\n')

    // Get an integer input from the user
    fmt.Print("Enter an integer: ")
    intInputStr, _ := reader.ReadString('\n')
    intInput, _ := strconv.ParseInt(intInputStr, 10, 64)

    // Get a float input from the user
    fmt.Print("Enter a float: ")
    floatInputStr, _ := reader.ReadString('\n')
    floatInput, _ := strconv.ParseFloat(floatInputStr, 64)

    // Get a boolean input from the user
    fmt.Print("Enter a boolean (true/false): ")
    boolInputStr, _ := reader.ReadString('\n')
    boolInput, _ := strconv.ParseBool(boolInputStr)

    // Print out the values of all the inputs
    fmt.Printf("String input: %s\n", strInput)
    fmt.Printf("Integer input: %d\n", intInput)
    fmt.Printf("Float input: %f\n", floatInput)
    fmt.Printf("Boolean input: %t\n", boolInput)
}
