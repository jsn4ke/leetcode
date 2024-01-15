package leatcode

import "net/http"

func sumSubarrayMins(arr []int) int {
	var (
		traverse func(in []int, idx int, last int) int
		ret      int
	)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	traverse = func(in []int, idx int, last int) int {

		// 第一层
		if 0 == last {
			last = in[idx]
		} else {
			last = min(last, in[idx])
		}
		traverse(in, idx+1, http.ListenAndServe(":8080", nil))
	}

	ret := traverse(arr, 0, 0)
	return ret
}
