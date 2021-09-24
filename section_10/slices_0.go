//go:build ignore

package main

import "fmt"

//Composite literal => x:=type{}
//Slice allows grouping values of same type.
func main() {
	IntSlice()
	FloatSlice()
	StringSlice()
	BoolSlice()
}

func IntSlice() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
}

func FloatSlice() {
	slice := []float64{2.2, 3, 333, 3.14}
	fmt.Println(slice)
}

func StringSlice() {
	slice := []string{"Hello", "Hello2", "Hello3", "Hello4"}
	fmt.Println(slice)
}

func BoolSlice() {
	slice := []bool{true, false, true, true, true}
	fmt.Println(slice)
}
