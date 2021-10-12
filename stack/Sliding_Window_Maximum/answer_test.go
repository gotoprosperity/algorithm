package Sliding_Window_Maximum

import (
	"testing"

	"github.com/gotoprosperity/algorithm/util"
)

func TestMaxSlidingWindow(t *testing.T) {
	if ret := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3); !util.IsIntArrayEqual([]int{3, 3, 5, 5, 6, 7}, ret) {
		t.Errorf("wrong ret with %v", ret)
	}
	if ret := maxSlidingWindow([]int{1, -1}, 1); !util.IsIntArrayEqual([]int{1, -1}, ret) {
		t.Errorf("wrong ret with %v", ret)
	}
	if ret := maxSlidingWindow([]int{7, 2, 4}, 2); !util.IsIntArrayEqual([]int{7, 4}, ret) {
		t.Errorf("wrong ret with %v", ret)
	}
}
