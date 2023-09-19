package main

import "fmt"

func sum(s []int, c chan int) {

	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //send sum to channel c
}

func main() {

	s := []int{7, 2, 8, -9, 10}

	c := make(chan int)

	go sum(s[:len(s)/2], c)

	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c //receive from c

	fmt.Println(x, y, x+y)

	b := make(chan int, 2)
	b <- 2
	b <- 1
	fmt.Println(<-b)
	fmt.Println(<-b)
	b <- 3
	fmt.Println(<-b)
}
