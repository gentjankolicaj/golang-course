//go:build ignore

package main

import "fmt"

func main() {
	//arr is an array because has fixed length
	arr := [5]int{0, 0, 1, 1, 1}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 2*3
	}

	for index, value := range arr {
		fmt.Printf("Idx %d Val %d | Types %T - %T \n", index, value, index, value)
	}
}
