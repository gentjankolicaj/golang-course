//go:build ignore

package main

import "fmt"

//Func template : func (r receiver) identifier(param1 type,....) (type1,...){}
//Here we are going to unfurl a slice
//Variadic means 0 or n number of elements
//How to omit values returned by func : => we use underscore _
func main() {
	intSlice := []int{1, 2, 3, 5, 654, 223, 45, 5, 75, 2, 54, 76, 8} //slice literal

	//We pass the slice of int into function argument as variadic params
	// var slice:=[]type,  foo(slice...) where foo is function with variadic params
	// Here se unfurl slice intSlice...
	sum := sumElements(intSlice...)
	fmt.Println(sum)

}

func sumElements(v ...int) int {
	var sum int = 0;
	for _, value := range v { //Using underscore we omit value returned
		sum = sum + value;
	}
	return sum
}
