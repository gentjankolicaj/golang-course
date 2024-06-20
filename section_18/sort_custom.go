//go:build ignore

package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

//ByAge implements sort.Interface, interface for []Person based on age field
type ByAge []person

func (b ByAge) Len() int           { return len(b) }
func (b ByAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByAge) Less(i, j int) bool { return b[i].age < b[j].age }

func main() {
	a := person{
		first: "Andrea",
		age:   22,
	}

	b := person{
		first: "Baci",
		age:   22,
	}

	d := person{
		first: "Dane",
		age:   22,
	}

	c := person{
		first: "Cesar",
		age:   32,
	}

	z := person{
		first: "Zed",
		age:   34,
	}

	persons := []person{a, z, c, z, b, d}

	fmt.Println("Unsorted ", persons)

	//Sort slice of struct
	sort.Sort(ByAge(persons))
	fmt.Println("Sorted ", persons)
}
