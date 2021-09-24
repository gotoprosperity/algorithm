package find_k_closest_elem

import (
	"testing"

	"github.com/gotoprosperity/algorithm/util"
)

func TestFindClosestElements(t *testing.T) {
	var (
		arr, ret []int
		k, x     int
	)
	arr = []int{1, 2, 3, 4, 5}
	k, x = 4, 3
	ret = findClosestElements(arr, k, x)
	if !util.IsIntArrayEqual(ret, []int{1, 2, 3, 4}) {
		t.Errorf("wrong ret is %v", ret)
	}
	arr = []int{1, 1, 1, 10, 10, 10}
	k, x = 1, 9
	ret = findClosestElements(arr, k, x)
	if !util.IsIntArrayEqual(ret, []int{10}) {
		t.Errorf("wrong ret is %v", ret)
	}
	arr = []int{1, 1, 2, 2, 2, 2, 2, 3, 3}
	k, x = 3, 3
	ret = findClosestElements(arr, k, x)
	if !util.IsIntArrayEqual(ret, []int{2, 3, 3}) {
		t.Errorf("wrong ret is %v", ret)
	}
}
