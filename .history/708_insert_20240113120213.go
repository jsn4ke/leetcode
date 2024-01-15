package leatcode

type Node struct {
	Val  int
	Next *Node
}

func insert_708(aNode *Node, x int) *Node {
	if nil == aNode {
		return &Node{
			Val:  x,
			Next: nil,
		}
	}
}
