package bubble

// 冒泡排序
func BubbleSort(arr []int) []int {
	len := len(arr)
	var temp int
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
