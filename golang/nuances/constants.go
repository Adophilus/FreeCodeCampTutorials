package main

import (
	"fmt"
)

const name string = "adophilus" // constant string

const (
	cat = iota // zero-value for this block of constants
	dog
)

const (
	_ = iota // throw away the zero-value
	male
	female
)

const (
	_ = iota + 2
	boy
	girl
)

const (
	light = 1 << iota
	dark
	yinYang
)

func main () {
	const calculation int = 0x1 >> 2 // not allowed
	fmt.Println(cat)
	fmt.Println(male)
	fmt.Println(girl)
	fmt.Println(light)
	fmt.Println(dark)
	fmt.Println(yinYang)
}
