package bubble

import (
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	arr := []int{1, 5, 23, 6, 3, 1, 7}
	arr = BubbleSort(arr)
	t.Log(arr)
}
