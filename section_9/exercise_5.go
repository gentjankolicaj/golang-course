//go:build ignore

package main

import "fmt"

func main() {

	for a := 10; a <= 100; a++ {
		var remainder int = a % 4
		fmt.Println(remainder)
	}

}
