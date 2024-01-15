package leatcode

func convert(s string, numRows int) string {
	var (
		strs []string = make([]string, numRows)
	)

	for i, v := range s {
		v := v
		mod := 2*numRows - 2
		idx := i % mod
		if idx < numRows {
			strs[idx] += string(v)
		}
	}
}
