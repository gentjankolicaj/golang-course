//go:build ignore

package main

import "fmt"

type Mobile struct {
	brand   string
	version string
	imei    int
}

func main() {
	slice := make([]Mobile, 7, 10)
	var startCap int = cap(slice)
	fmt.Println("Before for loop \n", slice)

	for j := 0; j < startCap+1; j++ {
		var tmp Mobile = Mobile{
			brand:   "NEXUS",
			version: "V5",
			imei:    1322323232323454,
		}
		slice = append(slice, tmp) //appending on slice not reassign values to initial indexes
	}

	fmt.Println("After for loop \n", slice)
}
