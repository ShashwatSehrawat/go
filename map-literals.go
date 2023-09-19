package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		10.2, 13.5,
	},
	"Google": {
		13.0, 11.1,
	},
}

func main() {
	fmt.Println(m["Google"])
}
