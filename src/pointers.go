package main

import "fmt" // import the fmt package for printing to the console

func main() {
    // Declare a variable i with an initial value of 10
    i := 10

    // Declare a pointer variable p that points to the memory address of i
    p := &i

    // Print the value of i using the variable itself and the dereferenced pointer
    fmt.Printf("i = %d\n", i)     // i = 10
    fmt.Printf("*p = %d\n", *p)   // *p = 10

    // Change the value of i using the pointer
    *p = 20

    // Print the new value of i using the variable itself and the dereferenced pointer
    fmt.Printf("i = %d\n", i)     // i = 20
    fmt.Printf("*p = %d\n", *p)   // *p = 20

    // Declare a new variable j
    j := 30

    // Assign the value of i to j
    j = i

    // Print the value of j and i using the variable itself and the dereferenced pointer
    fmt.Printf("i = %d\n", i)     // i = 20
    fmt.Printf("*p = %d\n", *p)   // *p = 20
    fmt.Printf("j = %d\n", j)     // j = 20

    // Change the value of i using the variable
    i = 40

    // Print the new value of i using the variable itself and the dereferenced pointer
    fmt.Printf("i = %d\n", i)     // i = 40
    fmt.Printf("*p = %d\n", *p)   // *p = 40
    fmt.Printf("j = %d\n", j)     // j = 20
}
