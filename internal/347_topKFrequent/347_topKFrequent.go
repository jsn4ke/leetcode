package leetcode_347_topKFrequent

func topKFrequent(nums []int, k int) []int {
	var (
		hash = map[int]int{}
	)
	for _, v := range nums {
		hash[v]++
	}
	for len(hash) != k {
		for k, v := range hash {
			if 1 == v {
				delete(hash, k)
			} else {
				hash[k]--
			}
		}
	}
	var (
		res []int
	)
	for k := range hash {
		res = append(res, k)
	}
	return res
}
