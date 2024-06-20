//go:build ignored

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

//Context : A package to use with concurrent design patterns
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Error check-1 ", ctx.Err())
	fmt.Println("Nr goroutine-1 ", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working...")
			}
		}

	}()

	time.Sleep(time.Millisecond * 20)
	fmt.Println("Error check-2 ", ctx.Err())
	fmt.Println("Nr goroutine-2 ", runtime.NumGoroutine())
	fmt.Println("About to cancel context")

	cancel() //Invoke function
	fmt.Println("Error check-3 ", ctx.Err())
	fmt.Println("Nr goroutine-3 ", runtime.NumGoroutine())
}
