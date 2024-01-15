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
		for x > cur.Next.Val {
			if cur.Val < cur.Next.Val {

			} else if cur.Val > cur.Next.Val {
				break
			} else if cur.Val == cur.Next.Val {
				var (
					p    = cur
					mask int
				)
				for {
					if cur.Val != cur.Next.Val {
						mask = 1
						break
					}
					if  p == cur.Next{
						mask = 2
						break
					}
					cur = cur.Next
				}
				if 2 == mask {
					break
				} else {
					cur = 
				}
			}
		}
		node := &Node{}
		node.Val = x
		node.Next = cur.Next
		cur.Next = node
	} else {
		start := cur
		// 找到开头
		for cur.Val >= cur.Next.Val && cur.Next != start {
			cur = cur.Next
		}

	}

	return aNode
}
