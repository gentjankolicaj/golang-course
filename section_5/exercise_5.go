//go:build ignore

package main

import "fmt"

type type1 float64
type type2 type1

var x type1 = 1
var y float64 = -99

func performXOperations() {
	s := fmt.Sprintf("%v\t%T", x, x)
	fmt.Println(s)

	//Reassign value
	x = 111
	fmt.Println(fmt.Sprintf("%v\t%T", x, x))
}

func performYOperations() {
	var convertedValue float64 = float64(x) //converted type x to int (its underlying type)
	fmt.Println(y)

	//Reassign converted value
	y = convertedValue
	fmt.Println(y)
}

func main() {
	performXOperations()
	performYOperations()
}
