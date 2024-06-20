//go:build ignore

package main

import "fmt"

func main() {
	IntArray()
	FloatArray()
	StringArray()
	BoolArray()
}

func IntArray() {
	var array [10]int
	fmt.Println(array)

	array[3] = 27
	fmt.Println(array)
	fmt.Println("Array length ", len(array))
}

func FloatArray() {
	var arr [5]float64
	arr[1] = 3.14
	fmt.Println("Array length ", len(arr))
	fmt.Println("Array content ", arr)
}

func StringArray() {
	var arr [4]string
	arr[0] = "Hello World of array of string"
	fmt.Println("Array content ", arr)
}

func BoolArray() {
	var arr [5]bool
	arr[2] = true
	fmt.Println("Array content", arr)
}
