//go:build ignore

package main

import "fmt"

func main() {
	v0 := foo()
	i, s := bar()

	fmt.Println("foo ", v0)
	fmt.Println("bar ", i, s)
}

func foo() int {
	return 9966
}

func bar() (int, string) {
	return 101, "bar function"
}
