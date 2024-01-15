package leatcode

import "sort"

func repeatLimitedString(s string, repeatLimit int) string {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] > bs[j]
	})
	var (
		dup      = 0
		startIdx = 0
	)
	for i := 0; i < len(bs); i++ {
		if bs[i] == bs[startIdx] {

			if i-startIdx > repeatLimit {
				for j := i; j < len(bs); j++ {
					if bs[i] == bs[startIdx] {

					}
				}
			}
		}
	}
}
