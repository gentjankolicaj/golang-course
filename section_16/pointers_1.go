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
	x := 0

	//When x is passed as argument,is passed its value not reference or address
	foo(x)

	//Get a pointer from int (from int to *int)
	xPtr := &x
	boo(xPtr)
	fmt.Println("x is update from function boo() ", x)
}

func foo(y int) {
	fmt.Println("\n foo()", y)
	y = 23
	fmt.Println(y)
}

func boo(y *int) {
	fmt.Println("\n boo()", *y)
	//Note:by updating pointer value,we update original value at address
	*y = 101

	fmt.Println(*y)
}
