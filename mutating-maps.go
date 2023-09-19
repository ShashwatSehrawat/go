package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["value"] = 42
	fmt.Println("The value:", m["value"])
	m["value"] = 48
	fmt.Println("The value:", m["value"])

	delete(m, "value")
	fmt.Println("The value:", m["value"])

	v, ok := m["value"]
	fmt.Println(v, ok)

}
