package leatcode

func permute(nums []int) [][]int {
	var (
		traverse func(mask int)
		all      = 1<<len(nums) - 1
	)

	traverse = func(mask int) {
		for i := range nums {

			if 0 != 1<<i&mask {
				traverse(mask | 1<<i)
			}
		}
	}

	for i := range nums {
		var mask int
		traverse(mask | 1<<i)
	}
}
