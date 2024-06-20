//go:build ignored

package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 23
	}()

	v, ok := <-c
	fmt.Println(v, ok)

}
