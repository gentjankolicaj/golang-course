//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	ID        int
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Address   address
}

type address struct {
	zipcode string
	street  string
	city    string
	country string
}

//Writer interface
func main() {

	john := person{
		ID:        101,
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "johndoe@none.com",
		Password:  "sdfsafsaf",
		Address: address{
			zipcode: "1033",
			street:  "St. Tirana",
			city:    "Tirana",
			country: "Albania",
		},
	}

	storeJson(&john)
}

func storeJson(p *person) {
	byteSlice, err := json.Marshal(*p)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	n, err := file.Write(byteSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Number of bytes saved ", n)
	}
}
