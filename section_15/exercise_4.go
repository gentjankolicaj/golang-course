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

type person struct {
	first string
	last  string
	age   int
}

//Since this function has receiver of type person=>This is a method
//of type person
func (p person) speak() {
	fmt.Println(p.first, p.last, "-", p.age)
}

func main() {

	p0 := person{
		first: "john",
		last:  "doe",
		age:   23233,
	}

	p1 := person{
		first: "jane",
		last:  "doe",
		age:   1000,
	}

	p0.speak()
	p1.speak()

}
