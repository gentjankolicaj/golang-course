//go:build ignore

package main

import "fmt"

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("Original slice ", s)

	//Append value 52
	s = append(s, 52)
	fmt.Println("Appended to slice ", s)

	s = append(s, 53, 54, 55)
	fmt.Println("Appended values to a slice ", s)

	values := []int{56, 57, 58, 59, 60}
	s = append(s, values...)
	fmt.Println("Appended slice to a slice ", s)
}
