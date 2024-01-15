package leetcode_206_reverseRecursive

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var (
		dummy = &ListNode{
			Val: 0,
		}
	)
	for nil != head {
		q := head.Next
		head.Next = dummy.Next
		dummy.Next = head
		head = q
	}
	return dummy.Next
}
