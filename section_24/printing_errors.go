//go:build ignored

package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Err happened ", err)
	}
}
