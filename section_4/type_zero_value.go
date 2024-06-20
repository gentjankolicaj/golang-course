//go:build ignore

package main

import "fmt"

var a bool

var b uint

var c string

var d int

var e float64

var f interface{}

var g byte //uint8 2^8

//When we declare a variable and not initialize default values are assigned
//Boolean => false
//Integer => 0
//Double/float => 0.0
//String =>""
//interface{} => nil
//Pointer => nil
//functions => nil
//slices =>nil
//map => nil
//channels => nil
func main() {

	fmt.Println(a)
	fmt.Printf(" %T", a)

	fmt.Println(b)
	fmt.Printf(" %T", b)

	fmt.Println(c)
	fmt.Printf(" %T", c)

	fmt.Println(d)
	fmt.Printf(" %T", d)

	fmt.Println(e)
	fmt.Printf(" %T", e)

	fmt.Println(f)
	fmt.Printf(" %T", f)

	fmt.Println(g)
	fmt.Printf(" %T", g)
}
