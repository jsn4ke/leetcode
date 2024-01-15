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
		p    = aNode.Next
		mask int
	)
	for {
		if p.Val != p.Next.Val {
			break
		}
		if p == aNode {
			node := &Node{
				Val:  x,
				Next: nil,
			}
			node.Next = aNode
			aNode.Next = node
			return node
		}
		p = p.Next
	}
	var (
		cur = aNode
	)
	if x >= cur.Val {
		for x > cur.Next.Val {
			if cur.Val > cur.Next.Val {
				break
			}
		}
		node := &Node{}
		node.Val = x
		node.Next = cur.Next
		cur.Next = node
	} else {

	}

	return aNode
}
