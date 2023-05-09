/*
Channels are another important feature of Go concurrency.
They provide a way for goroutines to communicate with each other and synchronize their execution.
Here's an example:
*/

package main

import (
    "fmt"
)

func main() {
    // create an unbuffered channel of integers
    ch := make(chan int)

    // start a new goroutine that sends a value on the channel
    go func() {
        ch <- 42
    }()

    // receive the value from the channel in the main goroutine
    value := <-ch
    fmt.Println("received value:", value)
}
