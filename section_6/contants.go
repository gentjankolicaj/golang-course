//go:build ignore

package main

import (
	"fmt"
)

//Constants are entities that don't change value.
//Constants on golang use 'const' keyword

//Typed constants where type of variable is declared
const a int = 10

const b float64 = 3.14
const c string = "John Doe"

//Untyped constants where type of variable is undeclared
const (
	d = 222.22
	e = 91919
	f = true
)

func main() {
	PrintTyped()
	PrintUntyped()

}

func PrintTyped() {
	fmt.Println("Typed constants", a, b, c)
}

func PrintUntyped() {
	fmt.Println("Untyped constants", d, e, f)
}
