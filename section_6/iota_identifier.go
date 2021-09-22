//go:build ignore

package main

import "fmt"

//IOTA a special predeclared identifier for constants of type numbers N
//Where N includes (integer,float,complex)

const (
	a = iota + 1
	b = iota + 11.0
	c = iota + complex(1, 2)
)

func main() {
	fmt.Println("Iota : ", a, b, c)
}
