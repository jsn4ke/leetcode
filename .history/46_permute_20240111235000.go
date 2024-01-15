package leatcode

func permute(nums []int) [][]int {
	var (
		traverse func(mask int)
	)

	traverse = func(mask int) {
		for i := range nums {
			if 1 << i & mask {

			}
		}
	}

	for i := range nums {
		var mask int
		traverse(mask | 1<<i)
	}
}
