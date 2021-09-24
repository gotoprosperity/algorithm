package util

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

func GenerateLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	root := &ListNode{}
	tmp := root
	for _, val := range vals {
		tmp.Next = &ListNode{Val: val}
		tmp = tmp.Next
	}
	return root.Next
}

func LengthOfLinkedList(node *ListNode) int {
	length := 0
	if node == nil {
		return length
	}
	for {
		length++
		node = node.Next
		if node == nil {
			break
		}
	}
	return length
}

func PrintLinkedList(node *ListNode) {
	if node == nil {
		fmt.Println("empty Linked List")
		return
	}
	for {
		if node == nil {
			break
		}
		fmt.Printf("%d,", node.Val)
		node = node.Next
	}
	fmt.Println()
}
