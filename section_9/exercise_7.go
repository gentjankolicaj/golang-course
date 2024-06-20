//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var val = 101
	if val == 0 {
		fmt.Printf("0-Dec %d , UTF-8 code-point %U \n", val, val)
	} else if val == 1 {
		fmt.Printf("1-Dec %d , UTF-8 code-point %U \n", val, val)
	} else if val == 10 {
		fmt.Printf("10-Dec %d , UTF-8 code-point %U \n", val, val)
	} else if val == 101 {
		fmt.Printf("101-Dec %d , UTF-8 code-point %U \n", val, val)
	}
}
