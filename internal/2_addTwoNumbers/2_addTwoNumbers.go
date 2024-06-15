package leetcode_2_addTwoNumbers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	reverse := func(node *ListNode) *ListNode {
		var pre *ListNode
		for nil != node {
			tmp := node.Next
			node.Next = pre
			pre = node
			node = tmp
		}
		return pre
	}
	_ = reverse
	// l1 = reverse(l1)
	// l2 = reverse(l2)
	var (
		remain int
		dummy  = new(ListNode)
		p      = dummy
	)

	for nil != l1 && nil != l2 {
		node := new(ListNode)
		sum := l1.Val + l2.Val + remain
		remain = sum / 10
		node.Val = sum % 10
		p.Next = node
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if nil == l1 {
		l1 = l2
	}
	for 0 != remain || nil != l1 {
		node := new(ListNode)
		sum := remain
		if nil != l1 {
			sum += l1.Val
			l1 = l1.Next
		}
		remain = sum / 10
		node.Val = sum % 10
		p.Next = node
		p = p.Next
	}
	return dummy.Next
}
