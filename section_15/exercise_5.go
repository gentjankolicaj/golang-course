//go:build ignore

package main

import (
	"fmt"
	"math"
)

//Function syntax : func (r receiver) identifier(params) (returns){}
//Methods : functions with receiver.
//Callback : function is passed as argument
//Defer : Defer function execution till closing function is returned
//Closure : Closing scope of variable in some scope
//Recursion : Function calling itself
//Interface : type on golang for defining behaviour
//Custom types on golang : keyword identifier type
//Anon func : functions without identifier
//Variadic params : A variable number of parameters in function
//Func expression : Assign to a variable a value of func

//Type & method
type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

//Type & method
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//Interface : Anything that has area() method is of type shape
type shape interface {
	area() float64
}

func main() {
	var a float64 = 5
	var b float64 = 7

	square := square{
		length: a,
	}
	circle := circle{
		radius: b,
	}

	//Print area of square
	info(square)

	//Print area of circle
	info(circle)

}

func info(s shape) {
	fmt.Println(s.area())
}
