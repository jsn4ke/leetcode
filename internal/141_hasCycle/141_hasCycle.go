package leetcode_141_hasCycle

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if nil == head || nil == head.Next {
		return false
	}
	var (
		slow = head
		fast = head
	)
	for nil != slow && nil != fast && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
