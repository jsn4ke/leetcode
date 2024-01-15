package leatcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var (
		dummy = &ListNode{}
		cur   = head
	)
	dummy.Next = head
	for nil != cur {
		last := cur
		for nil != cur.Next && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

	}
}
