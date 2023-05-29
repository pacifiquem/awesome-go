/*
To generate a random number within a specific range in Go (Golang), you can use the rand package from the standard library.
Here's an example that generates a random number between a given minimum and maximum value:
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialize the random number generator with a seed
	rand.Seed(time.Now().UnixNano())

	// Define the range
	min := 1
	max := 100

	// Generate a random number within the range
	randomNum := rand.Intn(max-min+1) + min

	// Print the random number
	fmt.Println(randomNum)
}


/*
* In the above code, we first import the necessary packages (fmt, math/rand, and time). We then use the time.Now().
* UnixNano() function to generate a unique seed for the random number generator. Next, we define the minimum and maximum values of the desired range. Finally, we use rand.
* Intn() to generate a random number within the range, and we add the minimum value to ensure the number falls within the desired range.
* The generated random number is then printed to the console.
*/