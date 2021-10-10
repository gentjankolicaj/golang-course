//go:build ignored

package main

import (
	"context"
	"fmt"
)

//Context : A package to use with concurrent design patterns
func main() {
	backgroundPrint()

}

func backgroundPrint() {
	ctx := context.Background()
	fmt.Println("Context ", ctx)
	fmt.Println("Context err ", ctx.Err())
	fmt.Println("Context type ", ctx.Done())
	fmt.Println("-------------------")
}
