//go:build ignore

package main

import (
	"fmt"
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
//Variable of type func : var f func()

func main() {
	//Create a slice of func type
	var f []func() = make([]func(), 7, 10)

	for i := 0; i < len(f); i++ {
		f[i] = func() {
			//because i & tmp is enclosed,value of tmp & i the same on each func index of slice
			tmp := i
			fmt.Println(" i -", tmp)
		}
	}

	for j := 0; j < len(f); j++ {
		fmt.Printf("j - %v ", j)
		f[j]() //Invoke function in slice
	}
}
