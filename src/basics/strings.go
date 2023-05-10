package main

import (
    "fmt"
    "strings"
)

func main() {
    // Strings in Go are represented as UTF-8 encoded byte sequences.

    // Creating a new string:
    str := "Hello, world!"

    // Length of a string:
    fmt.Println(len(str)) // Output: 13

    // Indexing a string:
    fmt.Println(str[0]) // Output: 72 (the ASCII code for 'H')

    // Concatenating strings:
    str2 := " How are you?"
    str3 := str + str2
    fmt.Println(str3) // Output: Hello, world! How are you?

    // Checking if a string contains another string:
    fmt.Println(strings.Contains(str3, "world")) // Output: true

    // Counting the occurrences of a substring within a string:
    fmt.Println(strings.Count(str3, "o")) // Output: 3

    // Converting a string to uppercase or lowercase:
    fmt.Println(strings.ToUpper(str)) // Output: HELLO, WORLD!
    fmt.Println(strings.ToLower(str2)) // Output:  how are you?

    // Splitting a string into substrings based on a separator:
    str4 := "apple,banana,orange"
    fruits := strings.Split(str4, ",")
    fmt.Println(fruits) // Output: [apple banana orange]

    // Joining a slice of strings into a single string with a separator:
    joined := strings.Join(fruits, " and ")
    fmt.Println(joined) // Output: apple and banana and orange

    // Trimming leading and trailing whitespace from a string:
    str5 := "   some text   "
    trimmed := strings.TrimSpace(str5)
    fmt.Println(trimmed) // Output: "some text"

    // Replacing substrings within a string:
    str6 := "The quick brown fox jumps over the lazy dog"
    replaced := strings.Replace(str6, "fox", "cat", -1)
    fmt.Println(replaced) // Output: "The quick brown cat jumps over the lazy dog"
}
