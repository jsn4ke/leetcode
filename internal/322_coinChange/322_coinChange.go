package leetcode_322_coinChange

import "sort"

func coinChange(coins []int, amount int) int {
	if 0 == amount {
		return 0
	}
	sort.Ints(coins)
	var dp = make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	for i := coins[0]; i <= amount; i++ {
		for _, v := range coins {
			if v > i {
				break
			}
			if i == v {
				dp[v] = 1
			} else {
				if -1 == dp[i-v] {
					continue
				} else {
					if -1 == dp[i] {
						dp[i] = dp[i-v] + 1
					} else {
						dp[i] = min(dp[i], dp[i-v]+1)
					}
				}

			}
		}
	}
	return dp[amount]
}
