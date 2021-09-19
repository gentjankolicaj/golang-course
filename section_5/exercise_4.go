//go:build ignore

package main

import "fmt"

//Custom type
type life_form int

var human life_form
var martian life_form
var kriptonian life_form

func PerformOperations() {
	s := fmt.Sprintf("%v\t%v\t%v", human, martian, kriptonian)
	fmt.Println(s)

	//Value assignment
	human = 1
	martian = 2
	kriptonian = 999

	s = fmt.Sprintf("%v\t%v\t%v", human, martian, kriptonian)
	fmt.Println(s)

}

func main() {

	PerformOperations()
}
