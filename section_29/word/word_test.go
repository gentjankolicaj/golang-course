package word

import (
	"testing"
	"fmt"
)


func TestUseCount(t *testing.T){
	s:="a , b ,c d e f g a g b b b b f c c c f f a a aaa a aaa aa  ba babb"
	m:=UseCount(s)
	aCounter:=m["a"]
	if aCounter!=5{
		t.Error("Expected ",5,"Actual ",aCounter)
	}

}

func TestCount(t *testing.T){
	s:="a , b , c b b b a a a aa , a a a "
	counter:=Count(s)
	if counter!=16{
		t.Error("Expected ",16,"Actual ",counter)
	}

}

func BenchmarkUseCount(b *testing.B){
	s:="a , b ,c d e f g a g b b b b f c c c f f a a aaa a aaa aa  ba babb"
	for i:=0;i<b.N;i++{
		UseCount(s)
	}
}

func BenchmarkCount(b *testing.B){
	s:="a , b ,c d e f g a g b b b b f c c c f f a a aaa a aaa aa  ba babb"
	for i:=0;i<b.N;i++{
		Count(s)
	}
}


func ExampleUseCount(){
	s:="a , b ,c d e f g a g b b b b f c c c f f a a aaa a aaa aa  ba babb"
	fmt.Println(UseCount(s))
	//Output : 
	//
}

func ExampleCount(){
	s:="a , b ,c d e f g a g b b b b f c c c f f a a aaa a aaa aa  ba babb"
	fmt.Println(Count(s))
	//Output : 
	//
}
