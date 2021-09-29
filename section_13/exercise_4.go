//go:build ignore

package main

import (
	"fmt"
)

func main() {

	//Create & use anonymous struct
	superHero := struct {
		first, last   string
		age, height   int
		id, code_name string
		country, city string
		friends       []string
	}{
		first:     "Bruce",
		last:      "Wayne",
		age:       33,
		height:    185,
		id:        "Batman",
		code_name: "Dark knight",
		country:   "USA",
		city:      "Gotham",
		friends:   []string{"Superman,Flash"},
	}

	fmt.Println("Super-hero : ", superHero)
}
