package main

import (
	"fmt"
)

type T struct {
	S string
}

type I interface {
	M()
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println(t.S)
	}
}
func main() {
	var i I
	var t *T
	i = t
	describe(i)
	i.M()

	t = &T{"hello"}
	i = t
	describe(t)
	i.M()

}
func describe(i I) {
	fmt.Printf("%T : %v\n", i, i)
}
