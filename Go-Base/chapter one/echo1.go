package main

import (
	"fmt"
)

func main() {
	s, ep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
