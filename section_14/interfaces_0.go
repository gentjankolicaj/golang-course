//go:build ignore

package main

import "fmt"

//Interfaces allow us to define behaviour
//How to define interface : keyword identifier type => type human interface
//How to define custom type : keyword identifier type

//A value can be of more than one type.

//Custom struct & methods
type person struct {
	first string
	last  string
}

func (p person) speak() {
	fmt.Println("Hello from ", p.first, "-", p.last)
}

//Interface & methods
type human interface {
	speak()
}

//A VALUE CAN BE OF MORE THAN ONE TYPE
func main() {

	john := person{
		first: "JOHN",
		last:  "DOE",
	}

	jane := person{
		first: "JANE",
		last:  "DOE",
	}

	//JOHN & JANE is of type person & human.
	//Human because interface human has method speak()
	fmt.Println(jane)
	fmt.Println(john)

	jane.speak()
	john.speak()

}
