package leetcode_139_wordBreak

func wordBreak(s string, wordDict []string) bool {
	var (
		hash           = make(map[string]bool)
		n              = len(s)
		small, big int = len(s), 0
		cache          = make(map[int]bool)
	)
	for _, v := range wordDict {
		hash[v] = true
		small = min(small, len(v))
		big = max(big, len(v))
	}
	var dfs func(start int) bool
	dfs = func(start int) bool {
		if start == n {
			return true
		}
		for size := start + small; size <= n && size <= start+big; size++ {
			if cache[size] {
				continue
			}
			if _, ok := hash[s[start:size]]; ok {
				if dfs(size) {
					return true
				}
				cache[size] = true
			}
		}
		return false
	}
	return dfs(0)
}
