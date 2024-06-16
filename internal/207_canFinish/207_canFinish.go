package leetcode_207_canFinish

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges     = make([][]int, numCourses)
		inDegrees = make([]int, numCourses)
		stack     []int
		depth     int
		finish    int
	)
	for _, pre := range prerequisites {
		edges[pre[1]] = append(edges[pre[1]], pre[0])
		inDegrees[pre[0]]++
	}
	for i, v := range inDegrees {
		if 0 == v {
			stack = append(stack, i)
		}
	}
	for 0 != len(stack) {
		idx := stack[0]
		finish++
		stack = stack[1:]
		for _, next := range edges[idx] {
			inDegrees[next]--
			if 0 == inDegrees[next] {
				stack = append(stack, next)
			}
		}
		depth++
	}
	return finish == numCourses
}
