//go:build ignored

package main

import (
	"fmt"
	"golang_course/section_29/dog"
)



type canine struct{
	name string
	age int
}

func main(){
   rex:=canine{
	   name:"Rex",
	   age:8,
   }

   lilo:=canine{
	   name:"lilo",
	   age:5,
   }

   fmt.Println("Strc ",rex,dog.DogToHumanYears(rex.age))
   fmt.Println("Strc : ",lilo,dog.DogToHumanYears(lilo.age))

   

	
}

