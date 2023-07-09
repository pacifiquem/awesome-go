/*
	Concurrency in Go is achieved with goroutines but by default we do not wait for the goroutines to finish.

	WaitGroups are used to wait for a collection of goroutines to finish executing. The control flow of the program is blocked until all goroutines finish executing. Lets see an example:
*/

package main

import (
	"fmt"
	"sync"
)

func doWork(id int, job int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	// mark the worker as done when the function returns
	defer wg.Done()

	for i := 0; i < job; i++ {
		fmt.Println("Worker", id, "processing job", i)
	}
}

func returnWork(id int, job int, wg *sync.WaitGroup, result chan int) {
	fmt.Printf("Worker %d starting\n", id)

	// mark the worker as done when the function returns
	defer wg.Done()

	for i := 0; i < job; i++ {
		fmt.Println("Worker", id, "processing job", i)
	}

	// send the result to the channel
	result <- id * 2
}

func main() {
	/*
		Basic WaitGroup example
	*/

	// define a WaitGroup
	var wg sync.WaitGroup

	fmt.Println("Starting goroutines")
	for i := 0; i < 3; i++ {
		// add a goroutine to the WaitGroup
		wg.Add(1)

		// launch a goroutine and pass in the WaitGroup as a pointer
		// see line 82 for explanation
		go doWork(i, 5, &wg)
	}
	// block the main thread until all goroutines finish executing
	wg.Wait()
	fmt.Println("End goroutines\n\n")

	/*
		Basic WaitGroup example with channels
	*/

	// create a channel to receive the result
	result := make(chan int, 2)

	fmt.Println("Starting goroutines w/ channels")
	for i := 0; i < 3; i++ {
		// add a goroutine to the WaitGroup
		wg.Add(1)

		// launch a goroutine and pass in the WaitGroup as a pointer
		// see line 82 for explanation
		go returnWork(i, 5, &wg, result)

		// print the result from the channel
		fmt.Println(<-result)
	}
	// block the main thread until all goroutines finish executing
	wg.Wait()
	fmt.Println("End goroutines w/ channels")

	/*
		Why pass WaitGroups as a pointer?

		Since WaitGroups keep track of the number of goroutines that are 
		running, we need to pass them as a pointer to the function.

		Not doing so will result in a copy of the WaitGroup being created and used inside the function, while the original WaitGroup will never receive any changes to the counter.

		Calling wg.Wait() as seen in line 80 will do nothing since our WaitGroup never got its counter incremented.
	*/
}
