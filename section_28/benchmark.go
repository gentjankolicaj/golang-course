//go:build ignored
package main

import "fmt"

//NOTE:Allows to measure performance of code
//HOW TO RUN BENCHMARK : go test -bench .
func main(){

	fmt.Println(Greet("King"))
    fmt.Println(Greet("Tony"))
}

func Greet(s string)string{
	return fmt.Sprint("Welcome ",s)
}