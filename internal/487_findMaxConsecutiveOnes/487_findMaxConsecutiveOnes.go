package leetcode_487_findMaxConsecutiveOnes

func findMaxConsecutiveOnes(nums []int) int {
	var (
		input func(v int)
		ans   int
		left  = 0
		keep  = 0
	)
	input = func(v int) {
		keep++
		if v == 1 {
			ans = max(ans, left+keep)
		} else {
			ans = max(ans, keep)
			left = keep
			keep = 0
		}
	}
	for _, v := range nums {
		input(v)
	}
	return ans
}
