package leetcode_394_decodeString

func decodeString(s string) string {
	// 2[abs3[sxx43[fe]45[xx]]]er
	// 2[abc]erf

	var h1 func(startIdx int) (sub string, endIdx int)
	h1 = func(startIdx int) (sub string, endIdx int) {
		at := startIdx
		var record string
		var prefix int
		at++
		for at < len(s) {
			c := s[at]
			switch c {
			case ']':
				return record, at
			case '[':
				tmpSub, tmpEndIdx := h1(at)
				for i := 0; i < prefix; i++ {
					record += tmpSub
				}
				prefix = 0
				at = tmpEndIdx + 1
			default:
				if c >= '0' && c <= '9' {
					prefix = prefix*10 + int(c-'0')
				} else {
					record += string(c)
				}
				at++
			}
		}
		return record, at
	}
	ans, _ := h1(0)
	return ans
}
