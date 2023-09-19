package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello! how are you")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("n=%v err=%v bytes:: %v value:: %s\n", n, err, b, b)
		fmt.Printf("b[:%v]=%q\n", n, b[:n])
	}
}
