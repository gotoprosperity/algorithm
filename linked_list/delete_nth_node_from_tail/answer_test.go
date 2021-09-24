package delete_nth_node_from_tail

import (
	"testing"

	"github.com/gotoprosperity/algorithm/util"
)

func TestDelNode(t *testing.T) {
	var l, ret *util.ListNode
	l = util.GenerateLinkedList([]int{1, 2, 3})
	ret = delNode(l, 1)
	util.PrintLinkedList(ret)
	if ret.Next.Next != nil {
		t.Error("error")
	}
	l = util.GenerateLinkedList([]int{1, 2, 3})
	ret = delNode(l, 2)
	util.PrintLinkedList(ret)
	if ret.Next.Next != nil && ret.Next.Val != 3 {
		t.Error("error")
	}
	l = util.GenerateLinkedList([]int{1, 2, 3})
	ret = delNode(l, 3)
	util.PrintLinkedList(ret)
	if ret.Next.Next != nil && ret.Val == 1 {
		t.Error("error")
	}
}
