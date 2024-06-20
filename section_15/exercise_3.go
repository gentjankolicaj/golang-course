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

	f1()
	defer f2() //By defer function f2() =>f2() is executed after main function has finished executing
	f3()
}

func f1() {
	fmt.Println("f1() invoked")
}

func f2() {
	fmt.Println("Deferred f2() invoked ")
}

func f3() {
	fmt.Println("f3() invoked")
}
