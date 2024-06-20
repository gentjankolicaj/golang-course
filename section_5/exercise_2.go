//go:build ignore

package main

import "fmt"

//Variables with package level scope
var x int
var y string
var z bool
var a complex128

func main() {

	//Zero values assigned by default
	fmt.Println(x, y, z, a)

}
