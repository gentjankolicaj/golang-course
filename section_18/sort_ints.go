//go:build ignore

package main

import (
	"fmt"
	"sort"
)

//Writer interface
func main() {
	s := []int{2, 3, 5, 6, 723, 3, 3, 4, 2, 54, 43}
	fmt.Println("Unsorted ", s)

	sort.Ints(s)
	fmt.Println("Sorted ", s)
}
