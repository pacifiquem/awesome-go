package main

import "fmt"

func main() {
    // create an empty map with string keys and integer values
    scores := make(map[string]int)

    // add some key-value pairs to the map
    scores["Alice"] = 90
    scores["Bob"] = 80
    scores["Charlie"] = 70

    // retrieve a value from the map
    aliceScore := scores["Alice"]
    fmt.Println("Alice's score is:", aliceScore)

    // check if a key exists in the map
    _, ok := scores["Bob"]
    if ok {
        fmt.Println("Bob's score exists in the map")
    }

    // iterate over all key-value pairs in the map
    for key, value := range scores {
        fmt.Printf("%s's score is %d\n", key, value)
    }
}
