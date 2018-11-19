package main

import (
	"fmt"
	"time"
)

func main() {
	tab := []int{1, 3, 0, 5}

	ch := make(chan int)
	for _, value := range tab {
		go func(val int) {
			time.Sleep(time.Duration(int64(val)) * 10000000)
			fmt.Print(val, " ")
			ch <- val
		}(value)
	}
	for _ = range tab {
		<-ch
	}
}
