//go:build ignore

package main

import "fmt"

type Cube struct {
	length, width, height int
}

func main() {
	var cubeMap map[int]Cube = make(map[int]Cube)
	fmt.Println(cubeMap, len(cubeMap))

	for i := 0; i < 7; i++ {
		var cubeTmp Cube = Cube{
			height: i + 2,
			width:  i + 34,
			length: i + 54,
		}
		cubeMap[i] = cubeTmp
	}

	for key, value := range cubeMap {
		fmt.Printf("\n %d , %v ", key, value)
	}
}
