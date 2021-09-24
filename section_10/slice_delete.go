//go:build ignore

package main

import "fmt"

//Composite literal => x:=type{}
//Slice allows grouping values of same type.
//Slicing-slice=> x[a:b] where a=startIndex included & b=endIndex not included
//Appending-slice=> builtin package has function append(Slice,...types)[]T => append(slice,a,b,c,c,c)
//Delete-slice => using slicing[] & append() function

func main() {
	slice := []int{100, 2, 32, 40, 3, 1, 6, 7, 8, 434, 2, 3, 43}

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
