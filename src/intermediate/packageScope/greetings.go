package main

import "fmt"

var num int32 = 12 //this variable is available everywhere in this package "main" since it's not inside any function

// num := 12 not allowed because it's outside function's body.

//function greet is availabe everywhere in this package "main" but variables decraled inside it are not available in other files.
func greet(name string) string {

	response  :=  fmt.Sprintf("Hello %v It's been %v days since we lastly talked.", name, num)

	return response

}