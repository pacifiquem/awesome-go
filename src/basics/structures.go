package main

import "fmt" // import the fmt package for printing to the console

// define a custom type called "money"
type money float64

// define a struct called "person"
type person struct {
    name string
    age  int
}

func main() {
    // create a variable of type "money"
    var m money = 100.50

    // print the value of the money variable
    fmt.Println("Money:", m)

    // create a variable of type "person"
    p := person{name: "Alice", age: 30}

    // print the values of the person variable
    fmt.Println("Person:", p)
    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
}

// In this example, we first define a custom type called money using the type keyword. The money type is essentially an alias for the float64 type, so we can create variables of type money and perform arithmetic operations on them just like regular float64 values.

// Next, we define a struct called person using the type keyword. The person struct has two fields, name and age.

// In the main function, we create a variable m of type money and assign it a value of 100.50. We print the value of m using the fmt.Println function.

// We then create a variable p of type person using the person struct definition. We set the value of name to "Alice" and the value of age to 30. We print the values of the p variable using the fmt.Println function, and also print the values of name and age using the dot notation (p.name and p.age, respectively).

// Note that in Golang, custom types and structs are defined using the type keyword. Custom types can be created by specifying the underlying type, while structs can be created by specifying a collection of fields. You can create variables of custom types and structs using the type or struct name, respectively. Fields in structs can be accessed using the dot notation.