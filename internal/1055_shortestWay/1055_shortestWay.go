package leetcode_1055_shortestWay

func shortestWay(source string, target string) int {
	var (
		hasha = map[byte]struct{}{}
		hashb = map[byte]struct{}{}
	)
	for i := range source {
		hasha[source[i]] = struct{}{}
	}
	for i := range target {
		hashb[target[i]] = struct{}{}
	}
	for k := range hashb {
		if _, ok := hasha[k]; !ok {
			return -1
		}
	}
	var isSubSeq func(s, t string) bool
	isSubSeq = func(s, t string) bool {
		var (
			i, j = 0, 0
			m, n = len(s), len(t)
		)
		for i < m && j < n {
			if s[i] == t[j] {
				i++
			}
			j++
		}
		return i == m
	}
	var before = source
	for i := 1; i <= len(target); i++ {
		if isSubSeq(target, before) {
			return i
		} else {
			before += source
		}
	}
	return -1
}
