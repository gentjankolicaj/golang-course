//go:build ignore

package main

import "fmt"

type person struct {
	first, last     string
	age             int
	email, password string
	address         address
}

type address struct {
	street  string
	city    string
	country string
}

//Methods : normal functions with receiver => func (r receiver) identifier(params) (returns){}
//func (r receiver) identifier(params) (returns) {}
//When we have a receiver to a function,function is attached to the type => This transforms function into method
//Function becomes method because it is related to state of type

//This method belongs to struct person
func (p person) toString() string {
	fmt.Println("I am ", p.first, "", p.last)
	return p.first + " " + p.last
}

func main() {
	johnDoe := person{
		first:    "John",
		last:     "Doe",
		age:      32,
		email:    "Noneyourconcern@doe.com",
		password: "lalalalal",
		address: address{
			street:  "234",
			city:    "Tirana",
			country: "Albania",
		},
	}
	fmt.Println(johnDoe)

	//Call method of struct person
	johnDoe.toString()

}
