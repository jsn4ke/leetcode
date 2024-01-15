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
		return node
	}
	var (
		cur = aNode
	)
	if x >= cur.Val {
		start := cur
		// 往后插，直到头了，或者小于下一个
		for x > cur.Next.Val || cur.Val >= cur.Next.Val && cur.Next != start {
			cur = cur.Next
		}
		node := &Node{
			Val:  x,
			Next: nil,
		}
		node.Next = cur.Next
		cur.Next = node
	} else {
		start := cur
		for cur.Val >= cur.Next.Val && cur.Next != start {
			cur = cur.Next
		}
	}

	return aNode
}
