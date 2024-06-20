//go:build ignore

package main

import "fmt"

//Composite literal => x:=type{}
//Slice allows grouping values of same type.
//A slice is built on top of an array with difference that slices are dynamic

//Creating-slice => builtin make([]T,length,capacity)
//Length=>Initial length with allocated values
//Capacity=>Number of spots in underlying array to use.Once capacity reached,new array is created with double capacity of previous

//Slicing-slice => x[a:b] where a=startIndex included & b=endIndex not included
//Appending-slice => builtin package has function append(Slice,...types)[]T => append(slice,a,b,c,c,c)
//Delete-slice => using slicing[] & append() function

//Multi dim slices => [][][]...Type from 1 to n where n is number of dimensions

func main() {
	greetingSlice := []string{"Hello composite literal", "How are you", "Where do you live", "How is everything going"}
	nameSlice := []string{"James Bond", "Jack Bauer", "Bruce Wayne"}

	multiDimSlice := [][]string{greetingSlice, nameSlice}
	fmt.Println(multiDimSlice)

	//Multi dimension slice

	_2dSlice := make([][]int, 10)
	fmt.Println("\n2D slice using make", _2dSlice)

	_3dSlice := make([][][]int, 10)
	fmt.Println("\n3D slice using make ", _3dSlice)
}
