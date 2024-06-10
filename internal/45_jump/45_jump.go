package leetcode_45_jump

func jump(nums []int) int {
	var (
		n  = len(nums)
		dp = make([]int, n)
	)
	for i := n - 2; i >= 0; i-- {
		dp[i] = -1
		for j := i + 1; j <= i+nums[i] && j < n; j++ {
			if -1 == dp[j] {
				continue
			}
			if -1 == dp[i] || dp[i] > 1+dp[j] {
				dp[i] = 1 + dp[j]
			}
		}
	}
	return dp[0]
}
