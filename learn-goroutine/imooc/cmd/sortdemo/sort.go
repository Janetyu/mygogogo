package main

import (
	"sort"
	"fmt"
)

func main() {
	// Creates a slice of int
	a := []int{3,6,2,1,9,10,8}
	sort.Ints(a)

	for i, v := range a {
		fmt.Println(i,v)
	}
}
