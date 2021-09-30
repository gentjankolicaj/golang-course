//go:build ignore

package main

//Pointer : value pointing to some memory address
//Pointer : We get pointer value by using operator &,it gives memory address
//Pointer is different from type that is pointing to
//Operator & : Gives address where value is stored
//Operator * : Gives the value that address is holding (also called de-referencing)
//Everything in GOLANG is pass by value

//Pointer : Good to be used when we have a large group of data, and we don't want to pass it around
//Pointer : To change a value at certain location

//METHOD SETS  : methods attached to a type
type square struct {
	length float64
}

//NON-POINTER RECEIVER : works with pointer & non-pointer (t T) => *T & T
func (s square) nonPointer() float64 {
	return s.length
}

//POINTER RECEIVER : works with pointer  (t *T) => *T
func (s *square) pointer() float64 {
	return (*s).length
}

func main() {

	s := square{
		length: 28,
	}

	s.nonPointer() //works with pointer & non-pointer (t T) => *T & T
	s.pointer()    //works with pointer & non-pointer (t T) => *T & T

	sPtr := &s //sPtr is of type *square
	sPtr.pointer()
	sPtr.nonPointer() //compile error

}
