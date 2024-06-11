package leetcode_75_sortColors

func sortColors(nums []int) {
	var (
		n = len(nums)
	)
	var (
		low  = 0
		high = n - 1
		i    = low - 1
	)
	for j := low; j < high; j++ {
		if nums[j] < 1 {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	low = i + 1
	i = low - 1
	for j := low; j < high; j++ {
		if nums[j] < 2 {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
}
