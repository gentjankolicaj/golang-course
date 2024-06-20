//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
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

	var slice2D [][]byte = make([][]byte, 7, 10)
	for i, v := range persons {
		bs, _ := json.Marshal(v)
		slice2D[i] = bs
		fmt.Println("2d slice marshalled ", string(bs))
	}

	bs, _ := json.Marshal(persons)
	fmt.Println("\n", "Persons slice marshalled ", string(bs))
}
