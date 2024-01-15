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
		p = aNode.Next
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
			break
		}
		p = p.Next
	}
	var (
		cur = aNode
	)
	if x >= cur.Val {
		for x > cur.Next.Val {
			if cur.Val < cur.Next.Val {
				cur = cur.Next
			} else if cur.Val > cur.Next.Val {
				break
			} else if cur.Val == cur.Next.Val {
				var (
					p    = cur
					mask int
				)
				for {
					// 找到不同即可
					if cur.Val != cur.Next.Val {
						mask = 1
						break
					}
					if p == cur.Next {
						// 套圈了，所有列表节点都一样
						mask = 2
						break
					}
					cur = cur.Next
				}
				if 2 == mask {
					break
				}
			}
		}
		node := &Node{}
		node.Val = x
		node.Next = cur.Next
		cur.Next = node
	} else {
		for 
		node := &Node{}
		node.Val = x
		node.Next = cur.Next
		cur.Next = node
	}

	return aNode
}
