//go:build ignore

package main

import "fmt"

//Bit shifting in golang.Below are bit shifting operators
//Shift left <<
//Shift right >>

func main() {
	x := 2
	fmt.Printf("Initial value of x %d , %b \n", x, x)

	//Bitshift left
	//Shifting bits 2 places to the left,and reassign value to x
	x = x << 2
	fmt.Printf("Left-shift value of x %d , %b \n", x, x)

	//Bitshift right
	x = x >> 2
	fmt.Printf("Right-shift value of x %d , %b \n", x, x)

	// 1<<10 shifts 1 to 10 positions left => 1*2^10

	ByteUnitsExample()
}

func ByteUnitsExample() {
	kb := 1024
	mb := kb * 1024
	gb := mb * 1024

	fmt.Printf("kb %d | %b", kb, kb)
	fmt.Printf("mb %d | %b", mb, mb)
	fmt.Printf("gb %d | %b", gb, gb)

}
