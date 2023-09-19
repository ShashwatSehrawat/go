package main

import (
	"fmt"
)

func adder() func(int) int {
	// values of the function closures persist between different calls when assigned to function variables.
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()

	for i := 0; i < 5; i++ {
		fmt.Println(
			i,
			pos(i),
			neg(-2*i),
		)
	}
}
