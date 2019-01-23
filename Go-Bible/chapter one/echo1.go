// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	// Args[0] 是命令本身的名字
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
