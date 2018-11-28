package main

import (
	"fmt"
)

func main() {
	getScan()
}

func getScan() {
	//	var a, b, c string
	var s []int
	n, _ := fmt.Scanln(&s)
	for {
		if n == 0 {
			break
		} else {
			fmt.Println(n)
			break
		}
	}
}
