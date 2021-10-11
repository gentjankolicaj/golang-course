//go:build ignored

package main

import (
	"errors"
	"log"
	"strconv"
)

type customError struct {
	Id  int
	err error
}

//IMPLEMENT error interface by having a method of Error()string
func (c customError) Error() string {
	return "Id " + strconv.Itoa(c.Id) + " err " + c.err.Error()
}

var packageError customError = customError{101, errors.New("NotEmpty")}

func main() {
	var localError customError = customError{202, errors.New("NotEmpty")}

	fx(packageError)
	fx(localError)

}

func fx(e error) {
	log.Println(e)
}
