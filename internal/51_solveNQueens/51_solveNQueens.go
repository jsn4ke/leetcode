package leetcode_51_solveNQueens

func solveNQueens(n int) [][]string {
	var ans [][]string
	var (
		arrs []string
	)
	for i := 0; i < n; i++ {
		var str = make([]byte, n)
		for j := 0; j < n; j++ {
			str[j] = '.'
		}
		str[i] = 'Q'
		arrs = append(arrs, string(str))
	}
	var checkValid func(row int, in []int) bool
	checkValid = func(row int, in []int) bool {
		var (
			top    = 0
			botton = len(in)
			left   = 0
			right  = n
		)
		for i := 0; i < len(in); i++ {
			x, y := i, in[i]
			for step := top; step < botton; step++ {
				if step == x {
					continue
				}
				if in[step] == y {
					return false
				}
			}
			// 右下
			for i := 1; i < botton; i++ {
				sx := x + i
				sy := y + i
				if sx >= botton || sy >= right {
					break
				}
				if in[sx] == sy {
					return false
				}
			}
			// 右上
			for i := 1; i < botton; i++ {
				sx := x - i
				sy := y + i
				if sx < top || sy >= right {
					break
				}
				if in[sx] == sy {
					return false
				}
			}
			// 左下
			for i := 1; i < botton; i++ {
				sx := x + i
				sy := y - i
				if sx >= botton || sy < left {
					break
				}
				if in[sx] == sy {
					return false
				}
			}
			// 左上
			for i := 1; i < botton; i++ {
				sx := x - i
				sy := y - i
				if sx < top || sy < left {
					break
				}
				if in[sx] == sy {
					return false
				}
			}
		}
		return true
	}
	var dfs func(row int, before ...int)
	dfs = func(row int, before ...int) {
		if row == n-1 {
			var tmp []string
			for _, v := range before {
				tmp = append(tmp, arrs[v])
			}
			ans = append(ans, tmp)
			return
		}
		for j := 0; j < n; j++ {
			tmp := append(before, j)
			if checkValid(row+1, tmp) {
				dfs(row+1, tmp...)
			}
		}
	}
	for i := 0; i < n; i++ {
		dfs(0, i)
	}
	return ans
}
