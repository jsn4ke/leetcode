package leatcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var (
		dummy = &ListNode{}
		p     = dummy
	)
	dummy.Next = head

	for nil != p {
		if p.Val == p.Next.Val {

		}
	}
}
