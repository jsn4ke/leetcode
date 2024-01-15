package leatcode

type Node struct {
	Val  int
	Next *Node
}

func insert_708(aNode *Node, x int) *Node {
	if nil == aNode {
		aNode = &Node{
			Val:  x,
			Next: nil,
		}
		aNode.Next = aNode
		return aNode
	}
	if aNode.Next == aNode {
		node := &Node{
			Val:  x,
			Next: nil,
		}
		node.Next = aNode
		aNode.Next = node
		return aNode
	}
	var (
		cur = aNode.Next
	)
	for {
		if cur.Val != cur.Next.Val {
			break
		}
		if cur == aNode {
			node := &Node{
				Val:  x,
				Next: nil,
			}
			node.Next = cur.Next
			cur.Next = node
			return aNode
		}
		cur = cur.Next
	}
	if x >= cur.Val {
		for x >= cur.Next.Val && cur.Val <= cur.Next.Val {
			cur = cur.Next
		}
	} else {
		for cur.Val <= cur.Next.Val {
			cur = cur.Next
		}
		for x > cur.Next.Val {
			cur = cur.Next
		}
	}
	node := &Node{}
	node.Val = x
	node.Next = cur.Next
	cur.Next = node
	return aNode
}
