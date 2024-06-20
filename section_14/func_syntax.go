//go:build ignore

package main

import "fmt"

//function syntax: func (r receiver) identifier(param type,....) (type1,...){}
//We define our func with params if any
//We call our func with arguments if any
//Everything is go functions is PASS BY VALUE,which means that value is passed as func argument
//In Golang we can have multiple return types for functions & methods.
func main() {
	printPas("Important : Everything in go is PASS BY VALUE");
	printTriangle()

	s, b, i, f := multipleReturn(); //Call a function that returns multiple return types.
	fmt.Println(s, b, i, f)
}

func printPas(s string) (string) {
	//Note:Scope of s is inside printPas() func body
	fmt.Println(s)
	return fmt.Sprintf("Hello from printPass : ", s)
}

func printTriangle() {
	fmt.Println("Hello from printTriangle() function")
	var maxRow = 19;
	for i := 0; i < maxRow; i++ {
		fmt.Println("")
		var str string = ""
		for j := 0; j < i; j++ {
			str = str + "*"
		}
		fmt.Println(str)
	}
	fmt.Println("")
}

func multipleReturn() (string, bool, int, float64) {
	return "MultipleReturn", true, 101, 3.14;
}
