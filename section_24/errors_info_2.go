//go:build ignored

package main

import (
	"fmt"
	"log"
)

//NOTE:EACH variable & function that starts with UPPER-CASE is exported outside package
//Exported outside package means is visible outside package
func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf(" math: square root of negative number")
	} else {
		return 10, nil
	}
}
