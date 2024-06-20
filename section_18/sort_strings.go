//go:build ignore

package main

import (
	"fmt"
	"sort"
)

//Writer interface
func main() {
	s := []string{"Alea", "James", "Ben", "Tim", "Dave"}
	fmt.Println("Unsorted ", s)

	sort.Strings(s)
	fmt.Println("Sorted ", s)
}
