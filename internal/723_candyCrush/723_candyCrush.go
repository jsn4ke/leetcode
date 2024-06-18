package leetcode_723_candyCrush

func candyCrush(board [][]int) [][]int {
	var (
		m = len(board)
		n = len(board[0])
	)
	for {
		loop := false
		for i := 0; i < m; i++ {
			var (
				pre = 0
			)
			for j := 1; j < n; j++ {
				for j < n && 0 == board[i][j] {
					j++
				}
				if j == n {
					break
				}
				pre = j - 1
				for j < n && board[i][j] == board[i][pre] {
					j++
				}
				if j-pre >= 3 {
					loop = true
					for ; pre < j; pre++ {
						board[i][pre] = -board[i][pre]
					}
				}
				pre = j
			}
		}
		abs := func(i int) int {
			if i < 0 {
				return -i
			}
			return i
		}
		for j := 0; j < n; j++ {
			for i := 1; i < m; i++ {
				for i < m && 0 == board[i][j] {
					i++
				}
				if i == m {
					break
				}
				pre := i - 1
				for i < m && abs(board[i][j]) == abs(board[pre][j]) {
					i++
				}
				if i-pre >= 3 {
					loop = true
					for ; pre < i; pre++ {
						if board[pre][j] > 0 {
							board[pre][j] = -board[pre][j]
						}
					}
				}
				pre = i
			}
		}
		if !loop {
			break
		}
		for j := 0; j < n; j++ {
			var (
				start = 0
			)
			for i := m - 1; i >= 0; i-- {
				if board[i][j] < 0 {
					continue
				} else {
					board[m-1-start][j] = board[i][j]
					start++
				}
			}
			for ; start < m; start++ {
				board[m-1-start][j] = 0
			}
		}
	}
	return board
}
