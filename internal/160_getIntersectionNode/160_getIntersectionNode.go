package leetcode_160_getIntersectionNode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		p, q = headA, headB
	)
	if nil == p || nil == q {
		return nil
	}
	for nil != q.Next {
		q = q.Next
	}
	for nil != p.Next {
		p = p.Next
	}
	if p != q {
		return nil
	}
	first := true
	p, q = headA, headA
	for first || p != q {
		first = false
		p = p.Next
		if nil == p {
			p = headB
		}
		if nil == q.Next {
			q = headB.Next
		} else {
			q = q.Next.Next
		}
		if nil == q {
			q = headB
		}
	}
	q = headA
	for p != q {
		p = p.Next
		q = q.Next
		if nil == p {
			p = headB
		}
		if nil == q {
			q = headB
		}
	}
	return q
}
