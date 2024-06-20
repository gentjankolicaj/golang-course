//go:build ignore

package main

import "fmt"

var literal string = `Hello this is a raw string "Inside-some string" literal`

func main() {
	fmt.Println(literal)
}
