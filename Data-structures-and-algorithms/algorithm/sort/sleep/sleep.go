package main

import (
	"fmt"
	"time"
)

// 睡排，用协程编写的，有待完善

func sleepSort(arr []int) {
	c := make(chan int)

	for _, value := range arr {
		go func(val int) {
			time.Sleep(time.Duration(int64(val)) * time.Second)
			fmt.Print(" ", val)
			c <- val
		}(value)
	}

	for _ = range c {
		<-c
	}
}

func main() {
	a := []int{2, 3, 1, 4}
	go sleepSort(a)
	time.Sleep(5 * time.Second)
	//	fmt.Println(a)
}
