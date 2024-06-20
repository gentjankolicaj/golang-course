//go:build ignored

package main

import (
	"fmt"
	"runtime"
	"sync"
)

//NOTE: to check for race-condition run => go run -race exercise_3.go
//WaitGroup : Used for making main-routine to wait or other go-routines
func main() {

	var wg sync.WaitGroup

	var counter int = 0
	maxGoroutines := 10

	wg.Add(maxGoroutines) //Add max goroutine number to wait group in order to wait
	for i := 0; i < maxGoroutines; i++ {
		go func() { //ANON function declared & started go-routine

			routineValue := counter //Read counter value & stored into local variable
			runtime.Gosched()       //Yield processor to
			routineValue = routineValue + 2
			counter = routineValue //Write to counter new value

			wg.Done() //Decrement wait-group counter with 1
		}()
	}

	wg.Wait() //Wait for all goroutines to finish executing
	fmt.Println("Go-routines : ", runtime.NumGoroutine())
	fmt.Println("Count value : ", counter)
}
