package remove_repeated_node_in_sorted_linked_list

import (
	"testing"

	"github.com/gotoprosperity/algorithm/linked_list/util"
)

func TestRmNode(t *testing.T) {
	var (
		l      *util.Node
		length int
	)
	// case 1
	l = util.GenerateLinkedList([]int{1, 2, 3})
	RmNode(l)
	length = util.LengthOfLinkedList(l)
	if length != 3 {
		t.Error("length is not 3")
	}
	util.PrintLinkedList(l)
	// case 2
	l = util.GenerateLinkedList([]int{1, 1, 2, 3})
	RmNode(l)
	length = util.LengthOfLinkedList(l)
	if length != 3 {
		t.Error("length is not 3")
	}
	util.PrintLinkedList(l)
	// case 3
	l = util.GenerateLinkedList([]int{1, 1, 2, 2, 2, 2, 2, 3, 4, 4, 4})
	RmNode(l)
	length = util.LengthOfLinkedList(l)
	if length != 4 {
		t.Error("length is not 4")
	}
	util.PrintLinkedList(l)
}
