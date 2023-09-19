package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type IPAddr [4]byte

func (p Person) String() string {
	return fmt.Sprintf("%v (%d years)\n", p.Name, p.Age)
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v:%v:%v:%v\n", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	a := Person{"hari", 100}
	b := Person{"Ram", 20}
	fmt.Println(a, b)
	hosts := map[string]IPAddr{
		"localhost": {127, 0, 0, 0},
		"goolgeDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v :: %v", name, ip)
	}
}
