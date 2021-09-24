//go:build ignore

package main

import "fmt"

//Composite literal => x:=type{}
//Slice allows grouping values of same type.
func main() {
	IntSlice()
	FloatSlice()

	IntSliceIndexAccess()

}

func IntSlice() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println("Looping on Int slice : ")
	for index, value := range slice {
		fmt.Printf("Index %d , Value %d \n", index, value)
	}
	fmt.Println("\n")
}

func FloatSlice() {
	slice := []float64{2.2, 3, 333, 3.14}
	fmt.Println(slice)
	fmt.Println("Looping on Float slice : ")
	for index, value := range slice {
		fmt.Printf("Index %d , Value %d \n", index, value)
	}
	fmt.Println("\n")
}

func IntSliceIndexAccess() {
	slice := []int{11, 21, 31, 14, 51}
	fmt.Println("Access slice values by index : ")
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
}
