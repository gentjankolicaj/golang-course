//go:build ignore

package main

import "fmt"

//Func template : func (r receiver) identifier(param1 type,....) (type1,...){}
//Variadic param : Is an operator in golang for variable number of params at functions
//Operator ... helps create variadic params
func main() {
	variadicParam(0, 1, 2, 3, 4, 5, 6, 7)

	//Calculate sum using variadic params
	sum := v2(1, 2, 3, 4, 4, 5, 56, 5, 67676, 565, 343)
	fmt.Println("Total sum : ", sum)
}

func variadicParam(x ...int) { //type of param here is a slice of int
	fmt.Println(x)
	fmt.Printf("%T \n", x)
}

func v2(x ...int) int {
	sum := 0;
	for i, v := range x {
		fmt.Printf("%d , %d \n", i, v)
		sum += v
	}
	return sum;
}
