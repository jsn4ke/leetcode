package leetcode_159_lengthOfLongestSubstringTwoDistinct

func lengthOfLongestSubstringTwoDistinct(s string) int {
	var (
		left  = 0
		right = 0
		hash  = map[byte]int{}
		ans   int
	)
	for right < len(s) {
		c := s[right]
		hash[c]++
		right++
		for len(hash) > 2 {
			lc := s[left]
			if 1 == hash[lc] {
				delete(hash, lc)
			} else {
				hash[lc]--
			}
			left++
		}
		ans = max(ans, right-left)
	}
	return ans
}
