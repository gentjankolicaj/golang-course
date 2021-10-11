//go:build ignored

package main

import "fmt"

//Error: is of type interface,
//type error interface{
// Error() string
//}

//Any type that implements Error()string method ,is of type error.
//NOTE:Errors are just another type in golang
//NOTE:Any type that implements all methods of an interface is a type of that interface.

func main() {
	n, err := fmt.Println("Hello errors")
	if err == nil {
		fmt.Println("No error in fmt.Println()")
	} else {
		fmt.Println("Error on fmt.Println() ", err)
	}
	fmt.Println("N return from fmt.Println() ", n)
}
