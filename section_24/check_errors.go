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
	f, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
		return //returns main function
	}
	defer f.Close() //Invoke Close() method after main is returned

	r := strings.NewReader("Hi,how are you doing ?")
	io.Copy(f, r)

}
