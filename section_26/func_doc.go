//go:build ignored

//package provides some math functions
package main

import "fmt"

//Function to calculate sum of variadic params
func Sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

//NOTE:
//go doc <package>.<function> show on terminal the documentation
//godoc -http=:8080 starts an http server for on localhost for official go documentation

func main() {
	sum := Sum(1, 2, 3, 3, 4, 5, 5)
	fmt.Println(sum)
}
