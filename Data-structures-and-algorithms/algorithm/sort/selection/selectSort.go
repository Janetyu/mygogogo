package selection

// 选择排序

func SelectSort(arr []int) []int {
	len := len(arr)
	var temp, minindex int
	for i := 0; i < len-1; i++ { // 注意 i < len - 1
		minindex = i                   // 把每一次要比较的i都当成是最小的
		for j := i + 1; j < len; j++ { // 注意 j < len
			if arr[j] < arr[minindex] {
				minindex = j // 不断往后遍历，找到最小的进行选择替换
			}
		}
		temp = arr[minindex]
		arr[minindex] = arr[i]
		arr[i] = temp
	}
	return arr
}
