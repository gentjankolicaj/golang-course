//go:build ignored

package main

import (
	"fmt"
	"golang_course/section_29/speech"
	"golang_course/section_29/word"
)

func main(){

	alCount:=word.Count(speech.ALincolnSpeech)
	tjCount:=word.Count(speech.TJeffersonSpeech)
	fmt.Println("A Lincoln speech count : ",alCount)
	fmt.Println("T Jefferson speech count : ",tjCount)

	for k,v:=range word.UseCount(speech.ALincolnSpeech){
		fmt.Println(k,v)
	}

	
}