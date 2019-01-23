// 把全部输入数据读到内存中，一次分割为多行，然后处理它们。
package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFile  函数返回一个字节切片（byte slice），用string(data)转成string类型
		// bufio.Scanner 、ioutil.ReadFile 和 ioutil.WriteFile
		// 都使用 *os.File 的 Read 和 Write 方法
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data),"\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}