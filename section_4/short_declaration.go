//go:build ignore

package main

import "fmt"

//Short declaration operator => := ,and used only locally not package scope
func main() {
	BooleanType()
	SignedIntegerType()
	SignedDoubleType()
	StringType()
}

func BooleanType() {
	//Short declaration of boolean
	flag := true
	fmt.Println(flag)
	flag = false
	fmt.Println(flag)
}

func SignedIntegerType() {
	//Short declaration
	value := 10
	fmt.Println(value)
}

func SignedDoubleType() {
	//Short declaration
	value := -2.1
	fmt.Println(value)
}

func StringType() {
	//Short declaration
	str := "StringType"
	fmt.Println(str)
}
