//go:build ignore

package main

import (
	"fmt"
)

//Creating slice of slice
func main() {
	slice0 := []int{1, 2, 3}
	slice1 := []int{23, 4, 45, 34}
	fmt.Println("Slice0", slice0)
	fmt.Println("Slice1", slice1)

	sliceOfSlice := [][]string{
		{"User", "Tester", "Admin", "PR"},
		{"Ben", "Tim", "John", "Katia"},
	}

	fmt.Println(sliceOfSlice)
	fmt.Println(sliceOfSlice[0])
	fmt.Println(sliceOfSlice[1])

	for index, value := range sliceOfSlice {
		fmt.Printf("Index %d ,value %d \n", index, value)
		for idx, vl := range value {
			fmt.Println("Looping with value identifier : ", idx, vl)
		}
	}
}
