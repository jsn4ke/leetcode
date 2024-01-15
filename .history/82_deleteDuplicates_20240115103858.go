package leatcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var (
		dummy = &ListNode{
			Val:  -101,
			Next: head,
		}
		last = dummy
		cur  = head
	)

	for nil != cur {
		for nil != cur && cur.Val == last.Val {
			cur = cur.Next
		}
	}
}
