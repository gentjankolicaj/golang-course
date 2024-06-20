//go:build ignore

package main

import "fmt"

//Func template : func (r receiver) identifier(param1 type,....) (type1,...){}
//Variadic means 0 or n number of elements
//How to omit values returned by func : => we use underscore _

//Defer : A defer statement invokes a function whose execution is deferred
//to the moment the surrounding function returns

func main() {

	//Defer function f1(),this means f1 will be invoked after
	//Surrounding function of f1() is out of stack
	//In our case after main() has finished executing

	//Invoking & defer functions
	defer f1()

	f2()
	f3()

}

func f1() {
	fmt.Println("Function f1() invoked ")
}

func f2() {
	fmt.Println("Function f2() invoked ")
}

func f3() {
	fmt.Println("Function f3() invoked ")
}
