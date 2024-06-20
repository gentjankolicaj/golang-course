//go:build ignore

package main

import "fmt"

var a float64 = 3.14

const (
	b = 12
)

func main() {
	fmt.Printf("a values on different numeral systems %d | %b | %x \n", a, a, a)
	fmt.Printf("b values on different numeral systems %d | %b | %x \n", b, b, b)
}
