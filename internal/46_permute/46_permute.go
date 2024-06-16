package leetcode_46_permute

func permute1(nums []int) [][]int {
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
	return ret
}

func permute(nums []int) [][]int {
	var (
		ans [][]int
		dfs func(in []int, res ...int)
	)
	dfs = func(in []int, res ...int) {
		if 0 == len(in) {
			ans = append(ans, res)
			return
		}
		for i, v := range in {
			var arr []int
			arr = append(arr, in[:i]...)
			arr = append(arr, in[i+1:]...)
			var newRes = append([]int{}, res...)
			dfs(arr, append(newRes, v)...)
		}
	}
	dfs(nums)
	return ans
}
