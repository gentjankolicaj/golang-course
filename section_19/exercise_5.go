//go:build ignore

package main

import (
	"fmt"
	"sort"
)

type person struct {
	Id   int
	Name string
}

//NOTE: By creating new types we are  able to make a slice of person to be sorted

//Create a type & its method for sorting a slice of persons by id
type ById []person

func (b ById) Len() int {
	return len(b)
}
func (b ById) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ById) Less(i, j int) bool {
	return b[i].Id < b[j].Id
}

//Create a type & its methods for sorting a slice of persons by name
type ByName []person

func (b ByName) Len() int {
	return len(b)
}
func (b ByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByName) Less(i, j int) bool {
	return b[i].Name < b[j].Name
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

	sort.Sort(ById(persons)) //Pass value of persons as value of ById which has underlying a slice of persons
	fmt.Println("Sorted by id ", "\n", persons)

	sort.Sort(ByName(persons)) //Pass value of persons a value of ByName which has underlying a slice of persons
	fmt.Println("Sorted by name \n", persons)

}
