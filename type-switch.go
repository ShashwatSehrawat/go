package main

import "fmt"

func do(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Integer Value:", i)
	case string:
		fmt.Println("String Value:", i)
	default:
		fmt.Printf("No operation for Type:%T\n", i)
	}
}

func main() {
	do(10)
	do("hello")
	do(true)

}
