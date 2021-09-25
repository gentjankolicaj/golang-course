//go:build ignore

package main

import "fmt"

type address struct {
	street  string
	city    string
	country string
	zipcode int
}

//Struct person has embedded a struct of type address
//If fields of a struct are not initialized by developer,zero-value is assigned during compilation
type person struct {
	first   string
	last    string
	age     int
	address address
}

func main() {
	johnDoe := person{
		address: address{
			street:  "Kohlmarkt",
			city:    "Vienna",
			country: "Austria",
			zipcode: 1010,
		},
		first: "John",
		last:  "Doe",
		age:   28,
	}

	janeDoe := person{
		address: address{},
		first:   "Jane",
		last:    "Doe",
		age:     28,
	}

	fmt.Println(johnDoe)
	fmt.Println(janeDoe)
}
