//go:build ignore

package main

import "fmt"

func main() {
	var start int = 65
	var end int = 90

	for ; start <= end; start++ {
		fmt.Println(fmt.Printf("Dec %d | char %c | unicode code-point %U", start, start, start))
	}

}
