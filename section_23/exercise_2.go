//go:build ignored

package main

import "fmt"

func main() {
	//Send channel
	cs := make(chan<- int)

	go func() {
		cs <- 22
		cs <- 232
	}()

	fmt.Printf("%T  , %v \n", cs, cs)

}
