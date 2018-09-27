package search

import (
	"log"
	"sync"
)

// 注册用于搜索的匹配器的映射
var matchers = make(map[string]Matcher)

// Run执行逻辑,searchTerm是要搜索的搜索项
func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	// 构造一个waitGroup，以便处理所有的数据源
	// WaitGroup是一个计数信号量，利用它来统计所有的goroutine是否都完成了工作
	var waitGroup sync.WaitGroup

	// 设置需要等待处理
	// 每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		//获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		//启动一个goroutine来执行搜索
		//goroutine是一个独立于其他函数运行的函数
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			// WaitGroup的值没有作为参数传入匿名函数，但匿名函数依旧访问到了这个值
			// 这里运用了闭包，searchTerm和results也是同样的道理
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个goroutine来监控是否所有的工作都做完了
	go func() {
		// 使用了闭包来访问WaitGroup和results变量
		// 等待所有任务完成，Wait方法会导致goroutine阻塞，直到WaitGroup内部的计数到达0
		waitGroup.Wait()

		// 用关闭通道的方式，通知Display函数
		// 可以退出程序了
		close(results)
	}()

	// 启动函数，显示返回的结果，并且在最后一个结果显示完后返回
	Display(results)
}

// Register调用时，会注册一个匹配器，提供给后面的程序使用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
