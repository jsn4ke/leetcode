package leetcode_121_maxProfit

func maxProfit(prices []int) int {
	var (
		minVal = prices[0]
		ans    int
	)
	for _, v := range prices[1:] {
		ans = max(ans, v-minVal)
		minVal = min(minVal, v)
	}
	return ans
}
