package leatcode

func permute(nums []int) [][]int {
	var (
		traverse func(mask int)
		arr      []int
	)

	for i := range nums {
		var mask int
		traverse(mask | 1<<i)
	}
}
