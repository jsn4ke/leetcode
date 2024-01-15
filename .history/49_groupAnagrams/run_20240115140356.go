package leetcode_49_groupAnagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	var mp = map[string]int{}
	for _, v := range strs {
		bt := []byte(v)
		sort.Slice(bt, func(i, j int) bool {
			return bt[i] < bt[j]
		})
	}
}
