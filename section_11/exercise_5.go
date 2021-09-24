//go:build ignore

package main

import "fmt"

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("Original slice ", s)

	newSlice := append(s[:3], s[6:]...)
	fmt.Println("New slice ", newSlice)
}
