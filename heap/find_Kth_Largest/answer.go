package find_Kth_Largest

import (
	"container/heap"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	//return findKthLargestAfterSortingAllElements(nums, k)
	//return findKthLargestWithHeap(nums, k)
	return findKthLargestWithPartion(nums, k)
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

// 快排partion思路，持续找一个中间值，当这个中间值左边的数为k-1时，这个数即为第k大的数，不需要全排序
func findKthLargestWithPartion(nums []int, k int) int {
	low, high := 0, len(nums)
	for low < high {
		i, j := low, high-1
		pivot := nums[low] // 取low位置的值
		for i <= j {
			for i <= j && nums[i] >= pivot {
				i++
			}
			for i <= j && nums[j] <= pivot {
				j--
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[low], nums[j] = nums[j], nums[low]
		if j == k-1 {
			return nums[j]
		} else if j < k-1 {
			low = j + 1
		} else {
			high = j
		}
	}
	return 0
}
