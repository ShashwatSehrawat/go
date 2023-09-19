package main

import (
	"fmt"
	"os"
)

const s string = "Constant"

func main() {
	fmt.Fprintln(os.Stdout, s)
	fmt.Println(s)
}
