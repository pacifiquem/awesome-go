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


// In this example, we declare a variable i with an initial value of 10 using the := operator. We then declare a pointer variable p using the & operator to get the memory address of i.

// We print the value of i using the variable itself (fmt.Printf("i = %d\n", i)) and the dereferenced pointer (fmt.Printf("*p = %d\n", *p)), which gives the same output of 10.

// We change the value of i using the pointer by assigning a new value to the dereferenced pointer (*p = 20). This also changes the value of i, since the pointer points to the same memory address as i.

// We then declare a new variable j and assign the value of i to j (j = i). This assigns the value of i to j, but does not create a new pointer to i.

// We print the value of j and i using the variable itself and the dereferenced pointer, which gives the same output of 20.

// We change the value of i using the variable by assigning a new value to i (i = 40). This also changes the value of the dereferenced pointer, but not the value of j.

// Note that pointers in Golang are declared using the * operator before the variable name (var p *int in this case). To get the memory address of a variable, you use the & operator (p = &i in this case). To dereference