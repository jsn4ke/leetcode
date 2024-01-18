package leetcode_238_productExceptSelf

func productExceptSelf(nums []int) []int {
	var (
		length = len(nums)
		left   = make([]int, length)
		right  = make([]int, length)
		ret    []int
	)

	for i := range nums {
		if i == 0 {
			left[i] = 1
			right[length-1-i] = 1
		} else {
			left[i] = left[i-1] * nums[i-1]
			right[length-1-i] = right[length-i] * nums[length-i]
		}
	}
	for i := 0; i < length; i++ {
		ret = append(ret, left[i]*right[i])
	}
	return ret
}
