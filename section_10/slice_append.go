//go:build ignore

package main

import "fmt"

//Composite literal => x:=type{}
//Slice allows grouping values of same type.
//Slicing-slice=> x[a:b] where a=startIndex included & b=endIndex not included
//Appending-slice=> builtin append(Slice,...types)[]T => append(slice,a,b,c,c,c)

func main() {
	slice := []int{93, 2, 32, 4, 3, 1, 6, 7, 8, 434, 2, 3, 43}

	//Slicing a slice x[startIndex:endIndex] where endIndex not included
	fmt.Println(slice)
	fmt.Println(slice[:])

	//Appending to slice
	slice2 := append(slice, 0, 0, 1, 1, 1, 1, 1, 1)
	fmt.Println("Appended slice2 ", slice2)

	slice3 := append(slice2, slice2...)
	fmt.Println("Appended slice3 ", slice3)

}
