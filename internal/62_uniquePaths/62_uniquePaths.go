package leetcode_62_uniquePaths

var cache = map[[2]int]int{}

func uniquePaths(m int, n int) int {
	if 1 == m && 1 == n {
		return 1
	}
	key := [2]int{m, n}
	if v, ok := cache[key]; ok {
		return v
	}
	var ans int
	if m > 1 {
		ans += uniquePaths(m-1, n)
	}
	if n > 1 {
		ans += uniquePaths(m, n-1)
	}
	cache[key] = ans
	return ans
}
