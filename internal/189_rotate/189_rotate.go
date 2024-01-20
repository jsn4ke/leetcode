package leetcode_189_rotate

func rotate(nums []int, k int) {
	var (
		length = len(nums)
	)
	for i := 0; i < length>>1; i++ {
		nums[i], nums[length-1-i] = nums[length-1-i], nums[i]
	}
	for i := 0; i < k>>1; i++ {
		nums[i], nums[k-i-1] = nums[k-i-1], nums[i]
	}
	for i := 0; i < (length-k)>>1; i++ {
		nums[k+i], nums[length-i-1] = nums[length-i-1], nums[k+i]
	}
}
