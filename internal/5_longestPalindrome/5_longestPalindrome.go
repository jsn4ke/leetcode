package leetcode_5_longestPalindrome

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		length := max(len1, len2)
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func longestPalindrome2(s string) string {
	var (
		size = len(s)
	)
	if size <= 1 {
		return s
	}
	var (
		dp = make([][]bool, size)
	)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, size)
		dp[i][i] = true
	}
	var (
		maxLen = 1
		start  = 0
	)
	for l := 2; l <= size; l++ {
		for i := 0; i < size; i++ {
			j := i + l - 1
			if j >= size {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if l == 2 || dp[i+1][j-1] {
					dp[i][j] = true
					if l > maxLen {
						maxLen = l
						start = i
					}
				}
			}
		}
	}
	return s[start : start+maxLen]
}
