//go:build ignore

package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	//Slicing the slice
	part0 := slice[1:5]
	fmt.Println(part0)

	part1 := slice[0 : len(slice)/2]
	fmt.Println(part1)

	part2 := slice[len(slice)/2:]
	fmt.Println(part2)

	part3 := slice[5:10]
	fmt.Println(part3)
	fmt.Println(part3[0:3])

}
