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
	for {
		for nil != cur.Next && cur.Val != cur.Next.Val {
			cur = cur.Next
		}
	}
}
