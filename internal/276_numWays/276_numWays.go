package leetcode_276_numWays

func numWays(n int, k int) int {
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}
	var dp = make([]int, n+1)
	dp[0] = k
	dp[1] = k * k
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]*(k-1) + 1*dp[i-2]*(k-1)
	}
	return dp[n-1]
}
