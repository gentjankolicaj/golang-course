//go:build ignore

package main

import "fmt"

//Composite literal => x:=type{}
//Slice allows grouping values of same type.
//A slice is built on top of an array with difference that slices are dynamic

//Creating-slice => builtin make([]T,length,capacity)
//Length=>Initial length with allocated values
//Capacity=>Number of spots in underlying array to use.

//Slicing-slice => x[a:b] where a=startIndex included & b=endIndex not included
//Appending-slice => builtin package has function append(Slice,...types)[]T => append(slice,a,b,c,c,c)
//Delete-slice => using slicing[] & append() function

func main() {
	slice := make([]int, 10, 100)
	fmt.Println("Length & capacity ", len(slice), cap(slice))

	//Slicing a slice x[startIndex:endIndex] where endIndex not included
	fmt.Println("Initial slice", slice)
	fmt.Println("Initial slice", slice[:])

	//Appending to slice
	slice2 := append(slice, 0, 0, 1, 1, 1, 1, 1, 1)
	fmt.Println("Appended slice2 ", slice2)

	slice3 := append(slice2, slice2...)
	fmt.Println("Appended slice3 ", slice3)

	//In slice in generally I delete using slicing & append
	slice3 = append(slice3[:4], slice3[7:]...)
	fmt.Println("Deleted in slice ", slice3)

}
