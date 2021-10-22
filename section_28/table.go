//go:build ignored

package main

import (
	"fmt"
)

func main() {

	fmt.Println("1,2,3,4,5 => sum = ", Sum(1, 2, 3, 4, 5))
}

//Sum ... calculates the sum of variadic param values
func Sum(x ...int) int {
	var sum int = 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
