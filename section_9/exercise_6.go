//go:build ignore

package main

import "fmt"

func main() {
	var val = 101
	if val == 0 {
		fmt.Printf("Dec %d , UTF-8 code-point %U \n", val, val)
	}
	fmt.Printf("Dec %d , UTF-8 code-point %U \n", val, val)
}
