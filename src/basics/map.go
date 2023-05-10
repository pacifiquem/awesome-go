package main

import "fmt" // import the fmt package for printing to the console

func main() {
    // Declare an empty map with string keys and int values
    var ages map[string]int

    // Initialize the map using a map literal
    ages = map[string]int{
        "Alice": 25,
        "Bob":   30,
        "Eve":   28,
    }

    // Print the value associated with the key "Alice"
    fmt.Printf("Alice is %d years old\n", ages["Alice"])

    // Add a new key-value pair to the map
    ages["Charlie"] = 35

    // Check if a key is present in the map
    if age, ok := ages["Dave"]; ok {
        fmt.Printf("Dave is %d years old\n", age)
    } else {
        fmt.Println("Dave's age is unknown")
    }

    // Iterate over the keys and values in the map
    for name, age := range ages {
        fmt.Printf("%s is %d years old\n", name, age)
    }
}

// In this example, we declare an empty map with string keys and int values using the var keyword. We then initialize the map using a map literal with three key-value pairs.

// We print the value associated with the key "Alice" using fmt.Printf and the map indexing syntax (ages["Alice"]).

// We add a new key-value pair to the map using the indexing syntax (ages["Charlie"] = 35).

// We check if a key is present in the map using the indexing syntax with a conditional statement (if age, ok := ages["Dave"]; ok {...}).

// We iterate over the keys and values in the map using a for loop with the range keyword (for name, age := range ages {...}). Note that the order of the keys and values is not guaranteed in a map, so they may not be printed in the same order as they were added to the map.

// Note that maps in Golang are created using the map keyword, followed by the type of the keys and values enclosed in square brackets (map[string]int in this case). To access a value in a map, you use the indexing syntax with the key (ages["Alice"] in this case). To add a new key-value pair to a map, you simply assign a value to a new key (ages["Charlie"] = 35 in this case). To check if a key is present in a map, you use the indexing syntax with a conditional statement (if age, ok := ages["Dave"]; ok {...} in this case). And to iterate over the keys and values in a map, you use a for loop with the range keyword (for name, age := range ages {...} in this case).

// All keys must have the same type & all values must have the same time.




