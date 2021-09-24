//go:build ignore

package main

import "fmt"

//Creating maps with make
func main() {
	myMap := make(map[string][]string)
	myMap["doe_john"] = []string{"Likes to train", "Likes to run", "Likes to joke"}
	myMap["doi_john"] = []string{"Likes to not train", "Likes to not run", "Likes to not joke"}
	myMap["dou_john"] = []string{"Likes to write", "Likes to read", "Likes to sleep"}

	for mi, mv := range myMap {
		fmt.Println("Map-key ", mi)
		for si, sv := range mv {
			fmt.Println("Slice idx,vl ", si, sv)
		}
		fmt.Println("\n")
	}
}
