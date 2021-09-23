//go:build ignore

package main

import "fmt"

//Conditional control flow templates
//1. if condition {}
//2. if {} else {}
//3. else if condition{}
//4. if expression; condition{}
func main() {
	IfSample()
	ElseIfSample()

	//4.if expression; condition{}
	if x := 22; x == 22 {
		//Scope of x limited inside if
		fmt.Println(x)
	}
}

func IfSample() {
	var index int = 10
	if index > 11 {
		fmt.Println("Condition 2 ", index > 11)
	}

	if index == 11 {
		fmt.Println("Condition 2 ", index == 11)
	}

	if index == 12 {
		fmt.Println("Condition 3 ", index == 12)
	}

	if index == 14 {
		fmt.Println("Condition 4 ", index == 14)
	}
}

func ElseIfSample() {
	var value int = 0
	if value == 1 {
		//do something
	} else if value == 2 {
		//do something
	} else if value == 3 {
		//do something
	} else if value == 4 {
		//do something
	} else {
		fmt.Println("Last condition : ", value)
	}

}
