package leetcode_323_countComponents

func countComponents(n int, edges [][]int) int {
	var (
		arr    = make([]int, n)
		parent func(i int) int
		union  func(i, j int)
	)
	parent = func(i int) int {
		if i == arr[i] {
			return i
		}
		arr[i] = parent(arr[i])
		return arr[i]
	}
	union = func(i, j int) {
		pi, pj := parent(i), parent(j)
		if pi == pj {
			return
		}
		arr[pi] = pj
	}
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	for _, edge := range edges {
		union(edge[0], edge[1])
	}
	var ans int
	for i, v := range arr {
		if i == v {
			ans++
		}
	}
	return ans
}
