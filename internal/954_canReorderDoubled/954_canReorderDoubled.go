package leetcode_954_canReorderDoubled

import (
	"sort"
)

func canReorderDoubled(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i] < 0 && arr[j] < 0 {
			return arr[i] > arr[j]
		}
		if arr[i] > 0 && arr[j] > 0 {
			return arr[i] < arr[j]
		}
		return arr[i] < arr[j]
	})
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
			re := mp[half] - val
			if re > 0 {
				mp[half] = re
			} else if re < 0 {
				return false
			} else {
				delete(mp, half)
			}
		}
		left++
		right--
	}
	return 0 == len(mp)
}
