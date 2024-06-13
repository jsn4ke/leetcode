package leetcode_300_lengthOfLIS

func lengthOfLIS(nums []int) int {
	// dp[i] = max (dp[i] , dp[j]+1) if nums[j] > nums[i]
	var (
		n   = len(nums)
		ans int
		dp  = make([]int, n)
	)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i], 1)
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j-1]+1)
			}
		}
	}
	for _, v := range dp {
		ans = max(ans, v)
	}
	return ans
}
