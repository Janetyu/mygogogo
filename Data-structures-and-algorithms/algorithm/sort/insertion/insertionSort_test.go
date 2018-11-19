package insertion

import (
	"testing"
)

func Test_insertSort(t *testing.T) {
	arr := []int{1, 5, 23, 6, 3, 1, 7}
	arr = InsertSort(arr)
	t.Log(arr)
}
