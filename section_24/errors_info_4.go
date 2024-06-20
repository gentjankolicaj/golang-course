//go:build ignored

package main

import (
	"fmt"
	"log"
)

//NOTE:Errors are just another type in golang
//NOTE:Any type that implements all methods of an interface is a type of that interface.

type customError struct {
	lat int
	lon int
	err error
}

//Implement method Error()string =>this makes customError of type error
func (c customError) Error() string {
	return fmt.Sprintf("Custom error : %v %v %v ", c.lat, c.lon, c.err)
}

var cErr error = customError{lat: 22, lon: 11, err: nil}

func main() {
	fmt.Printf("%T \n ", cErr)

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, cErr
	} else {
		return 10, nil
	}
}
