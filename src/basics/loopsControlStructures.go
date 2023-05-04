package main

import ("fmt")

func main() {

	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Printf("For loop iteration #%d\n", i)
	}

	// While loop
	j := 1
	for j <= 5 {
		fmt.Printf("While loop iteration #%d\n", j)
		j++
	}

	// Infinite loop with break statement
	k := 1
	for {
		if k > 5 {
			break
		}
		fmt.Printf("Infinite loop iteration #%d\n", k)
		k++
	}

	// Continue statement
	for l := 1; l <= 5; l++ {
		if l == 3 {
			continue
		}
		fmt.Printf("For loop iteration #%d\n", l)
	}

	// Nested loop with labeled break
OuterLoop:
	for m := 1; m <= 3; m++ {
		for n := 1; n <= 3; n++ {
			fmt.Printf("Nested loop iteration #%d-%d\n", m, n)
			if m == 2 && n == 2 {
				break OuterLoop
			}
		}
	}

	// Range loop
	arr := []string{"apple", "banana", "cherry"}
	for index, value := range arr {
		fmt.Printf("Range loop iteration #%d: %s\n", index, value)
	}
}
