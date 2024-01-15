package leatcode

func countWords(words1 []string, words2 []string) int {
	var (
		key1 = map[string]int{}
		key2 = map[string]int{}
	)
	for _, v := range words1 {
		key1[v]++
	}
}
