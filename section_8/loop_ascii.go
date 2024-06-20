//go:build ignore

package main

import "fmt"

func main() {
	LoopAscii()
}

func LoopAscii() {
	var i int = 0
	for ; i < 200; i++ {
		fmt.Printf("\n Hex %x , Dec %d , Ascii char %c , code-point %U", i, i, i, i)
	}

}
