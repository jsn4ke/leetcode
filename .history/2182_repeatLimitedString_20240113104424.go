package leatcode

import "sort"

func repeatLimitedString(s string, repeatLimit int) string {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] > bs[j]
	})
	var (
		startIdx = 0
	)
	for i := 0; i < len(bs); i++ {
		if bs[i] != bs[startIdx] {
			startIdx = i
			continue
		}
		if i-startIdx > repeatLimit {
			for j := i + 1; j < len(bs); j++ {
				if bs[i] != bs[j] {
					bs[i], bs[j] = bs[j], bs[i]
					break
				}
			}
			startIdx = i
		}
	}
	return string(bs)
}
