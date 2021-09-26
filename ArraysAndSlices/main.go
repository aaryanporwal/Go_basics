package main

import "fmt"

func main () {
	// Arrays
		// Actual declaration: var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30} // Type inferred

	names:=[4]string{"yoshi", "mario", "peach", "apple"}
	names[1] = "luigi"

	fmt.Println(ages, len(ages)) // Works!
	fmt.Println(names, len(names))

	// Slices
	var scores = []int{100,50,60}
	scores[2] = 25
	scores = append(scores, 23)
	fmt.Println(scores)

	// slice ranges
	range1 := names[1:3]
	range2 := names[2:]
	range3 := names[:3] // from 0th index(inclusive) to 3rd index(exclusive)
	fmt.Println(range1, range2, range3)
}
