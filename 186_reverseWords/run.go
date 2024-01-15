package leetcode_186_reverseWords

func reverseWords(s []byte) {
	reverse := func(s []byte) {
		var (
			left  = 0
			right = len(s) - 1
		)
		for left < right {

			s[left], s[right] = s[right], s[left]
			left++
			right--
		}
	}
	reverse(s)
	var (
		lastIdx = 0
		curIdx  = 0
	)
	for ; curIdx < len(s); curIdx++ {
		if s[curIdx] == ' ' {
			reverse(s[lastIdx:curIdx])
			lastIdx = curIdx + 1
		}
	}
	reverse(s[lastIdx:curIdx])
}
