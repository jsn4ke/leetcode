package leetcode_2786_maxScore

import "math"

func maxScore(nums []int, x int) int64 {
	// dp[i] = dp[j]  - x || dp[j] + i
	var (
		n   = len(nums)
		ans int64
		dp  [2]int64
	)
	for i := range dp {
		dp[i] = math.MinInt32
	}
	dp[0] = int64(nums[0])
	for j := 1; j < n; j++ {
		val := int64(nums[j])
		idx := val % 2
		tmp := max(dp[idx]+val, dp[1-idx]+val-int64(x))
		dp[idx] = max(dp[idx], tmp)
		ans = max(ans, dp[idx])
	}
	return ans
}
