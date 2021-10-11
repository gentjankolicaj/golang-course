//go:build ignored

package main

import "fmt"

//Error: is of type interface,
//type error interface{
// Error() string
//}

//Any type that implements Error()string method ,is of type error.

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
