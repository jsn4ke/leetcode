package leetcode_200_numberOfIslands

type UnionFind struct {
	count  int
	parent []int
}

func (uf *UnionFind) init(grid [][]byte) {
	uf.count = 0
	uf.parent = make([]int, len(grid)*len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				uf.parent[i*len(grid[0])+j] = i*len(grid[0]) + j
				uf.count++
			}
		}
	}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX != rootY {
		uf.parent[rootX] = rootY
		uf.count--
	}
}

func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	uf := UnionFind{}
	uf.init(grid)
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				for _, direction := range directions {
					newX := i + direction[0]
					newY := j + direction[1]
					if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == '1' {
						uf.union(i*len(grid[0])+j, newX*len(grid[0])+newY)
					}
				}
			}
		}
	}
	return uf.count
}

func numIslands(grid [][]byte) int {
	var (
		m = len(grid)
		n = len(grid[0])
	)
	var outBound = func(x, y int) bool {
		return x < 0 || x >= m || y < 0 || y >= n
	}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if outBound(x, y) || grid[x][y] != '1' {
			return 0
		}
		grid[x][y] = '2'
		return 1 + dfs(x+1, y) + dfs(x-1, y) + dfs(x, y+1) + dfs(x, y-1)
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, dfs(i, j))
		}
	}
	return res
}
