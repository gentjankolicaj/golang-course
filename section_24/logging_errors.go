//go:build ignored

package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("file.txt")
	if err != nil {
		log.Println("Err happened ", err)
	}

	logFile()
}

func logFile() {
	file1, err := os.Create("log.txt")
	if err != nil {
		log.Println(err)
	}
	defer file1.Close()
	log.SetOutput(file1) //Since file implements a method of writer interface,is of type writer

	file2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("Error happened ", err)
	}
	defer file2.Close()
	log.Println("Check log.txt file")
	log.Fatalln("Fatal message : ") //Invokes os.Exit()
	//Defered functions are not run because of os.Exit()
}
