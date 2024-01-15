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
	if x > cur.Val {

		// 往后插，直到头了，或者小于下一个
		for x > cur.Val || cur.Val >= cur.Next.Val {
			cur = cur.Next
		}
	}

	return aNode
}
