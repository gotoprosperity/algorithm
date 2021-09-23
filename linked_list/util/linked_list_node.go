package util

type Node struct {
	Next *Node
	Val  int
}

func GenerateLinkedList(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	root := &Node{}
	tmp := root
	for _, val := range vals {
		tmp.Next = &Node{Val: val}
		tmp = tmp.Next
	}
	return root.Next
}

func LengthOfLinkedList(node *Node) int {
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
