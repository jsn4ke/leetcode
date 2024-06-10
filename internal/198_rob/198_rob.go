package leetcode_198_rob

func rob2(nums []int) int {
	var (
		n = len(nums)
	)
	if 1 == n {
		return nums[0]
	}
	if 2 == n {
		return max(nums[0], nums[1])
	}
	var (
		left  = make([]int, n)
		right = make([]int, n)
	)
	left[0] = nums[0]
	left[1] = max(nums[0], nums[1])

	right[n-1] = nums[n-1]
	right[n-2] = max(nums[n-1], nums[n-2])

	for i := 2; i < n; i++ {
		left[i] = -1
		right[n-1-i] = -1
	}

	var getLeft func(int) int
	getLeft = func(i int) int {
		if i < 0 {
			return 0
		}
		if -1 != left[i] {
			return left[i]
		}
		left[i] = max(getLeft(i-1), getLeft(i-2)+nums[i])
		return left[i]
	}
	var getRight func(int) int
	getRight = func(i int) int {
		if i >= n {
			return 0
		}
		if -1 != right[i] {
			return right[i]
		}
		right[i] = max(getRight(i+1), getRight(i+2)+nums[i])
		return right[i]
	}
	var ans int
	// left[i-1] + right[i+1]
	// left[i-2] + nums[i] + right[i+2]
	for i := 0; i < n; i++ {
		a := getLeft(i-2) + nums[i] + getRight(i+2)
		b := getLeft(i-1) + getRight(i+1)
		ans = max(ans, max(a, b))
	}
	return ans
}

func rob(nums []int) int {
	var (
		n = len(nums)
	)
	if 1 == n {
		return nums[0]
	}
	if 2 == n {
		return max(nums[0], nums[1])
	}
	k2, k1 := nums[0], max(nums[0], nums[1])
	var ans = k1
	for i := 2; i < len(nums); i++ {
		a := max(k2+nums[i], k1)
		k2, k1 = k1, a
		ans = max(ans, k1)
	}
	return ans
}
