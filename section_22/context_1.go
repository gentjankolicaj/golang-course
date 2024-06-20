//go:build ignored

package main

import (
	"context"
	"fmt"
)

//Context : A package to use with concurrent design patterns
func main() {
	ctx := context.Background()
	fmt.Println("Context ", ctx)
	fmt.Println("Context err ", ctx.Err())
	fmt.Println("Context type ", ctx.Done())
	fmt.Println("-------------------")

	childCtx, cancel := context.WithCancel(ctx)
	fmt.Println("Child-context ", childCtx)
	fmt.Println("Child-context err ", childCtx.Err())
	fmt.Println("Child-context type ", childCtx.Done())
	fmt.Println("-------------------")

	cancel() //cancel context
	fmt.Println("After cancel Child-context ", childCtx)
	fmt.Println("After cancel Child-context err ", childCtx.Err())
	fmt.Println("After cancel Child-context type ", childCtx.Done())
	fmt.Println("-------------------")

}
