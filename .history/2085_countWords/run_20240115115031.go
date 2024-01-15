package leetcode_2085_countWords

func countWords(words1 []string, words2 []string) int {
	var (
		key1 = map[string]int{}
		key2 = map[string]int{}
		res  int
	)
	for _, v := range words1 {
		key1[v]++
	}
	for _, v := range words2 {
		key2[v]++
	}
	for k, v := range key1 {
		if 1 != v {
			continue
		}
		if num, ok := key2[k]; !ok || 1 != num {
			continue
		}
		res++
	}
	return res
}
