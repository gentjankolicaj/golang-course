//go:build ignore

package main

import "fmt"

var x bool //Zero-Value: false =>Assigned by compiler on compile-time

func main() {

	fmt.Println("X value : ", x)
	x = true
	fmt.Println("X update-value : ", x)

	//Short-declaration operator
	a := 7
	b := 8

	//equals operator
	var flag bool = a == b
	fmt.Println("a==b", flag)

	//Golang comparison operators :
	// ==
	// >=
	// <=

}
