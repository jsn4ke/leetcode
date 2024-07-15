package leetcode__top_interview_150_52_totalNQueens

import "fmt"

func totalNQueens(n int) int {
	var dfs func(idx int, colMask, left, right uint16)
	var ans int
	dfs = func(idx int, colMask, left, right uint16) {
		var combine = colMask | left | right
		fmt.Printf("idx %v combine %016b col %016b left %016b right %016b\n",
			idx, combine, colMask, left, right)
		if idx == n {
			ans++
			return
		}

		for i := uint16(0); i < uint16(n); i++ {
			mask := uint16(1 << i)
			if 0 != mask&combine {
				continue
			}
			dfs(idx+1, colMask|mask, (left|mask)<<1, (right|mask)>>1)
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println("====")
		mask := uint16(1 << i)
		dfs(1, mask, mask<<1, mask>>1)
	}
	return ans
}
