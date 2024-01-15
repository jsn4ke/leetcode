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
		pre := cur
		cur := cur.Next
		var dup *Node
		// 往后插，直到头了，或者小于下一个
		for x > cur.Val {
			// 正常递增迭代
			if pre.Val < cur.Val {
			} else if pre.Val == cur.Val {
				// 重复
				if nil == dup {
					dup = cur
				} else if dup == cur {
					break
				}
			} else if pre.Val > cur.Val {
				// 到头了
				break
			}
			pre = cur
			cur = cur.Next
		}
		node := &Node{
			Val:  x,
			Next: nil,
		}
		node.Next = pre.Next
		pre.Next = node
	} else {
		start := cur
		// 找到开头
		for cur.Val >= cur.Next.Val && cur.Next != start {
			cur = cur.Next
		}

	}

	return aNode
}
