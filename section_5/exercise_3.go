//go:build ignore

package main

import "fmt"

//Variables with package level scope
var x int = 42
var y string = "James Bond"
var z bool = true
var a complex128 = complex(1, 2)

func main() {
	s := fmt.Sprintf("%v\t%v\t%v\t%v", x, y, z, a)
	fmt.Println(s)
}
