//go:build ignore

package main

import "fmt"

func main() {
	var currentYear = 2021
	for {
		fmt.Println(currentYear)
		if currentYear <= 1994 {
			break
		}
		currentYear--
	}

}
