package convert_6

func convert(s string, numRows int) string {
	var (
		strs []string = make([]string, numRows)
	)

	for i, v := range s {
		v := v
		mod := 2*numRows - 2
		idx := i % mod
		if idx >= numRows {
			idx = mod - idx
		}
		strs[idx] += string(v)
	}
	var ret string
	for _, v := range strs {
		ret += v
	}
	return ret
}
