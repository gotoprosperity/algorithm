package find_Kth_Largest

import "sort"

func findKthLargest(nums []int, k int) int {
	return findKthLargestAfterSortingAllElements(nums, k)
}

// 全排序然后取第K位，数据量大的时候效率不好，因为实际上不需要关注所有元素的排序
func findKthLargestAfterSortingAllElements(nums []int, k int) int {
	sort.Ints(nums) // 升序
	return nums[len(nums)-k]
}
