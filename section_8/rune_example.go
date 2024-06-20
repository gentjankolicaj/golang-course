//go:build ignore

package main

import "fmt"

const (
	A = rune(22)
	B = rune(23)
)

func main() {
	var a rune = 12
	var b rune = rune(223)
	fmt.Println("Package constants ", A, B)
	fmt.Println("Local vars ", a, b)

	var index int = 100
	var runeVal = rune(index)
	fmt.Println("int to rune : ", index, runeVal)
}
