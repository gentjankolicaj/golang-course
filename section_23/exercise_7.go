//go:build ignored

package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}

	for j := 0; j < 100; j++ {
		fmt.Println("index ", j, " - value ", <-c)
	}
	close(c)

}
