/*
In Go, type casting is the process of converting a value of one type to another type.
This can be useful when you need to convert between types in your program.
Here are some examples of type casting in Go:
*/

package main

import (
    "fmt"
    "strconv"
)

// Define a struct type Person with two fields: name and age.
type Person struct {
    name string
    age  int
}

func main() {
    // Example 1: Convert an int to a float64.
    x := 42
    y := float64(x)
    fmt.Printf("Example 1: x is of type %T and y is of type %T\n", x, y)

    // Example 2: Convert a string to an int.
    s := "42"
    i, _ := strconv.Atoi(s) // Convert s to an int using strconv.Atoi.
    fmt.Printf("Example 2: s is of type %T and i is of type %T\n", s, i)

    // Example 3: Convert a struct to an interface.
    p := Person{"Alice", 30}
    var iface interface{} = p // Convert p to an interface{} value.
    fmt.Printf("Example 3: p is of type %T and iface is of type %T\n", p, iface)
}

/*
In this program, we define a Person struct type with two fields: name and age. We then use three examples to demonstrate type casting:

We convert an integer value x to a float64 value y using the float64() function.
We convert a string value s to an integer value i using the strconv.Atoi() function.
We convert a Person value p to an interface{} value iface by assigning p to iface.
*/