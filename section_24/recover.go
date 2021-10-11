//go:build ignored

package main

import "fmt"

//Just like try/catch block in exception in languages like Java, C#, etc. are used to catch exception
//similarly in Go language, recover function is used to handle panic.
//It is an inbuilt function which is defined under the builtin package of the Go language.
//The main use of this function is to regain control of a panicking Goroutine.
//Or in other words, it handles the panicking behavior of the Goroutine.
func main() {
	fx()
	fmt.Println("Returned normally from fx()")
}

func fx() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in fx()", r)
		}
	}()

	fmt.Println("Calling gx()")
	gx(0)
	fmt.Println("Returned normally from gx()")
}

func gx(i int) {
	if i > 3 {
		fmt.Println("Panicking !")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g ", i)
	gx(i + 1)
}
