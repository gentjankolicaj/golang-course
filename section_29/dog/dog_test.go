package dog

import (
	"fmt"
	"testing"
)

func TestHumanToDogYears(t *testing.T){
	years:=HumanToDogYears(10)

	if years!=1{
		t.Error("Expected ",1,"Actual ",years)
	}
}

func TestDogToHumanYears(t *testing.T){
	years:=DogToHumanYears(10)
	if years!=100{
		t.Error("Expected ",100,"Actual ",years)
	}
}

func BenchmarkDogToHumanYears(b *testing.B){
	for i:=0;i<b.N;i++{
		DogToHumanYears(i)
	}
}

func BenchmarkHumanToDogYears(b *testing.B){
	for i:=0;i<b.N;i++{
		HumanToDogYears(i)
	}
}


func ExampleHumanToDogYears(){
	fmt.Println(HumanToDogYears(10))
	//Output:
	//1
}


func ExampleDogToHumanYears(){
	fmt.Println(DogToHumanYears(10))
	//Output:
	//100
}


