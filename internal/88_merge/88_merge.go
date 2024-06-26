package leetcode_88_merge

func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		i = m - 1
		j = n - 1
		k = m + n - 1
	)
	for k >= 0 {
		if i < 0 {
			nums1[k] = nums2[j]
			j--
		} else if j < 0 {
			nums1[k] = nums1[i]
			i--
		} else if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}
