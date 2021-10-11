//go:build ignored

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("new_file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Read content : ")
	fmt.Println(string(bs))
}
