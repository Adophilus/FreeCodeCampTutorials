package main

import "fmt"

func main () {
	var a int = 10
	var (
		b int = 20
		c int = 30
	)
	d := 40
	var e int = 50
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	// fmt.Println(e) // raises compile time error because all variables created must be used
}
