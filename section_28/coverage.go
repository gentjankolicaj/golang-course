//go:build ignored
package main

import "fmt"

//NOTE:To run coverage => go test -cover .
func main(){

	fmt.Println(Greet("King"))
    fmt.Println(Greet("Tony"))
}

func Greet(s string)string{
	return fmt.Sprint("Welcome ",s)
}