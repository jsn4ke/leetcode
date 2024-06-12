package leetcode_105_buildTree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var (
		n    = len(preorder)
		hash = make(map[int]int, n)
	)
	for i, v := range inorder {
		hash[v] = i
	}
	var dfs func(p1, p2, i1, i2 int) *TreeNode
	dfs = func(p1, p2, i1, i2 int) *TreeNode {
		if i1 >= i2 {
			return nil
		}
		val := preorder[p1]
		idx := hash[val]
		root := new(TreeNode)
		root.Val = val
		leftNum := idx - i1
		root.Left = dfs(p1+1, p1+1+leftNum, i1, idx)
		root.Right = dfs(p1+1+leftNum, p2, idx+1, i2)
		return root
	}
	return dfs(0, n, 0, n)
}
