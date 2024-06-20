//go:build ignored

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Any last wishes ? ", "Just getting started"},
	}

	bs, cErr := toJSON(p)
	if cErr != nil {
		log.Println("Error found ", string(bs))
		log.Fatalln(cErr)
	} else {
		log.Println(string(bs))
	}

}

func toJSON(p person) ([]byte, error) {
	var customError error = fmt.Errorf("Custom error with json at : %v ", time.Now())
	byteSlice, err := json.Marshal(p)
	if err != nil {
		return []byte{}, customError
	} else {
		return byteSlice, customError
	}

}
