package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	in := "1 2 3"
	result := outOrder(strings.Fields(in))
	dictSort(result)
	//	s := format(result)
	fmt.Println(result)

}

//输入trainsNums，返回全部排列
//如输入[1 2 3]，则返回[123 132 213 231 312 321]
func outOrder(trainsNums []string) []string {
	COUNT := len(trainsNums)
	//检查
	if COUNT == 0 || COUNT > 10 {
		panic("Illegal argument. trainsNums size must between 1 and 9.")
	}

	//如果只有一个数，则直接返回
	if COUNT == 1 {
		return []string{trainsNums[0]}
	}

	//否则，将最后一个数插入到前面的排列数中的所有位置（递归）
	return insert(outOrder(trainsNums[:COUNT-1]), trainsNums[COUNT-1])
}

func insert(res []string, insertNum string) []string {
	//保存结果的slice
	result := make([]string, len(res)*(len(res[0])+1))

	index := 0
	for _, v := range res {
		for i := 0; i < len(v); i++ {
			//在v的每一个元素前面插入
			result[index] = v[:i] + insertNum + v[i:]
			index++
		}

		//在v最后面插入
		result[index] = v + insertNum
		index++
	}

	return result
}

//按字典顺序排序
func dictSort(res []string) {
	sort.Strings(res)
}

//按指定格式输出
//func format(res []string) string {
//	strs := make([]string, len(res))
//	for i := 0; i < len(res); i++ {
//		strs[i] = addWhiteSpace(res[i])
//	}

//	return strings.Join(strs, " ")
//}

////添加空格
//func addWhiteSpace(s string) string {
//	var retVal string
//	for i := 0; i < len(s); i++ {
//		retVal += string(s[i])

//		if i != len(s)-1 {
//			retVal += " "
//		}
//	}

//	return retVal
//}
