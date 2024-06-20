//go:build ignored

package main

import "fmt"

//DEFER FUNC: functions arguments are evaluated when defer statement is evaluated
//DEFER FUNC: Execution order is LIFO,last in first out
func main() {

	for i := 0; i < 10; i++ {
		defer fmt.Println("index ", i)
	}
}
