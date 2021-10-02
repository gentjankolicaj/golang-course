//go:build ignore

package main

import (
	"fmt"
	"sort"
)

//Sort a slice of ints & strings
func main() {
	intSlice := []int{94, 3, 4, 4, 23, 2, 4, 35, 4, 3, 43}
	stringSlice := []string{"Dave", "Bond", "Wayne", "Bruce", "Golang", "Java"}
	fmt.Println("Unsorted : ", intSlice)
	fmt.Println("Unsorted : ", stringSlice)

	sort.Ints(intSlice)
	sort.Strings(stringSlice)

	fmt.Println("Sorted : ", intSlice)
	fmt.Println("Sorted : ", stringSlice)

}
