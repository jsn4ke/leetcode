package leetcode_1134_isArmstrong

func isArmstrong(n int) bool {
	var (
		p   = n
		arr []int
	)
	for 0 != p {
		part := p % 10
		arr = append(arr, part)
		p /= 10
	}
	calc := func(a, num int) (ans int) {
		ans = 1
		for i := 0; i < num; i++ {
			ans *= a
		}
		return
	}
	var ans int
	for _, v := range arr {
		ans += calc(v, len(arr))
	}
	return ans == n
}
