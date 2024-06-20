//go:build ignore

package main

import "fmt"

//Closure : Closing scope of variable

func main() {
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	g := foo() //new variable,new x value
	fmt.Println(g())
	fmt.Println()

}

func foo() func() int {
	x := 0
	//Variable x is enclosed by func() int{} ,which means that each time fun() int is invoked
	//Value of x is incremented because x variable is enclosed
	return func() int {
		x++
		return x
	}
}
