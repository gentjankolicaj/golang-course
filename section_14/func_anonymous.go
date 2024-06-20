//go:build ignore

package main

import "fmt"

//NOTE: Anonymous functions don't have identifier
//Anonymous function :  func (){} => no identifier
//Anonymous functions are invoked at places where they are defined
// func(){ ....}()  =>Anonymous func run

func main() {

	//Defer anonymous function,
	//Function is executed because of () after {}
	defer func() {
		fmt.Println("Anon func with no params")
	}()

	func(i int) {
		fmt.Println("Anon func with params ", i)
	}(99101)

}

func print(s *string) {
	fmt.Println(s)
}
