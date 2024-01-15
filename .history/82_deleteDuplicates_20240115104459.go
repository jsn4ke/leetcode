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
		cur = head
		p   = dummy
	)

	for {
		if pre.Val != cur.Val {
			p.Next = pre
			p = pre
		} else {
			for nil != cur.Next && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
		}
	}

	return dummy.Next
}
