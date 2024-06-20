//go:build ignored

package main

import "fmt"

//Error: is of type interface,
//type error interface{
// Error() string
//}

//NOTE:Errors are just another type in golang
//NOTE:Any type that implements all methods of an interface is a type of that interface.

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func main() {

	err := New("Testing sample error")
	fmt.Println(err)
}
