package leetcode_100301_countCompleteDayPairs

func countCompleteDayPairs(hours []int) int64 {
	var (
		preSum = map[int]int{}
		ans    int
	)
	preSum[hours[0]%24] = 1
	for _, v := range hours[1:] {
		hash := v % 24
		if 0 != hash {
			hash = 24 - hash
		}
		if cnt, ok := preSum[hash]; ok {
			ans += cnt
		}
		preSum[v]++
	}
	return int64(ans)
}
