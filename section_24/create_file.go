//go:build ignored

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//IMPORTANT: Always check errors
func main() {
	f, err := os.Create("new_file.txt")
	if err != nil {
		fmt.Println(err)
		return //returns main function
	}
	defer f.Close() //Invoke Close() method after main is returned

	for i := 0; i < 10; i++ {
		r := strings.NewReader("Hi,how are you doing ?\n")
		io.Copy(f, r)
	}

}
