package leetcode_49_groupAnagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	for _, v := range strs {
		bt := []byte(v)
		sort.Slice(bt, func(i, j int) bool {
			bt[i] < bt[j]
		})
	}
}
