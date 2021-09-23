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
