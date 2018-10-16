package main

import (
	"fmt"
	"time"
)

/**
 * FIFO first In first Out
 * 头{[hello 1], [hello 2], [hello 3]}
 * 头{[hello 2], [hello 3], [hello 4]} [hello 1 I‘m goroutine]
 */

func sample(message chan string) {
	message <- "helloworld 1"
	message <- "helloworld 2"
	message <- "helloworld 3"
	message <- "helloworld 4"
}

func sample2(message chan string) {
	time.Sleep(2 * time.Second)
	str := <-message
	str = str + " I'm goroutine"
	message <- str
	close(message) // 明确close要用在通道结束的地方
}

func main() {
	var message = make(chan string, 3)
	go sample(message)
	go sample2(message)
	time.Sleep(3 * time.Second)
	for i := range message { // 用range关键字扫描通道时，注意要关闭通道
		fmt.Println(i)
	}
	//	fmt.Println(<-message)
	//	fmt.Println(<-message)
	//	fmt.Println(<-message)
	//	fmt.Println(<-message)
	fmt.Println("hello world!!")
}
