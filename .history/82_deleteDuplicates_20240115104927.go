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
			p.Next = cur
			p = p.Next
			cur = cur.Next
		} else {
			for nil != cur.Next && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			cur = cur.Next
		}
	}

	return dummy.Next
}
