package main
import "fmt"

func main () {
	var a int
	var a1 uint8 // a byte
	var a2 int8
	var b float32
	var b1 float64
	var c complex64
	var c1 complex128
	a = 10
	a1 = 10
	a2 = 10
	b = 10
	a1 = 10
	c = 2
	c1 = 10
	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", a1, a1)
	fmt.Printf("%T, %v\n", a2, a2)
	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", b1, b1)
	fmt.Printf("%T, %v\n", c, c)
	fmt.Printf("%T, %v\n", c1, c1)
}
