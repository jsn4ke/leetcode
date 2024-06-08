package leetcode_560_subarraySum

func subarraySum1(nums []int, k int) int {
	// 前缀和
	preSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if preSum[j]-preSum[i] == k {
				count++
			}
		}
	}
	return count
}

func subarraySum(nums []int, k int) int {
	var (
		size             = len(nums)
		left, right, sum int
		count            int
	)
	for right < size {
		sum += nums[right]
		if sum == k {
			count++
		} else if sum > k {
			for left < right && sum > k {
				sum -= nums[left]
				left++
				if sum == k {
					count++
				}
			}
		}
		right++
	}
	return count
}
