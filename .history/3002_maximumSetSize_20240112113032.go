package leatcode

func maximumSetSize(nums1 []int, nums2 []int) int {
	var (
		mp   = map[int]int{}
		pair = [][2]int{}
	)
	for _, v := range nums1 {
		mp[v]++
	}
	for _, v := range nums2 {
		mp[v]++
	}

}
