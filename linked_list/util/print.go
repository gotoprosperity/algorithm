package util

import "fmt"

func PrintLinkedList(node *Node) {
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
