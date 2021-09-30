//go:build ignore

package main

import "fmt"

//NOTE: A func expression a function assigned to a variable
// a:=func(){} => For calling function a()
//In Golang functions are first class citizens,means that functions can do what any other type can do.
//NOTE: Anonymous functions don't have identifier
func main() {

	f := func() {
		fmt.Println("Anon function assigned to variable f")
	}

	f() // calling anon function assigned to f

	funcVar := func(i int, name string) {
		fmt.Println("Argument value ", i, "Func variable name ", name)
	}

	//Call function
	funcVar(22, "funcVar")
}
