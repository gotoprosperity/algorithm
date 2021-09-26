package quick_sort

import (
	"testing"

	"github.com/gotoprosperity/algorithm/util"
)

func TestQuickSort1(t *testing.T) {
	var (
		arr []int
	)
	arr = []int{9, 7, 3, 2, 1, 5, 6, 4, 8}
	QuickSort1(arr)
	if !util.IsIntArrayEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("wrong arr with %v", arr)
	}
}

func TestQuickSort2(t *testing.T) {
	var (
		arr []int
	)
	arr = []int{9, 7, 3, 2, 1, 5, 6, 4, 8}
	QuickSort2(arr)
	if !util.IsIntArrayEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("wrong arr with %v", arr)
	}
}
