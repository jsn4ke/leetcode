package leetcode_136_singleNumber

func singleNumber(nums []int) int {
	var (
		ans = nums[0]
	)
	for _, v := range nums[1:] {
		ans ^= v
	}
	return ans
}
