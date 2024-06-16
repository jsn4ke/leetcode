package leetcode_100316_maximumTotalDamage

import "sort"

// 一个魔法师有许多不同的咒语。

// 给你一个数组 power ，其中每个元素表示一个咒语的伤害值，可能会有多个咒语有相同的伤害值。

// 已知魔法师使用伤害值为 power[i] 的咒语时，他们就 不能 使用伤害为 power[i] - 2 ，power[i] - 1 ，power[i] + 1 或者 power[i] + 2 的咒语。

// 每个咒语最多只能被使用 一次 。

// 请你返回这个魔法师可以达到的伤害值之和的 最大值 。
func maximumTotalDamage(power []int) int64 {
	freq := make(map[int]int)
	for _, p := range power {
		freq[p]++
	}

	uniquePowers := make([]int, 0, len(freq))
	for p := range freq {
		uniquePowers = append(uniquePowers, p)
	}
	sort.Ints(uniquePowers)
	dp := make(map[int]int64)
	ans := int64(0)
	var last = -1
	for _, p := range uniquePowers {
		if -1 != last {
			for i := p - 3; i < p; i++ {
				if 0 == dp[i] {
					dp[i] = dp[last]
				}
			}
		}
		// 若决定使用当前伤害值p
		var useP int64
		if p > 2 { // 保证p-3是有效的
			useP = dp[p-3] + int64(p)*int64(freq[p])
		} else {
			useP = int64(p) * int64(freq[p])
		}
		// 若不使用p，即延续前一个状态
		dp[p] = max(max(dp[p-1], dp[p-2]), useP)
		dp[p+1] = max(dp[p+1], dp[p])
		dp[p+2] = max(dp[p+2], dp[p])
		ans = max(ans, dp[p])
		last = p
	}

	return ans
}
