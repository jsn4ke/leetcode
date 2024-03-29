package leetcode_907_sumSubarrayMins

func sumSubarrayMins(arr []int) int {
	const mod = 1_000_000_000 + 7
	var ret int
	for i, v := range arr {
		var (
			right = i + 1
			left  = i - 1
		)
		for ; right < len(arr); right++ {
			if v > arr[right] {
				break
			}
		}
		for left >= 0 {
			if v > arr[left] {
				break
			}
			left--
		}
		ret = (ret + ((i - left) * (right - i) * v)) % mod
	}
	return ret
} // func sumSubarrayMins(arr []int) int {// 	const mod = 1_000_000_000 + 7// 	var (// 		dp  = make([]int, len(arr))// 		ret int// 	)// 	min := func(a, b int) int {// 		if a < b {// 			return a// 		}// 		return b// 	}// 	for i := range dp {// 		dp[i] = arr[i]// 		ret = (ret + arr[i]) % mod// 	}// 	for i := 1; i < len(arr); i++ {// 		for j := range dp {// 			if j+i >= len(arr) {// 				break// 			}// 			tmp := min(dp[j], arr[j+i])// 			dp[j] = tmp// 			ret = (ret + tmp) % mod// 		}// 	}// 	return ret// }// func sumSubarrayMins(arr []int) int {// 	const mod = 1_000_000_000 + 7// 	var (// 		traverse func(in []int, idx int, last int) int// 		ret      int// 	)// 	min := func(a, b int) int {// 		if a < b {// 			return a// 		}// 		return b// 	}// 	traverse = func(in []int, idx int, last int) int {// 		if idx >= len(in) {// 			return 0// 		}// 		// 第一层// 		if 0 == last {// 			last = in[idx]// 		} else {// 			last = min(last, in[idx])// 		}// 		return (last%mod + traverse(in, idx+1, last)%mod) % mod// 	}// 	for i := range arr {// 		ret = (ret%mod + traverse(arr, i, 0)%mod) % mod// 	}// 	return ret// }
