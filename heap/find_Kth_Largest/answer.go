package find_Kth_Largest

import (
	"container/heap"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	//return findKthLargestAfterSortingAllElements(nums, k)
	return findKthLargestWithHeap(nums, k)
}

// 全排序然后取第K位，数据量大的时候效率不好，因为实际上不需要关注所有元素的排序
func findKthLargestAfterSortingAllElements(nums []int, k int) int {
	sort.Ints(nums) // 升序
	return nums[len(nums)-k]
}

// 优先队列/堆
func findKthLargestWithHeap(nums []int, k int) int {
	h := IntHeap(nums)
	heap.Init(&h)
	var ret int
	for i := 0; i < k; i++ {
		ret = heap.Pop(&h).(int)
	}
	return ret
}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
