package main

import (
	"fmt"
	"strconv"
)

func main () {
	var a int = 42
	var b float32 = 3.142
	var c string = strconv.Itoa(a)
	fmt.Println(int(b))
	fmt.Println(float32(a))
	fmt.Printf("%v, %T", c, c)
}
