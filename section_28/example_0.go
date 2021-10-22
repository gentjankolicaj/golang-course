//go:build ignored
package main

import "fmt"

//Mult ... calculates product of variadic param values and returns it
func Mult(x ...int) int {
	result := 1
	for _, v := range x {
		result = result * v
	}
	return result
}

func main() {

	fmt.Println("Multiply 1,2,3,4,5,6", Mult(1, 2, 3, 4, 5, 6))
}
