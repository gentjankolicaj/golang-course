//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Id   int
	Name string
}

func main() {
	p0 := person{
		Id:   1,
		Name: "One",
	}
	p1 := person{
		Id:   2,
		Name: "Two",
	}
	p2 := person{
		Id:   3,
		Name: "Three",
	}
	p3 := person{
		Id:   4,
		Name: "Four",
	}

	var persons []person
	persons = []person{p3, p2, p0, p1, p2, p0, p3}
	fmt.Println("Persons struct ", persons, "\n")

	//Encode to json a slice of users
	encoderPtr := json.NewEncoder(os.Stdout)
	err := encoderPtr.Encode(persons)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successful encoding")
	}
}
