package leetcode_169_majorityElement

func majorityElement(nums []int) int {
	var (
		ans   int
		count int
	)
	for _, v := range nums {
		if v == ans {
			count++
		} else {
			if 0 == count {
				count = 1
				ans = v
			} else {
				count--
			}
		}
	}
	return ans
}
