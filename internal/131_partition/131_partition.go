package leetcode_131_partition

func partition(s string) [][]string {
	var (
		n  = len(s)
		dp = make([][]int8, n)
	)
	for i := 0; i < n; i++ {
		dp[i] = make([]int8, n)
		dp[i][i] = 1
	}
	var is func(l, r int) int8
	is = func(l, r int) int8 {
		if l >= r {
			return 1
		}
		if 0 != dp[l][r] {
			return dp[l][r]
		}
		dp[l][r] = -1
		if s[l] == s[r] {
			dp[l][r] = is(l+1, r-1)
		}
		return dp[l][r]
	}
	var (
		ans [][]string
		tmp []string
	)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			ans = append(ans, append([]string{}, tmp...))
			return
		}
		for i := idx; i < n; i++ {
			if 1 == is(idx, i) {
				tmp = append(tmp, s[idx:i+1])
				dfs(i + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	dfs(0)
	return ans
}
