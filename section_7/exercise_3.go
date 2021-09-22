//go:build ignore

package main

import "fmt"

//Type declaration
type FLAG bool
type STATE int

//Typed constants
const (
	ON_FLAG    FLAG  = true
	USER_STATE STATE = 0
)

//Untyped constants
const (
	a = 2
	b = 3
)

func main() {
	fmt.Printf("Typed constant %v , %T \n", ON_FLAG, ON_FLAG)
	fmt.Printf("Typed constant %v , %T \n", USER_STATE, USER_STATE)
	fmt.Printf("Untyped constant %v , %T \n", a, a)
	fmt.Printf("Untyped constant %v , %T \n", b, b)
}
