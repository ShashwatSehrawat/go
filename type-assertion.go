package main

import (
	"fmt"
)

func main() {
	var i interface{} = "Hello"
	a := i.(string)
	fmt.Println(a)
	b, k := i.(string)
	fmt.Println(b, k)
	c, j := i.(float64)
	fmt.Println(c, j)
	m, l := i.(float64)
	fmt.Println(m, l)
}
