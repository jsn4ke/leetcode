package leetcode_2813_findMaximumElegance

import "sort"

func findMaximumElegance(items [][]int, k int) int64 {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	})
	var (
		stack  []int
		hash   = map[int]bool{}
		ans    int64
		profit int64
	)
	for i := 0; i < len(items); i++ {
		pi, cat := items[i][0], items[i][1]
		if i < k {
			profit += int64(pi)
			if hash[cat] {
				stack = append(stack, pi)
			} else {
				hash[cat] = true
			}
		} else {
			size := len(stack)
			if 0 == size {
				return ans
			}
			if !hash[cat] {
				profit -= int64(stack[size-1] - int(pi))
				stack = stack[:size-1]
				hash[cat] = true
			}
		}
		ans = max(ans, profit+int64(len(hash)*len(hash)))
	}
	return ans
}

func findMaximumElegance2(items [][]int, k int) int64 {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	})
	categorySet := map[int]bool{}
	var res, profit int64
	var st []int
	for i, item := range items {
		if i < k {
			profit += int64(item[0])
			if categorySet[item[1]] {
				st = append(st, item[0])
			} else {
				categorySet[item[1]] = true
			}
		} else if !categorySet[item[1]] && len(st) > 0 {
			profit += int64(item[0] - st[len(st)-1])
			st = st[:len(st)-1]
			categorySet[item[1]] = true
		}
		res = max(res, profit+int64(len(categorySet)*len(categorySet)))
	}
	return res
}
