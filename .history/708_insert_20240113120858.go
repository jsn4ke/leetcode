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
	for {
		// 头了
		if cur.Next.Val < cur.Val {
			if x < cur.Next.Val {
				node := &Node{
					Val:  x,
					Next: nil,
				}
				node.Next = cur.Next
				cur.Next = node
				return
			}
		}
	}
	return aNode
}
