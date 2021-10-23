//package contains functions for working with words
package word

import "strings"


func UseCount(s string) map[string]int{
	stringSlice:=strings.Fields(s)
	var counterMap map[string]int=make(map[string]int)

	for _,value:=range(stringSlice){
		counterMap[value]=counterMap[value]+1 //or counterMap[value]++
	}
	return counterMap
}

func Count(s string)int{
	return len(strings.Fields(s))
}