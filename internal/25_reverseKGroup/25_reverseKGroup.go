package leetcode_25_reverseKGroup

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		n     int
		p     = head
		dummy = new(ListNode)
	)
	for nil != p {
		n++
		p = p.Next
	}
	var reverse = func(node *ListNode) *ListNode {
		var (
			dummy = new(ListNode)
		)
		for nil != node {
			tmp := node.Next

			pre := dummy.Next
			node.Next = pre
			dummy.Next = node

			node = tmp
		}
		return dummy.Next
	}
	p = dummy
	cur := head
	for nil != cur {
		j := 1
		tail := cur
		for ; j < k && nil != tail.Next; j++ {
			tail = tail.Next
		}
		if j != k {
			p.Next = cur
			break
		}
		next := tail.Next
		tail.Next = nil
		p.Next = reverse(cur)
		for nil != p.Next {
			p = p.Next
		}
		cur = next
	}
	return dummy.Next
}
