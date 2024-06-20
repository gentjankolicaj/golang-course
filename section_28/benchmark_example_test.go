
package main

import(
	"testing"
	"strings"
)

func TestCat(t *testing.T){
 //todo
}

func TestJoin(t *testing.T){
  //todo
}

func ExampleCat(){
 //todo
}

func ExampleJoin(){
	//todo
}

func BenchmarkCat(b *testing.B){
	ss:=strings.Split(s," ")
     for i:=0;i<b.N;i++{
          Cat(ss)
	 }
}

func BenchmarkJoin(b *testing.B){
	ss:=strings.Split(s," ")
	for i:=0;i<b.N;i++{
		 Join(ss)
	}
}