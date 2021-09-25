//go:build ignore

package main

import "fmt"

//Struct are known as aggregate/composite/complex data structure
//Structs allow us to compose together multiple/different types to create new type
//If struct fields are not initialized by developer ,then they receive zero-value during compilation
type person struct {
	first string
	last  string
	age   uint
}

type agent struct {
	first, last string
	age         int
	id          int
}

func main() {
	//Composite literal struct
	p1 := person{first: "JOhn",
		last: "Doe",
		age:  12,
	}

	p2 := person{
		first: "Jane",
		last:  "Doe",
		age:   12,
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.age, p2.first)

	agent47 := agent{
		first: "47",
		last:  "47",
		age:   30,
		id:    47,
	}
	fmt.Println(agent47)

}
