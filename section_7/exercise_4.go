//go:build ignore

package main

import "fmt"

var packageVar int = 8

func main() {
	fmt.Printf("Initial %d , %b , %x \n", packageVar, packageVar, packageVar)

	anotherVar := packageVar << 1
	fmt.Printf("Shifted-left %d , %b , %x ", anotherVar, anotherVar, anotherVar)

}
