//go:build ignore

package main

import "fmt"

//Func template : func (r receiver) identifier(param1 type,....) (type1,...){}
//Variadic means 0 or n number of elements
//How to omit values returned by func : => we use underscore _
func main() {
	intSlice := []int{1, 2, 3, 5, 654, 223, 45, 5, 75, 2, 54, 76, 8} //slice literal
	fmt.Println("intSlice in main ", intSlice)

	//Checking extend of pass by value
	operationOnSlice(intSlice...)
	fmt.Println("intSlice in main ", intSlice)

	//Note even though on operationOnSlice is passed a new value of slice as argument,
	//the underlying array is the same.Thats why changes to slice argument are reflected on underlying array

}

func operationOnSlice(variadic ...int) {
	variadic[0] = 0
	variadic[1] = 0
	variadic[2] = 0
	fmt.Println("variadic in operationOnSlice : ", variadic)
}
