package leetcode_53_maxSubArray

func maxSubArray(nums []int) int {
	var (
		n   = len(nums)
		ans int
		dp  = make([]int, n)
	)
	dp[0] = nums[0]
	ans = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], 0) + nums[i]
		ans = max(ans, dp[i])
	}
	return ans
}
