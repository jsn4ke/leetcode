package leetcode_70_climbStairs

func climbStairs(n int) int {
	if 1 == n {
		return 1
	}
	var a, b int = 1, 1
	for i := 2; i <= n; i++ {
		// dp[i] = dp[i-1] + dp[i-2]
		tmp := a + b
		b = a
		a = tmp
	}
	return a
}
