package main

import (
	"sync"
	"fmt"
)

func main() {
	users := []string{"aa","bb","cc","dd"}
	wg := sync.WaitGroup{}
	var mu sync.Mutex
	finished := make(chan bool, 1)
	wg.Add(lv en(users))
	var out []string
	for _, u := range users {
		go func(u string) {
			defer wg.Done()
			mu.Lock()
			out = append(out, u)
			mu.Unlock()
		}(u)
	}
	// 等 wg 为0时，关闭 finished 通道，退出并发操作
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <- finished:
	}

	// 按顺序复位所有userinfo
	for i, u := range out {
		fmt.Println(i,u)
	}
}