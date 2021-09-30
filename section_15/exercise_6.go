//go:build ignore

package main

import "fmt"

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

	anonFunc0 := func() {
		fmt.Println("AnonFunc0")
	}

	//Call anonFunc0
	anonFunc0()

	anonFunc1 := func(i int, v ...int) {
		fmt.Println("First argument ", i)
		fmt.Println("Second argument ", v)
	}

	//Call anonFunc1
	anonFunc1(101, 2, 3, 4, 4)

	//Self executing anon func
	func() {
		fmt.Println("Self executing anon func")
	}()

}
