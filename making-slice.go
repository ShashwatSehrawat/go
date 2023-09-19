package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	var b = make([]int, 0, 5)

	printSlice(a)
	printSlice(b)

	c := a[2:]
	d := b[:3]

	printSlice(c)
	printSlice(d)
}

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d \t %v\n", len(s), cap(s), s)
}
