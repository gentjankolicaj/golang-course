//go:build ignore

package main

import (
	"fmt"
)

//not correct
func printAsci() {
	var max byte = 255
	var i byte = 0
	for ; i < max; i++ {
		fmt.Println(" ")
		fmt.Printf("Value of %c = %d ", i, i)
	}

}

func printUtf8() {
	var max byte = 255
	var i byte = 0
	for ; i < max; i++ {
		fmt.Println(" ")
		fmt.Printf("Value of %c = %d ", i, i)
	}

}

func main() {
	printAsci()
	printUtf8()
}
