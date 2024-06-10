package leetcode_234_isPalindrome

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var (
		arr []int
	)
	for nil != head {
		arr = append(arr, head.Val)
		head = head.Next
	}
	var (
		left, right = 0, len(arr)
	)
	for left < right {
		if arr[left] != arr[right-1] {
			return false
		}
		left++
		right--
	}
	return true
}
