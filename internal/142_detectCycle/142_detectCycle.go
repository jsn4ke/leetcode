package leetcode_142_detectCycle

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	var (
		slow = head
		fast = head
	)
	for nil != slow {
		slow = slow.Next
		if nil == fast.Next {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			fast = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}
