package delete_nth_node_from_tail

import (
	"github.com/gotoprosperity/algorithm/linked_list/util"
)

func delNode(root *util.Node, n int) *util.Node {
	if n < 1 || root == nil {
		return nil
	}
	fast, slow := root, root
	// 先跑快指针
	step := 0
	for {
		if fast.Next == nil { // 遍历到最后一个节点，停止
			break
		}
		if step == n { // 快指针跑够数目，可以停止然后一起跑了
			break
		}
		fast = fast.Next
		step++
	}
	// 判断链表长度是否足够
	if step == n-1 { // 删除第一个
		root = root.Next
		return root
	} else if step < n-1 { // 长度不够，无法操作
		return nil
	}
	// 现在快慢指针一起往后走
	for {
		if fast.Next == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	// 此时慢指针的下一个node就是需要删除的节点
	// 因为删除操作需要知道删除节点的pre节点，所以慢指针定位到这个pre节点上

	// 执行删除
	slow.Next = slow.Next.Next
	return root
}
