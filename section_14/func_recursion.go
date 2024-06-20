//go:build ignore

package main

import "fmt"

//Recursion : A function calls itself
//Template : func foo(i int) int{ if i==0 return 1 else i*foo(i-1)}
//Note : Important for each recursive function | method is to have a stop/return condition

func main() {
	rTotal := factorial(8)
	fmt.Println("\n")
	fmt.Println("Recursive ", rTotal)

	lTotal := factorialLoop(8)
	fmt.Println("Loop ", lTotal)

}

func factorial(n int) int {
	fmt.Println(n)
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func factorialLoop(n int) int {
	var total int = 1
	for i := n; i > 1; i-- {
		total = total * i
	}
	return total
}
