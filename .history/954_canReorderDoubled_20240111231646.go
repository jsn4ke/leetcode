package leatcode

import (
	"sort"
)

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)
	var (
		mp = make(map[int]int, len(arr))
	)
	for _, v := range arr {
		mp[v]++
	}
}
