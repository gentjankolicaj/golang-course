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

func main() {
	//define variable of type func
	var variable func()

	variable = func() {
		fmt.Println("Local anonymous func assigned to variable")
	}

	//Calling variable
	variable()

}
