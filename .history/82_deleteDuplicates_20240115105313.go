package leatcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var (
		dummy = &ListNode{
			Val:  -101,
			Next: nil,
		}
		cur = dummy.Next
		p   = dummy
	)

	for nil != cur {
		if nil == cur.Next || cur.Val != cur.Next.Val {
			p.Next = &ListNode{Val: cur.Val}
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
