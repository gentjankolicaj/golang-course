//go:build ignore

package main

import "fmt"

//Closure : Enclosing a variable in some scope,can be done by code-blocks
//We want to keep the scope of variable as small as possible

var x int //Scope of x is this whole package,outside package x is undefined

func main() {
	var y int //Scope of y is function main,outside main y is undefined
	fmt.Println(y)

	{
		var z int // Scope of z is inside this block,outside code-block z is undefined
		fmt.Println(z)
	}

	a := inc()       //Returned value of inc() is actually a function of type func() int{}
	fmt.Println(a()) //Call function assigned at a
	fmt.Println(a())
	fmt.Println(a())

	//Note : variable a & b are on different memory locations
	b := inc()
	fmt.Println()
	fmt.Println(b())
	fmt.Println(b())

}

func inc() func() int {
	var x int = 101 //Scope of x is inside inc() function.X is also visible on anon-function inside foo

	return func() int {
		x++ //Access variable x
		return x
	}
}
