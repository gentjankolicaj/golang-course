//go:build ignore

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	//Normal struct ,visible on package
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   25,
	}
	fmt.Println("Normal struct", p1)

	//Anonymous struct declaration => variable:=struct{ name1 type1,name2 type2}
	//composite literal
	student := struct {
		first, last string
		age, id     int
		email       string
	}{
		first: "Tim",
		last:  "Doe",
		age:   18,
		id:    9,
		email: "notyourconcern@none.com",
	}

	fmt.Println("Anonymous struct", student)
}
