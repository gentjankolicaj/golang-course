//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Some json samples

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

	byteSlice, err := json.Marshal(john)
	if err != nil {
		fmt.Println("Error with json marshalling ,", err)
	}
	os.Stdout.Write(byteSlice)

}
