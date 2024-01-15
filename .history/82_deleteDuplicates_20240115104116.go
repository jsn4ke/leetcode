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
		p   = dummy
		cur = head
	)
	for nil != cur {
		for nil != cur.Next && cur.Val != cur.Next.Val {
			cur = cur.Next
		}
		p.Next = cur
		p = cur
	}
}
