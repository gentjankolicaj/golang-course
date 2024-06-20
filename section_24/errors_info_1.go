//go:build ignored

package main

import (
	"errors"
	"fmt"
	"log"
)

//NOTE:EACH variable & function that starts with UPPER-CASE is exported outside package
//Exported outside package means is visible outside package

var mathError error = errors.New("Math error : square root of negative number")

func main() {
	fmt.Printf("%T \n ", mathError)

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, mathError
	} else {
		return 10, nil
	}
}
