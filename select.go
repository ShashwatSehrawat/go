package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	select {
	case c <- x:
		x, y = y, x+y
	case <-quit:
		fmt.Println("quit")
		return
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	//go fibonacci(cap(c), c)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	//for i := range c {
	//	fmt.Println(i)
	fibonacci(c, quit)
}
