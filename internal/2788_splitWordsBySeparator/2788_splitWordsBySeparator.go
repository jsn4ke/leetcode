package leetcode_2788_splitWordsBySeparator

func splitWordsBySeparator(words []string, separator byte) []string {
	split := func(s string) []string {
		var (
			ret  []string
			last []byte
		)
		for i := 0; i < len(s); i++ {
			c := s[i]
			if c != separator {
				last = append(last, c)
			} else {
				if 0 != len(last) {
					ret = append(ret, string(last))
					last = nil
				}
			}
		}
		if 0 != len(last) {
			ret = append(ret, string(last))
		}
		return ret
	}
	var (
		ret []string
	)
	for _, s := range words {
		ret = append(ret, split(s)...)
	}
	return ret
}
