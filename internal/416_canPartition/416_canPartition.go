package leetcode_416_canPartition

import "sort"

func canPartition(nums []int) bool {
	var (
		n    = len(nums)
		half int
	)
	for i := 0; i < n; i++ {
		half += nums[i]
	}
	if half&1 == 1 {
		return false
	}
	half >>= 1
	sort.Ints(nums)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, half+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		val := nums[i]
		for cost := 1; cost <= half; cost++ {
			if cost < val {
				dp[i][cost] = dp[i-1][cost]
			} else {
				dp[i][cost] = dp[i-1][cost] || dp[i-1][cost-val]
			}
		}
	}
	return dp[n-1][half]
}
