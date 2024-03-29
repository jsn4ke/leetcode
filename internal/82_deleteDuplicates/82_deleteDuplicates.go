package leetcode_82_deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var (
		dummy = &ListNode{Val: -101, Next: nil}
		cur   = head
		p     = dummy
	)
	for nil != cur {
		if nil == cur.Next || cur.Val != cur.Next.Val {
			p.Next = cur
			p = p.Next
			cur = cur.Next
			p.Next = nil
		} else {
			for nil != cur.Next && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			cur = cur.Next
		}
	}
	return dummy.Next
}
