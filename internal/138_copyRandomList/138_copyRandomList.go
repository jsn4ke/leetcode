package leetcode_138_copyRandomList

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var (
		hash  = map[*Node]*Node{}
		dummy = new(Node)
		p     = dummy
		q     = head
	)
	for nil != q {
		node := new(Node)
		node.Val = q.Val
		hash[q] = node
		p.Next = node
		q = q.Next
		p = p.Next
	}
	q = head
	for nil != q {
		hash[q].Random = hash[q.Random]
		q = q.Next
	}
	return dummy.Next
}
