package leetcode_340_lengthOfLongestSubstringKDistinct

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	var (
		hash  = map[byte]int{}
		left  = 0
		right = 0
		ans   int
	)
	if 0 == k {
		return 0
	}
	for right < len(s) {
		char := s[right]
		if _, ok := hash[char]; !ok {
			hash[char]++
			if len(hash) > k {
				for left < right {
					c := s[left]
					left++
					if 1 == hash[c] {
						delete(hash, c)
						break
					} else {
						hash[c]--
					}
				}
			}
		} else {
			hash[char]++
		}
		right++
		ans = max(ans, right-left)
	}
	return ans
}
