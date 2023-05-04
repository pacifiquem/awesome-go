package main

import "fmt"

/*
* The basic difference between a slice and an array in Go (golang) is that array have fixed length while slices doesn't 
* also slices uses arrays under the hoods
*
*/

func main() {

	// arrays always have fixed length. You can't change it after you provided it.
	var array [5]int = [5]int{1, 2, 3, 4, 5} //this specify that we want array of 5 elements and each element is an integer.
	var array2  = [5]int{1, 2, 3, 4, 5} //shorthand of what's written above.
	array = array2

	names  := [2]string{"Pacifique", "Bwenge"} //array of names which must have 2 element of type string.

	arraySize := len(array) //getting size of an array

	fmt.Println(array, arraySize) // printing array and it's length
	fmt.Println(names, len(names)) // printing names array and it's length
	fmt.Println(names[1]) //getting specific element. we start counting from 0, so names[1] will output Bwenge.


	// slices: they are like arrays but we can add new element remove element and it doesn't have fixed size.
	var scores = []int{100, 85}
	scores = append(scores, 10) //adding new element into a slice. Also remember that append function doesn't change the origin array that's why I re-assigned it to a new returned value
	fmt.Println(scores, len(scores))

	//slice range : getting a elements in array for a specific range
	rangeOne := array[1:3] // it includes starting index but exlude the last index
	rangeTwo := array[1:] // it includes everything from start index up to the end
	rangeThree := array[:3] // it includes everything from start index up to the end index excluded.
	fmt.Println(rangeOne, rangeTwo, rangeThree)

}