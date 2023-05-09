/*Concurrency is an important feature of Go, and it's designed to make it easy to write concurrent programs.
	In Go, concurrency is achieved through goroutines and channels.
	Goroutines are lightweight threads of execution that are managed by the Go runtime.
	They allow multiple functions to execute simultaneously in the same address space.
	You can create a new goroutine by using the go keyword followed by the function call. Here's an example:
*/

package main

import (
    "fmt"
    "time"
)

func main() {
    // start a new goroutine that prints "world"
    go func() {
        time.Sleep(time.Second)
        fmt.Println("world")
    }()

    // print "hello"
    fmt.Println("hello")

    // wait for the goroutine to finish
    time.Sleep(time.Second * 2)
}
