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
		pre = dummy
		p   = dummy
	)

	for nil != pre {
		cur := pre.Next
		if pre.Val != cur.Val {
			p.Next = pre
			p = pre
			pre = cur
		} else {
			for nil != cur.Next && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			pre = cur.Next
		}
	}

	return dummy.Next
}
