package main

import "fmt" // import the fmt package for printing to the console

// define a struct called "person"
type person struct {
    name string
    age  int
}

// define a receiver function for the "person" struct
func (p person) sayHello() {
    fmt.Println("Hello, my name is", p.name, "and I am", p.age, "years old.")
}

func main() {
    // create a variable of type "person"
    p := person{name: "Alice", age: 30}

    // call the "sayHello" method on the "person" variable
    p.sayHello()
}

// Next, we define a receiver function for the person struct using the func keyword, followed by the name of the function (sayHello) and the type that the function is associated with ((p person)). This means that the sayHello function is associated with the person type and can be called on variables of that type. The function simply prints a message to the console using the fmt.Println function and the fields of the person struct.

// In the main function, we create a variable p of type person with the values of name set to "Alice" and age set to 30. We call the sayHello method on the p variable using the dot notation (p.sayHello()), which prints a message to the console.

// Note that in Golang, receiver functions are defined using the func keyword with the receiver type specified before the function name. Receiver functions can be used to define methods on a custom type, and can be called on variables of that type using the dot notation.