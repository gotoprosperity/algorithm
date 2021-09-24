package merge_k_sorted_lists

import (
	"container/heap"

	"github.com/gotoprosperity/algorithm/util"
)

func mergeKLists(lists []*util.ListNode) *util.ListNode {
	//return mergeKListsWithPQ(lists) // 优先队列
	return mergeKListsReduce(lists) // 归并合并
}

// 构造pq
type Item struct {
	Node *util.ListNode
	Idx  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Node.Val < pq[j].Node.Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Idx, pq[j].Idx = i, j
}

func (pq *PriorityQueue) Push(ln interface{}) {
	item := ln.(*util.ListNode)
	node := &Item{Node: item}
	n := len(*pq)
	node.Idx = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	(*pq)[n-1] = nil
	item.Idx = -1
	*pq = (*pq)[0 : n-1]
	return item.Node
}

//// 优先队列解法
func mergeKListsWithPQ(lists []*util.ListNode) *util.ListNode {
	pq := PriorityQueue{}
	// push
	for _, list := range lists {
		for list != nil {
			heap.Push(&pq, list)
			list = list.Next
		}
	}
	//
	head := &util.ListNode{}
	tmp := head
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*util.ListNode)
		tmp.Next = &util.ListNode{Val: item.Val}
		tmp = tmp.Next
	}
	return head.Next
}

//// 归并方式解法
func mergeKListsReduce(lists []*util.ListNode) *util.ListNode {
	if len(lists) == 0 {
		return nil
	}
	n := len(lists)
	for n > 1 {
		k := (n + 1) / 2
		for i := 0; i < n/2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[i+k])
		}
		n = k
	}
	return lists[0]
}

// 合并两个有序链表的原子方法
func mergeTwoLists(l1, l2 *util.ListNode) *util.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	ret := &util.ListNode{}
	tmp := ret
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return ret.Next
}
