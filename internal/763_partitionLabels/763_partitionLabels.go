package leetcode_763_partitionLabels

func partitionLabels(s string) []int {
	var (
		hash = map[byte]int{}
		ans  []int
	)
	for i := range s {
		c := s[i]
		hash[c] = i
	}
	var (
		begin int
		last  int
	)
	for i := range s {
		c := s[i]
		last = max(last, hash[c])
		if last <= i {
			ans = append(ans, i-begin+1)
			begin = i + 1
			last = i + 1
		}
	}
	return ans
}
