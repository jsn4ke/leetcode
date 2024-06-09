package leetcode_312_maxCoins

func maxCoins(nums []int) int {
	var (
		n   = len(nums)
		arr = make([]int, 2+n)
	)
	arr[0], arr[n+1] = 1, 1
	copy(arr[1:], nums)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for i := n; i >= 0; i-- {
		for j := i + 1; j < len(arr); j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+arr[k-1]*arr[k]*arr[k+1])
			}
		}
	}
	return dp[0][n]
}
