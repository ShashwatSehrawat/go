package main

import (
	"fmt"
	"math"
	"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func day(today time.Weekday) string {
	switch time.Saturday {
	case today + 0:
		return "It is today."
	case today + 1:
		return "It is tomorrow"
	case today + 2:
		return "It is day after tomorrow"
	default:
		return "It is far away"
	}
}

type Vertex struct {
	X, Y int
}

func array() {
	// var a = [2]int{1, 2}
	a := [2]int{1, 2}
	a[0] = 10
	b := [2]string{"hello", "world"}
	fmt.Println(a, b)
}

func main() {
	sum := 0
	// var sum int = 0
	for sum < 10 {
		sum += 1
		/*
			if sum > 5 {
				fmt.Println("Value is higher\n")
				return
			}
		*/
	}
	fmt.Printf("%T type of variable with valye %v\n", sum, sum)
	fmt.Println(
		pow(3, 2, 5),
		pow(3, 2, 20),
	)

	// defer println(day(time.Now().Weekday()))    //defere statement----executed but returned in the end
	/*
		switch os := runtime.GOOS; os {					//example of switch statement
		case "darwin":
			fmt.Println(os)
		default:
			fmt.Println("default")
		}
	*/

	//	Use of pointer
	/*
		var p *int
		i := 42
		p = &i
		fmt.Println(*p)
		*p = 21
		fmt.Println(*p)
	*/

	// use of Struct
	var (
		v1 = Vertex{2, 3}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p  = &Vertex{4, 5}
	)
	fmt.Println(v1, v2, v3, p.X)
	array()

	a := []int{1, 2, 3, 4}
	printSlice(a)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
