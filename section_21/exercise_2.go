//go:build ignored

package main

import "fmt"

type human interface {
	speak() string
}

//NOTE: Person is of type human because has a method with same signature as
//human interface
//When we have pointer receiver (t *T) we need to have a pointer to impl interface
//When we have value receiver (t T) we can have pointer/value to impl interface
type person struct {
	Id   string
	Name string
}

//Pointer-receiver => only pointer impls interface
func (p *person) speak() string {
	return "Hello this is " + p.Name + " with id " + p.Id
}

func main() {
	ana := person{
		Id:   "23232sfasd",
		Name: "Ana",
	}
	neo := person{
		Id:   "2wasdfsa",
		Name: "Neo",
	}

	//saySomething(ana,neo) => Compile error
	//CE because we have pointer receiver at function speak()
	// (t *T) => *T
	// (t T)=> T or *T

	saySomething(&ana, &neo)

}

func saySomething(ptrSlice ...human) {
	for _, val := range ptrSlice {
		fmt.Println(val.speak())
	}
}
