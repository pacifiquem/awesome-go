package main

import "fmt"

func main() {
    // Create a new map with string keys and int values
    myMap := make(map[string]int)

    // Insert some key-value pairs into the map
    myMap["foo"] = 42
    myMap["bar"] = 23

    // Retrieve a value from the map
    val, ok := myMap["foo"]
    if ok {
        fmt.Printf("The value for key 'foo' is %d\n", val)
    } else {
        fmt.Println("Key 'foo' not found in map")
    }

    // Delete a key-value pair from the map
    delete(myMap, "bar")

    // Iterate over all key-value pairs in the map
    for key, value := range myMap {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}
