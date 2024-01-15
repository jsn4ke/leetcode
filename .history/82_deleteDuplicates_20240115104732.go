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
		cur = dummy
		p   = dummy
	)

	for nil != cur.Next {

	}

	return dummy.Next
}
