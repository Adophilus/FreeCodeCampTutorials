package main

import (
	"fmt"
)

func main () {
	// array of ints that can hold at most 3 ints
	grades := [3]int{10, 20, 30}
	fmt.Printf("%v\n", grades)
	grades[0] = grades[0] >> 2
	grades[1] = grades[1] >> 2
	grades[2] = grades[2] >> 2

	// the size of the array is determined by the parameters passed into the array initialization
	classes := [...]int{1, 2, 3}
	fmt.Printf("%v\n", grades)
	fmt.Printf("%v\n", classes)
	fmt.Printf("No of classes: %v", len(classes))

	// classes = copy of grades
	classes := grades

	// classes = grades
	classes := &grades
}

