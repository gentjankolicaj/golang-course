//go:build ignore

package main

import (
	"fmt"
)

//Strings are immutable on golang & java
//Each code-point is known as rune
//Rune  alias for int32

//Package variable
var s1 string = "Hello string 1"
var (
	s2 string = "Hello string 2"
)

func main() {
	var s3 string = "Hello string 3"
	fmt.Println(s1, s2, s3)

	bs := []byte(s1)
	fmt.Println("Hello string 1, to int64 ", bs, getStringLength(s1))

	for i, v := range s1 {
		fmt.Println("")
		fmt.Printf(" At index %d |  HEX %#x | UTF-8 code-point %#U", i, v, v)
	}

	h := "H"
	//create a byte slice to get integer value of h valu
	byteSlice := []byte(h)
	fmt.Println("\n", byteSlice[0])
	fmt.Printf("\n %d , %b , %#x", byteSlice[0], byteSlice[0], byteSlice[0])

}

func getStringLength(str string) int {
	return len(str)
}
