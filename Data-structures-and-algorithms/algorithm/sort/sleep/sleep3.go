// test project main.go
package main

import (
	"fmt"
	"time"
)

var flag bool
var container chan bool
var count int

func main() {
	flag = true
	container = make(chan bool, 7)
	var arrays []int = []int{25, 12, 36, 24, 96, 54, 28}
	var i int
	for i = 0; i < len(arrays); i++ {
		go tosleep(arrays[i])
	}
	go listen(len(arrays))
	for flag {
		time.Sleep(1e9)
	}

}
func listen(size int) {
	for flag {
		select {
		case <-container:
			count++
			if count >= size {
				flag = false
				break
			}
		}
	}
}
func tosleep(data int) {
	time.Sleep(time.Duration(data))
	fmt.Println(data)
	container <- true
}
