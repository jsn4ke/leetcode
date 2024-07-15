package leetcode__top_interview_150_88_merge

func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		idx  = 0
		p, q = m - 1, n - 1
	)
	for p >= 0 && q >= 0 {
		if nums1[p] < nums2[q] {
			nums1[m+n-1-idx] = nums2[q]
			q--
		} else {
			nums1[m+n-1-idx] = nums1[p]
			p--
		}
		idx++
	}
	for p >= 0 {
		nums1[m+n-1-idx] = nums1[p]
		p--
		idx++
	}
	for q >= 0 {
		nums1[m+n-1-idx] = nums2[q]
		q--
		idx++
	}
}
