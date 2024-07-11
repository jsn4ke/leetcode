package leetcode_151_reverseWords

func reverseWords(s string) string {
	var (
		in = []byte(s)
	)
	reverse := func(l, r int) {
		r--
		for l < r {
			in[l], in[r] = in[r], in[l]
			l++
			r--
		}
	}
	reverse(0, len(in))
	var idx = 0
	for i := 0; i < len(in); i++ {
		if in[i] != ' ' {
			if 0 != idx {
				in[idx] = ' '
				idx++
			}
			j := i
			for ; j < len(in) && in[j] != ' '; j++ {
				in[idx] = in[j]
				idx++
			}
			reverse(idx-(j-i), idx)
			i = j
		}
	}
	in = in[:idx]
	return string(in)
}
