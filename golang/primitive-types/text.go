package main

import "fmt"

func main () {
	var name string = "adophilus"
	var initial rune = 'a' // used for encoding UTF-32 characters
	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", initial, initial)
}
