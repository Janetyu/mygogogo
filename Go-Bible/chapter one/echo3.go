package main

import (
	"fmt"
	"strings"
	"os"
	"time"
)

func main()  {
	start := time.Now()
	// strings 不会生成新的字符串，对原字符串进行操作
	fmt.Println(strings.Join(os.Args[1:]," "))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
