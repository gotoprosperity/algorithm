package quick_sort

import "fmt"

/*
	快速排序写法
	这里以go实现方式为例
*/

// 写法一：递归；在目标值两侧，依次找需要移动的，类似拉锁，最后合并到目标值上。
func QuickSort1(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSort1(arr, 0, len(arr)-1)
}

func quickSort1(arr []int, left, right int) {
	i, j := left, right
	pos := left
	pivot := arr[pos]
	for i <= j {
		for j >= pos && arr[j] >= pivot {
			j--
		}
		if j >= pos {
			arr[pos] = arr[j] // pos上改为找到的值
			pos = j           // 位置换到右边去
		}
		for i <= pos && arr[i] <= pivot {
			i++
		}
		if i <= pos {
			arr[pos] = arr[i]
			pos = i
		}
	}
	arr[pos] = pivot // 这一轮确定的pos位置，写入这一轮的标记元素

	// 如果两边有未处理的子集，分别做这件事
	if pos-left > 1 {
		quickSort1(arr, left, pos-1)
	}
	if right-pos > 1 {
		quickSort1(arr, pos+1, right)
	}
}

// 写法二：递归；从左边开始，遍历范围内所有值，丢到目标值左边或者右边。
func QuickSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mid, pos := arr[0], 1
	head, tail := 0, len(arr)-1
	fmt.Println("begin", arr, head, tail, mid, pos)
	for head < tail {
		if arr[pos] > mid {
			arr[pos], arr[tail] = arr[tail], arr[pos] // 大于标记位的，都丢到最后
			tail--
		} else {
			arr[pos], arr[head] = arr[head], arr[pos] // 小于等于标记位的，都丢到最前面
			head++
			pos++ // 不能丢，随着往前放，pos的值要向后移动判断
		}
		fmt.Println("pos", pos, arr)
	}
	arr[head] = mid // 一轮循环完，定位元素位置确认
	fmt.Println("end", arr, head, tail, mid, pos)
	// 两边接着来
	QuickSort2(arr[:head])
	QuickSort2(arr[head+1:])
}
