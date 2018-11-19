package insertion

// 插入排序

func InsertSort(arr []int) []int {
	len := len(arr)
	var preindex, current int
	for i := 1; i < len; i++ { // 注意 i < len
		preindex = i - 1 // 下标为0的元素默认为已排序
		current = arr[i] // 从下标为1开始
		for preindex >= 0 && arr[preindex] > current {
			arr[preindex+1] = arr[preindex] // 不断从后往前遍历,如果比它大则往后移
			preindex--
		}
		arr[preindex+1] = current // 如果找到前一个比当前元素小,则插入当前的数组位置
	}
	return arr
}
