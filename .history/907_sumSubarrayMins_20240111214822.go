package leatcode

func sumSubarrayMins(arr []int) int {
	const mod = 1_000_000_000 + 7
	var (
		dp = make([][]int, len(arr))
	)
	for i := range dp {
		dp[i] = make([]int, len(arr))
		dp[i][i] = arr[i]
	}
	for i := 0; i < len(arr); i++ {
		for j := range dp {
			if j+i >= len(arr) {
				break
			}
			ret += dp[j][j+i]
		}
	}
}

// func sumSubarrayMins(arr []int) int {
// 	const mod = 1_000_000_000 + 7
// 	var (
// 		traverse func(in []int, idx int, last int) int
// 		ret      int
// 	)
// 	min := func(a, b int) int {
// 		if a < b {
// 			return a
// 		}
// 		return b
// 	}

// 	traverse = func(in []int, idx int, last int) int {
// 		if idx >= len(in) {
// 			return 0
// 		}
// 		// 第一层
// 		if 0 == last {
// 			last = in[idx]
// 		} else {
// 			last = min(last, in[idx])
// 		}
// 		return (last%mod + traverse(in, idx+1, last)%mod) % mod
// 	}
// 	for i := range arr {
// 		ret = (ret%mod + traverse(arr, i, 0)%mod) % mod
// 	}
// 	return ret
// }
