//go:build ignored

package main

import "fmt"

//Channel:Place to store data.Transaction can occur when send&receive happen at same time
//Channel : Have blocking ability till send & receive are at same time
func main() {
	c := make(chan int, 2)

	go func() {
		c <- 42
		c <- 22
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)

}
