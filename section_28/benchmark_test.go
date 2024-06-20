//go:build ignored

package main

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T){
	expected:="Welcome James"
	s:=Greet("James")

	if s!=expected{
		t.Error("Actual ",s,"Expected ",expected)
	}

}

func ExampleGreet(){
	fmt.Println(Greet("James"))
	//Output:
	//Welcome James
}

func BenchmarkGreet(b *testing.B){
	for i:=0;i<b.N;i++{
		Greet("James")
	}
}

