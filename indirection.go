package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	a := Vertex{2, 3}
	a.scale(2)
	scaleFunc(&a, 10)
	b := &Vertex{1, 2}
	b.scale(2)
	scaleFunc(b, 10)
	fmt.Println(a, b)
}
