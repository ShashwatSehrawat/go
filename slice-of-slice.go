package main

import (
	"fmt"
	"strings"
)

func main() {
	b := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	//player turns

	b[0][0] = "x"
	b[2][2] = "o"
	b[1][2] = "x"
	b[1][0] = "o"
	b[0][2] = "x"

	printSlice(b)
}

func printSlice(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}
