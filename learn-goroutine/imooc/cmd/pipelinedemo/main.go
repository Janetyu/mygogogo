package main

import (
	"mygithub/learn-goroutine/laonanhai/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {
	const smallfilename = "small.in"
	const largefilename = "large.in"
	const smalln  = 50
	const largen  = 100000000
	
	file,err := os.Create(smallfilename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建数据源
	p := pipeline.RandomSource(smalln)
	pipeline.WriterSink(bufio.NewWriter(file),p)

	file,err = os.Open(smallfilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 添加buffer通道，快速读取
	p = pipeline.ReaderSource(bufio.NewReader(file))
	for v := range p {
		fmt.Println(v)
	}
}

func mergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(
			pipeline.ArraySource(0, 8, 5, 3, 8, 6)))
	//for  {
	//	if num,ok := <-p; ok {
	//		fmt.Println(num)
	//	} else {
	//		break
	//	}
	//}
	for v := range p {
		fmt.Println(v)
	}
}
