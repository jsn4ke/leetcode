package leetcode_32_longestValidParentheses

func longestValidParentheses(s string) int {
	var (
		n   = len(s)
		dp  = make([][]bool, n)
		ans int
	)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n-1; i++ {
		if s[i] == '(' && s[i] != s[i+1] {
			dp[i][i+1] = true
		}
	}
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			continue
		}
		var last = -1
		for j := i - 3; j >= 0; j-- {
			dp[j][i] =
				dp[j][i] ||
					(s[j] == '(' && s[j] != s[i] && dp[j+1][i-1]) ||
					(dp[i-1][i] && dp[j][i-2]) ||
					(dp[j][j+1] && dp[j+2][i])
			if dp[j][i] {
				last = j
			} else {
				if -1 != last {
					dp[j][i] = dp[j][last-1] && dp[last][i]
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if dp[i][j] {
				ans = max(ans, j-i+1)
			}
		}
	}
	return ans
}
