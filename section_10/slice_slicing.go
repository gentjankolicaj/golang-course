//go:build ignore

package main

import "fmt"

//Composite literal => x:=type{}
//Slice allows grouping values of same type.
//Slicing to a slice x[a:b] where a=startIndex included & b=endIndex not included

func main() {
	slice := []int{93, 2, 32, 4, 3, 1, 6, 7, 8, 434, 2, 3, 43}

	//Slicing a slice x[startIndex:endIndex] where endIndex not included
	fmt.Println(slice)
	fmt.Println(slice[:])
	fmt.Println(slice[0:1])           //Slicing a slice from index 0 to 1, not included
	fmt.Println(slice[5:])            //Slicing a slice from index 5 to end, not included
	fmt.Println(slice[len(slice)/2:]) //Slice in half

	slice = slice[:len(slice)/2]
	fmt.Println("New half of slice ", slice)
}
