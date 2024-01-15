package leatcode

import "sort"

func repeatLimitedString(s string, repeatLimit int) string {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] > bs[j]
	})
}
