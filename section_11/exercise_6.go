//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

//Creating slice with make() builtin function
func main() {
	x := make([]string, 50, 50)
	fmt.Println(cap(x), len(x))
	fmt.Println(x, "\n")

	for d := 0; d < len(x); d++ {
		x[d] = strconv.Itoa(d) //convert int to string
	}

	fmt.Println(x)

}
