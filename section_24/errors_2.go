//go:build ignored

package main

import (
	"fmt"
	"log"
	"os"
)

//NOTE:Errors are just another type in golang
//NOTE:Any type that implements all methods of an interface is a type of that interface.

type person struct {
	First string
	Last  string
}

func (p *person) identify() string {
	return p.First + " " + p.Last
}

func main() {
	jason := person{
		First: "Jason",
		Last:  "Todd",
	}

	logFile, err := os.Create("persons.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile) //Set log output destination on a log-file
	log.Println(jason.identify())

	json, err := json2.Marshal(jason)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(string(json))
}
