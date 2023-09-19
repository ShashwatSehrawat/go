package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Index(a, 5))

	b := []string{"hello", "how", "are", "you"}

	fmt.Println(Index(b, "how"))
}
