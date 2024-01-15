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
		cur = dummy.Next
		p   = dummy
	)

	for nil != cur && nil != cur.Next {
		if cur.Val != cur.Next.Val {

		}
	}

	return dummy.Next
}
