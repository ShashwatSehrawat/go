package main

import (
	"fmt"
	"math"
)

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return (math.Sqrt(v.X*v.X + v.Y*v.Y))
}
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

type Absr interface {
	Abs() float64
}

func main() {
	// a := Vertex{1, 2}
	b := MyFloat(2.2)
	var ir Absr = &Vertex{1, 2}

	// ir = &a
	fmt.Printf("%v\n", ir.Abs())
	ir = b
	fmt.Println(ir)

}
