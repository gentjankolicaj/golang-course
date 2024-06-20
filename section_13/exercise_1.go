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
		last:          "Doe",
		fav_ice_cream: nil,
	}

	fmt.Println("John foe info ", johnDoe)
	fmt.Println("Jane Doe info ", janeDoe)
}
