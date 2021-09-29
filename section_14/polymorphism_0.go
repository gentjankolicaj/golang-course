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

func callSpeak(h human) {
	fmt.Printf("From function callSpeack() ")
	h.speak()
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

	//This function has polymorphism because can take different types of arguments
	callSpeak(jane)
	callSpeak(john)

	//Conversion
	fmt.Println("\n \n")
	convert()

	fmt.Println("\n \n")
	assertType(john)

}

//CONVERSION OF TYPE
type COUNTER int

func convert() {
	var c COUNTER = 99
	var value int = 222
	fmt.Println(c, ",", value)

	//Convert
	value = int(c)
	fmt.Println(c, ",", value)
}

//ASSERTION OF TYPE
//Asserting types because can convert from interface{}
//because interface{} hasn't implemented methods of types
func assertType(x interface{}) {
	switch x.(type) { //asserting that x is of type NOT CONVERTING
	case int:
		fmt.Println("John is of type int")
	case float64:
		fmt.Println("John is of type float64")
	case string:
		fmt.Println("John is of type string")
	case bool:
		fmt.Println("John is of type bool")
	case human:
		x.(human).speak()
		fmt.Println(" is of type human")
	case person:
		fmt.Println(x.(person).first, " is of type person")
	}
}
