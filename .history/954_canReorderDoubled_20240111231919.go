package leatcode

import (
	"sort"
)

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)
	var (
		mp    = make(map[int]int, len(arr))
		left  = 0
		right = len(arr) - 1
	)
	for _, v := range arr {
		mp[v]++
	}

	for left < right {
		if val, ok := mp[left]; ok {
			double = left * 2
			re := mp[left*2] - val
			if re > 0 {
				mp[left*2] = re
			} else if re < 0 {
				return false
			} else {
				delete(mp, left*2)
			}
		}
	}

}
