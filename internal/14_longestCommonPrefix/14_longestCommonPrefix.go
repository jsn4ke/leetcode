package leetcode_14_longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if 0 == len(strs) {
		return ""
	}
	if 1 == len(strs) {
		return strs[0]
	}
	var (
		maxLen = len(strs[0])
	)
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(strs[0]); j++ {
			if j >= len(strs[i]) || strs[0][j] != strs[i][j] {
				if j < maxLen {
					maxLen = j
				}
			}
		}
	}
	return strs[0][:maxLen]
}
