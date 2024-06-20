//go:build ignore

package main

import (
	"fmt"
)

type person struct {
	first         string
	last          string
	fav_ice_cream []string
}

func main() {
	johnDoe := person{
		first:         "John",
		last:          "Doe",
		fav_ice_cream: []string{"Vanilla", "Chocolate", "Yogurt"},
	}

	janeDoe := person{
		first:         "Jane",
		last:          "Doe Go",
		fav_ice_cream: []string{"Banana", "Blueberries"},
	}

	//Declare & create map
	var personMap map[string]person = make(map[string]person)

	//Populate map
	personMap["Doe"] = johnDoe
	personMap["Doe Go"] = janeDoe

	for key, value := range personMap {
		fmt.Printf("Key %v value %v \n", key, value)
		for i := 0; i < len(value.fav_ice_cream); i++ {
			fmt.Printf("\t\t\tFav %d - %v \n", i, value.fav_ice_cream[i])
		}
	}
}
