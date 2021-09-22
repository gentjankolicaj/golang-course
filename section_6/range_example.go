//go:build ignore

package main

import "fmt"

func main() {
	LoopIntArray()
	LoopFloatArray()
}

func LoopIntArray() {
	fmt.Println("LoopIntArray : ")
	var array [202]int
	for i, v := range array {
		fmt.Println("Index ", i, " value ", v)
	}
}

func LoopFloatArray() {
	fmt.Println("\n", "LoopFloatArray : ")
	var array [200]float64
	for index, value := range array {
		fmt.Println("Index ", index, " value ", value)
	}
}
