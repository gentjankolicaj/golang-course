//go:build ignored

package main

import (
	"os"
)

func main() {
	file, err := os.Open("no-file.txt")
	if err != nil {
		panic(err) //Panic builtin function
		//Runs defered functions
	}
	defer file.Close()

}
