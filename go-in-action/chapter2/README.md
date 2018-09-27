## 快速开始一个Go语言

本章学习内容

1. 学习如何写一个复杂的Go程序
2. 声明类型、变量、函数和方法
3. 启动并同步操作goroutine
4. 使用接口写通用的代码
5. 处理程序逻辑和错误

项目结构

```yaml
sample
	data
		data.json 	-- 包含一组数据源
	matchers
		rss.go 		-- 搜索rss源的匹配器
	search
		default.go 	-- 搜索数据用的默认匹配器
		feed.go 	-- 用于读取json数据文件
		match.go 	-- 用于支持不同匹配器的接口
		search.go 	--执行搜索的主控制逻辑
	main.go 		-- 程序的入口
```

