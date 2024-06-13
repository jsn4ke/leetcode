package leetcode_994_orangesRotting

func orangesRotting(grid [][]int) int {
	var (
		m, n  = len(grid), len(grid[0])
		fresh = map[[2]int]struct{}{}
		stack [][2]int
		ans   int
	)
	for i, arr := range grid {
		for j, v := range arr {
			switch v {
			case 1:
				fresh[[2]int{i, j}] = struct{}{}
			case 2:
				stack = append(stack, [2]int{i, j})
			}
		}
	}
	ways := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	if 0 == len(fresh) {
		return 0
	}
	for 0 != len(fresh) && 0 != len(stack) {
		var tmp [][2]int
		for _, v := range stack {
			for _, way := range ways {
				x, y := v[0]+way[0], v[1]+way[1]
				if x < 0 || x >= m || y < 0 || y >= n {
					continue
				}
				key := [2]int{x, y}
				if _, ok := fresh[key]; ok {
					delete(fresh, key)
					tmp = append(tmp, key)
				}
			}
		}
		stack = tmp
		ans++
	}
	if 0 == len(fresh) {
		return ans
	}
	return -1
}
