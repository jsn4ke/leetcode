package leatcode

func permute(nums []int) [][]int {
	var (
		traverse func(mask int, in ...int)
		all      = 1<<len(nums) - 1
		ret      [][]int
	)

	traverse = func(mask int, in ...int) {
		if mask == all {
			ret = append(ret, in)
			return
		}
		for i := range nums {
			if 0 != 1<<i&mask {
				traverse(mask|1<<i, append(in, nums[i])...)
			}
		}
	}

	traverse(0)
}
