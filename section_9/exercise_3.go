//go:build ignore

package main

import "fmt"

func main() {
	var birthYear = 1994
	var actualYear = 2021

	//Years that I have been alive
	for birthYear <= actualYear {
		fmt.Println(birthYear)
		birthYear++
	}

}
