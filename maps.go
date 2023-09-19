package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}
type Car struct {
	Name  string
	Price int
}

var m map[string]Vertex
var c map[string]Car

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		-14.10, 10.23,
	}
	fmt.Println(m["Bell Labs"])

	c = make(map[string]Car)
	c["Suzuki"] = Car{
		"Suzuki", 10,
	}
	fmt.Println(c["Suzuki"])
}
