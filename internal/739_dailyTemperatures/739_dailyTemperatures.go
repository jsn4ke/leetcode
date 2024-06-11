package leetcode_739_dailyTemperatures

func dailyTemperatures(temperatures []int) []int {
	var (
		sk   [][2]int
		n    = len(temperatures)
		ans  = make([]int, n)
		last = temperatures[n-1]
	)
	sk = append(sk, [2]int{last, n - 1})
	ans[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		v := temperatures[i]
		for 0 != len(sk) && v >= sk[len(sk)-1][0] {
			sk = sk[:len(sk)-1]
		}
		if 0 == len(sk) {
			ans[i] = 0
		} else {
			ans[i] = sk[len(sk)-1][1] - i
		}
		sk = append(sk, [2]int{v, i})
	}
	return ans
}
