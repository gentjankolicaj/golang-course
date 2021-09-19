//go:build ignore

package main

import "fmt"

var a bool

var b uint

var c string

var d int

var e float64

var f interface{}

var g byte  //uint8 2^8

func main() {

	fmt.Println(a)
	fmt.Printf("%T \n",a)

	fmt.Println(b)
	fmt.Printf("%T \n",b)

	fmt.Println(c)
	fmt.Printf("%T \n",c)

	fmt.Println(d)
	fmt.Printf("%T \n",d)

	fmt.Println(e)
	fmt.Printf("%T \n",e)

	fmt.Println(f)
	fmt.Printf("%T \n",f)
	
	fmt.Println(g)
	fmt.Printf("%T \n",g)
}
