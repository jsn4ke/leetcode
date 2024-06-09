package leetcode_283_moveZeroes

func moveZeroes2(nums []int) {
	for i := 0; i < len(nums); i++ {
		if 0 == nums[i] {
			for j := i + 1; j < len(nums); j++ {
				if 0 != nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
				}
			}
		}
	}
}

func moveZeroes(nums []int) {
	var (
		idx int
	)
	for _, v := range nums {
		if 0 != v {
			nums[idx] = v
			idx++
		}
	}
	for i := idx; i < len(nums); i++ {
		nums[i] = 0
	}
}
