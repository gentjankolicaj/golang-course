//go:build ignore

package main

import "fmt"

func main() {
	slice := []float64{2, 3, 4, 2, 4, 5, 6, 7, 6, 2.3, 4.1}

	//Pas as argument=> function(callback) & slice
	result0 := transform(add, slice...)
	result1 := transform(subtract, slice...)

	fmt.Println("Result 0", result0)
	fmt.Println("Result 1", result1)
}

func transform(fx func(s ...float64) float64, s ...float64) float64 {
	result := fx(s...)
	return result
}

func add(s ...float64) float64 {
	var result float64 = 0
	for _, v := range s {
		result = result + v
	}
	return result
}

func subtract(s ...float64) float64 {
	var result float64 = 0
	for _, v := range s {
		result = result - v
	}
	return result
}
