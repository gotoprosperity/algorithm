package merge_k_sorted_lists

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKListsWithPQ(lists)
}

// 构造pq
type Item struct {
	Node *ListNode
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
	item := ln.(*ListNode)
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

// 优先队列解法
func mergeKListsWithPQ(lists []*ListNode) *ListNode {
	pq := PriorityQueue{}
	// push
	for _, list := range lists {
		for list != nil {
			heap.Push(&pq, list)
			list = list.Next
		}
	}
	//
	head := &ListNode{}
	tmp := head
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*ListNode)
		tmp.Next = &ListNode{Val: item.Val}
		tmp = tmp.Next
	}
	return head.Next
}