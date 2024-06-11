package leetcode_279_numSquares

func numSquares(n int) int {
	var (
		coins []int
		dp    = make([]int, n+1)
	)
	for i := 1; i <= 100; i++ {
		sq := i * i
		if sq < n {
			dp[sq] = 1
			coins = append(coins, sq)
		} else if sq == n {
			return 1
		} else {
			break
		}
	}
	// dp[钱数] = dp[钱数] + coins
	for i := 1; i <= n; i++ {
		if dp[i] == 1 {
			continue
		}
		for _, coin := range coins {
			diff := i - coin
			if diff < 0 {
				break
			}
			if 0 == dp[i] {
				dp[i] = dp[diff] + 1
			} else {
				dp[i] = min(dp[diff]+1, dp[i])
			}
		}
	}
	return dp[n]
}
