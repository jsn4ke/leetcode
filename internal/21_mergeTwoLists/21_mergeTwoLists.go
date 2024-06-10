package leetcode_21_mergeTwoLists

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		head = new(ListNode)
		q    = head
	)
	for nil != list1 && nil != list2 {
		if list1.Val > list2.Val {
			list1, list2 = list2, list1
		}
		q.Next = list1
		list1 = list1.Next
		q = q.Next
	}
	if nil == list1 {
		list1 = list2
	}
	q.Next = list1
	return head.Next
}
