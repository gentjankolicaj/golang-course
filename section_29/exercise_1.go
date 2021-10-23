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

//NOTE:
//To run coverage : go test -cover .
//To run benchmark : go test -bench .
//To create coverage profile : go test -coverageprofile filename.out
//To show coverage profile on browser : go tool cover -html=c.out
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

