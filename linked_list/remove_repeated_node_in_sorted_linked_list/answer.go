package remove_repeated_node_in_sorted_linked_list

import "github.com/gotoprosperity/algorithm/util"

func RmNode(root *util.ListNode) {
	if root == nil {
		return
	}
	tmp := root
	for {
		if tmp.Next == nil {
			return
		}
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return
}
