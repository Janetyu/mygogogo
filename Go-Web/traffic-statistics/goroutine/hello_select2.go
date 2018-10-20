package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
 * select等待机制
 */

func sample(ch chan string) {
	for i := 0; i < 19; i++ {
		ch <- "I'm sample1 num: " + strconv.Itoa(i)
		time.Sleep(3 * time.Second)
	}
}

func sample2(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(60 * time.Second)
	}
}

func main() {
	ch1 := make(chan string, 3)
	ch2 := make(chan int, 5)
	for i := 0; i < 10; i++ {
		go sample(ch1)
		go sample2(ch2)
	}

	// select 会对通道阻塞
	for {
		select {
		case str, ch1Check := <-ch1:
			if !ch1Check {
				fmt.Println("ch1 failed")
			}
			fmt.Println(str)
		case p, ch2Check := <-ch2:
			if !ch2Check {
				fmt.Println("ch2 failed")
			}
			fmt.Println(p)
		}
	}
}
