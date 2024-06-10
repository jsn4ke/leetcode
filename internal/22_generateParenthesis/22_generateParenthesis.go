package leetcode_22_generateParenthesis

func generateParenthesis(n int) []string {
	var dfs func(i int, in string, stack int, remain int)
	var ans []string
	// 1.没有待匹配的，直接分配括号
	// 2.有待匹配的
	// 2.1.匹配括号
	// 2.2.添加括号（如果可以）
	dfs = func(i int, in string, stack, remain int) {
		if i == n*2 {
			ans = append(ans, in)
			return
		}
		if 0 == stack {
			dfs(i+1, in+"(", stack+1, remain-1)
		} else {
			dfs(i+1, in+")", stack-1, remain)
			if 0 != remain {
				dfs(i+1, in+"(", stack+1, remain-1)
			}
		}
	}
	dfs(0, "", 0, n)
	return ans
}
