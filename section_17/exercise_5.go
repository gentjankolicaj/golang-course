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
	var m map[int]*int = make(map[int]*int) //Pointer map,hahahaha different from map pointer *map[int]int

	for i := -10; i < 0; i++ {
		tmp := i
		m[i] = &tmp
	}

	fmt.Printf("%v , %T \n", m, m)

	for k, v := range m {
		fmt.Printf("%v , %v , %v , %T \n", k, v, *v, *v)
	}

}
