
package main

import (
	"fmt"
	"strings"
)

const s="Golang Variable: A variable is just a name given to a particular memory space in the storage, by which the Go program manipulates that variable. The memory space in the go program is decided by the type of the variable, and each type has its pre-decided size in the go language. The type of variable decides the storage and the range of the value in the go program.A variable name consists of alphabets, digits and underscore, which represent data stored in the storage of your computer. At the time of the declaration, the compiler creates a memory block for the variable in the storage. The compiler uses the variable name for accessing that data during the execution of the go program."

var xs []string

func main(){
 xs=strings.Split(s," ")

 for _,v:=range(xs){
	 fmt.Println(v)
 }

 fmt.Printf("Printf: \n %s \n",Cat(xs))
 fmt.Printf("Printf: \n %s \n",Join(xs))
}


func Cat(xs []string) string{
	s:=""
	for _,v:=range(xs){
		s=s+v+" "
	}

	return s
}

func Join(xs []string)string{
	return strings.Join(xs," ")
}