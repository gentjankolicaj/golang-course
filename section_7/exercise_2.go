//go:build ignore

package main

import "fmt"

//Operators of comparison :
//===, >= , <= , > , <
var (
	a = 12
	b = (23 <= 12)
	c = (23 >= 12)
	d = (23 > 12)
	e = (23 < 12)
)

const (
	g = 12
	h = (23 <= 12)
	i = (23 >= 12)
	j = (23 > 12)
	k = (23 < 12)
)

func main() {
	fmt.Println(a, b, c, d, e)
	fmt.Println(g, h, i, j, k)
}
