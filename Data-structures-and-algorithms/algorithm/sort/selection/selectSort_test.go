package selection

import (
	"testing"
)

func Test_selectSort(t *testing.T) {
	arr := []int{1, 5, 23, 6, 3, 1, 7}
	arr = SelectSort(arr)
	t.Log(arr)
}
