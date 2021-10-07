//go:build ignored

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//NOTE : Atomic package used for atomic read-write operation on variables
func main() {
	var maxGoroutine = 20
	var counter int64 = 0
	var waitGroup sync.WaitGroup

	//Add goroutine number to wait-group
	waitGroup.Add(maxGoroutine)

	for i := 0; i < maxGoroutine; i++ {
		go func() {
			idx := i
			atomic.AddInt64(&counter, 5)                               //Atomic write to variable counter
			fmt.Println("Goroutine-", idx, atomic.LoadInt64(&counter)) //Atomic read at variable counter

			runtime.Gosched() //Yield processor when goroutine is finished exe
			waitGroup.Done()  //Decrement wait-group counter by 1
		}()
	}

	waitGroup.Wait() //Make main-routine wait for other go-routines to finish
	fmt.Println("Counter ", counter)

}
