//go:build ignore

package main

import "fmt"

//Composite literal => x:=type{}
//Slice allows grouping values of same type.
func main() {
	slice := []int{1, 2, 3, 4, 5, 2, 3443, 545, 667, 876, 2343}
	fmt.Println(slice)
	fmt.Println("Looping on Int slice using range : ")
	for index, value := range slice {
		fmt.Printf("Index %d , Value %d \n", index, value)
	}
	fmt.Println("\n")
}
