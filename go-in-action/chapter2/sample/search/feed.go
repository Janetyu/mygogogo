package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	// 使用了tag标记，标记里描述了JSON解码的元数据
	// 每个标记将结构类型里的字段对应到json文档里指定名字的字段
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds 读取并反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 当函数返回时
	// 关闭文件
	defer file.Close()

	// 将文件解码到一个切片里
	// 这个切片的每一项是一个指向一个Feed类型值的指针
	var feeds []*Feed
	// func (dec *Decoder) Decode(v interface{}) error
	err = json.NewDecoder(file).Decode(&feeds)

	// 这个函数不需要检查错误，调用者会做这件事
	return feeds, err
}
