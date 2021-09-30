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

	a := []int{2, 3, 4, 5, 6, 7, 7}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(foo(a...)) //Unfurled slice because foo() accepts int variadic params
	fmt.Println(bar(b))

}

func foo(variadic ...int) int {
	var sum int = 0
	for _, v := range variadic {
		sum = sum + v
	}
	return sum
}

func bar(slice []int) int {
	var sum int = 0
	for _, v := range slice {
		sum = sum + v
	}
	return sum
}
