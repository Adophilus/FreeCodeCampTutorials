package main

import (
	"fmt"
)

func main () {
	grades := []int{1, 2, 3}
	classes := []int{10, 20, 30}

	// just like in python
	classes2 := grades[:1]
	classes3 := grades[0:2] // [inclusive:exclusive] => [ 1, 2 ]

	// classes4 = grades
	classes4 := grades

	classes5 := append(classes3, 3, 4, 5)
	// retuns the number of elements presenlty in the slice
	fmt.Println(cap(grades))
	fmt.Println(len(classes))
	fmt.Println(len(classes2))
	fmt.Println(len(classes4))

	fmt.Println(classes3)
	fmt.Println(classes5)

}
