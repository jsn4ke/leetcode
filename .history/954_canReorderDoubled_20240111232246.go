package leatcode

import (
	"sort"
)

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)
	var (
		mp       = make(map[int]int, len(arr))
		leftIdx  = 0
		rightIdx = len(arr) - 1
	)
	for _, v := range arr {
		mp[v]++
	}

	for leftIdx < rightIdx {
		var (
			left  = arr[leftIdx]
			right = arr[rightIdx]
		)
		if val, ok := mp[left]; ok {
			double := left * 2
			re := mp[double] - val
			if re > 0 {
				mp[double] = re
			} else if re < 0 {
				return false
			} else {
				delete(mp, double)
			}
			delete(mp, left)
		}
		if val, ok := mp[right]; ok {
			if 1 == 1&right {
				return false
			}
			half := right / 2
		}
		left++
		right--
	}

}
