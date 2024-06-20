//go:build ignored

package main

import (
	"fmt"
	"runtime"
	"sync"
)

//NOTE: Mutex used for read-write synchronization on variables or sequence of instructions
func main() {
	var maxGoroutine = 20
	var counter = 0
	var waitGroup sync.WaitGroup
	var mx sync.Mutex //Variable for read-write synchronization

	//Add goroutine number to wait-group
	waitGroup.Add(maxGoroutine)

	for i := 0; i < maxGoroutine; i++ {
		go func() {
			mx.Lock() //Lock for upcoming go-routines to read-write
			v := counter
			v = v + 5
			runtime.Gosched() //Yield processor when goroutine is finished exe
			counter = v
			mx.Unlock()      //Unlock for upcoming go-routines to read-write
			waitGroup.Done() //Decrement wait-group counter by 1
		}()
	}

	waitGroup.Wait() //Make main-routine wait for other go-routines to finish
	fmt.Println("Counter ", counter)

}
