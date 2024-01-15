package leatcode

func sumSubarrayMins(arr []int) int {
	var (
		traverse func(in []int, idx int, min int) int
	)
	traverse = func(in []int, idx intm, min int) int {
		for i := idx; i < len(in); i++ {
			traverse(in, i+1)
		}
	}

	ret := traverse(arr, 0, 0)
	return ret
}
