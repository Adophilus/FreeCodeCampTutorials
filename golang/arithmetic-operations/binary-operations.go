package main

import "fmt"

func main () {
	a := 10 // 1010
	b := 3 // 0011
	c := 8 // 2^3 // 1000
	fmt.Println(a & b) // and
	fmt.Println(a | b) // or
	fmt.Println(a ^ b) // xor
	fmt.Println(a &^ b) // and-not

	fmt.Println(c << 3) // 2^3 * 2^3
	fmt.Println(c >> 3) // 2^3 / 2^3
}
