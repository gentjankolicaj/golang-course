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
//Variable of type func : var f func()

func main() {

	//g variable holds a value of type func(s string) string
	g := greeting()

	//Invoke func
	str := g("Invoked ")
	fmt.Println(str)
}

func greeting() func(s string) string {
	var s string = "Greeting from func greeting()"
	fmt.Println(s)

	//Declare of variable f
	var f func(s string) string

	//Assign value to f
	f = func(s string) string {
		return "Anon-func " + s
	}
	return f
}
