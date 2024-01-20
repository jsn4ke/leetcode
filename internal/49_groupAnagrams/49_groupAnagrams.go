package leetcode_49_groupAnagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	var (
		mp  = map[string]int{}
		res [][]string
	)
	for _, v := range strs {
		v := v
		bt := []byte(v)
		sort.Slice(bt, func(i, j int) bool {
			return bt[i] < bt[j]
		})
		if idx, ok := mp[string(bt)]; ok {
			res[idx] = append(res[idx], v)
		} else {
			mp[string(bt)] = len(res)
			res = append(res, []string{v})
		}
	}
	return res
}
