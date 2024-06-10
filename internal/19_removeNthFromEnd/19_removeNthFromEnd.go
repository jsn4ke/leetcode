package leetcode_19_removeNthFromEnd

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		p, q = head, head
	)
	for i := 0; i < n; i++ {
		q = q.Next
	}
	if nil == q {
		return head.Next
	}
	for nil != q.Next {
		q = q.Next
		p = p.Next
	}
	p.Next = p.Next.Next
	return head
}
