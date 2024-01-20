package leetcode_128_longestConsecutive

func longestConsecutive(nums []int) int {
	var (
		hash = map[int]struct{}{}
		ret  int
	)
	for _, v := range nums {
		hash[v] = struct{}{}
	}
	for k := range hash {
		// 比他还小的，不配
		if _, ok := hash[k-1]; ok {
			continue
		}
		next := k + 1
		for {
			if _, ok := hash[next]; !ok {
				break
			}
			next++
		}
		if next-k > ret {
			ret = next - k
		}
	}
	return ret
}
