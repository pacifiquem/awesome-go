Strings Package in Golang
=========================

The strings package in Go provides a collection of functions for manipulating and searching UTF-8 encoded strings. It offers a variety of operations for comparing, concatenating, splitting, and searching strings, as well as functions for converting between different string representations.

String Comparison
-----------------
To compare two strings lexicographically, use the Compare() function:

```Go
result := strings.Compare("apple", "banana")
if result < 0 {
    fmt.Println("apple is less than banana")
} else if result == 0 {
    fmt.Println("apple is equal to banana")
} else {
    fmt.Println("apple is greater than banana")
}
```
 
This will compare the two strings and print a corresponding message based on the outcome.

String Concatenation
--------------------
To combine multiple strings into a single string, use the + operator:

```Go
greeting := "Hello" + " " + "World"
fmt.Println(greeting)
```
 
This will concatenate the three strings and print the resulting string "Hello World".

String Splitting
----------------
To split a string based on a delimiter, use the Split() function:

```Go
words := strings.Split("This is a string", " ")
fmt.Println(words)
```
 
This will split the string based on spaces and store the resulting words in a slice.

String Searching
----------------
To find the occurrence of a substring within a string, use the Index() or LastIndex() functions:

```Go
position := strings.Index("Hello, World!", "World")
if position == -1 {
    fmt.Println("World was not found")
} else {
    fmt.Println("World found at position:", position)
}
```
 
This will search for the substring "World" in the string and print its position if found.

String Conversion
-----------------
To convert a string to lowercase or uppercase, use the ToLower() or ToUpper() functions:

```Go
lower_case := strings.ToLower("Hello, World!")
upper_case := strings.ToUpper("hello, world!")
fmt.Println(lower_case)
fmt.Println(upper_case)
```

This will convert the strings to lowercase and uppercase, respectively, and print the resulting strings.

For more in-depth information on the strings package, refer to the official documentation: https://pkg.go.dev/: https://pkg.go.dev/
