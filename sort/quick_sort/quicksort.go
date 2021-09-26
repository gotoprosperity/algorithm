package quick_sort

/*
	快速排序写法
	这里以go实现方式为例
*/

// 写法一：携带左右边界递归
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
