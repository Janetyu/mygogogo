// Echo2 prints its command-line arguments.
package main

import (
	"os"
	"fmt"
	"strconv"
	"time"
)

func main()  {
	start := time.Now()
	s, sep := "", " "
	// +=  连接原字符串、空格和下个参数，产生新字符串, 并把它赋值给 s
	// s 原来的内容已经不再使用，将在适当时机对它进行垃圾回收
	for index, arg := range os.Args[1:] {
		s += "index: " + strconv.Itoa(index) + sep + "value: " + arg + "\n"
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
