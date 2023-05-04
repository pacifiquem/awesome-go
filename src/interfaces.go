package main

import "fmt"

// Define an interface called "Animal" that requires a "Speak" method
type Animal interface {
    Speak() string
}

// Define a struct called "Dog" that implements the "Animal" interface
type Dog struct {
    Name string
}

// Implement the "Speak" method for the "Dog" struct
func (d Dog) Speak() string {
    return "Woof! My name is " + d.Name
}

// Define a struct called "Cat" that implements the "Animal" interface
type Cat struct {
    Name string
}

// Implement the "Speak" method for the "Cat" struct
func (c Cat) Speak() string {
    return "Meow! My name is " + c.Name
}

// Define a function that takes an "Animal" interface as an argument
func PrintAnimal(a Animal) {
    fmt.Println(a.Speak())
}

func main() {
    // Create a variable of type "Dog"
    myDog := Dog{Name: "Fido"}

    // Create a variable of type "Cat"
    myCat := Cat{Name: "Whiskers"}

    // Call the "PrintAnimal" function with the "myDog" variable
    PrintAnimal(myDog)

    // Call the "PrintAnimal" function with the "myCat" variable
    PrintAnimal(myCat)
}
