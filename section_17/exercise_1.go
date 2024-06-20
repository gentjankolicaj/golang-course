//go:build ignore

package main

import "fmt"

//Pointer : value pointing to some memory address
//Pointer : We get pointer value by using operator &,it gives memory address
//Pointer is different from type that is pointing to
//Operator & : Gives address where value is stored
//Operator * : Gives the value that address is holding (also called de-referencing)
//Everything in GOLANG is pass by value

//Pointer : Good to be used when we have a large group of data, and we don't want to pass it around
//Pointer : To change a value at certain location
func main() {
	var a bool = true
	var b uint = 12
	var c int = 13
	var d float64 = 3.14
	var e byte = 150
	var f rune = 23
	var g string = "hello exercise 1"

	var h = struct {
		id       string
		username string
	}{
		id:       "23242sdfasf",
		username: "johndoe",
	}

	var i interface {
		call() string
	}

	var j func() = func() {
		fmt.Println("Hello from inside func() assigned to j")
	}

	fmt.Printf("%v , %v , %T \n", a, &a, &a)
	fmt.Printf("%v , %v , %T \n", b, &b, &b)
	fmt.Printf("%v , %v , %T \n", c, &c, &c)
	fmt.Printf("%v , %v , %T \n", d, &d, &d)
	fmt.Printf("%v , %v , %T \n", e, &e, &e)
	fmt.Printf("%v , %v , %T \n", f, &f, &f)

	//Special types
	fmt.Printf("%v , %v , %T \n", g, &g, &g)
	fmt.Printf("%v , %v , %T \n", h, &h, &h)
	fmt.Printf("%v , %v , %T \n", i, &i, &i)
	fmt.Printf("%v , %v , %T \n", j, &j, &j)
}
