package leetcode

// 待优化
func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}
	m, n := len(grid), len(grid[0])

	connected := make([][]bool, m)
	for i := range connected {
		connected[i] = make([]bool, n)
	}
	connected[row][col] = true
	changed := true

	for changed {
		changed = false
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == grid[row][col] && !connected[i][j] {
					if i-1 >= 0 && connected[i-1][j] {
						connected[i][j] = true
						changed = true
					}
					if i+1 < m && connected[i+1][j] {
						connected[i][j] = true
						changed = true
					}
					if j-1 >= 0 && connected[i][j-1] {
						connected[i][j] = true
						changed = true
					}
					if j+1 < n && connected[i][j+1] {
						connected[i][j] = true
						changed = true
					}
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !connected[i][j] {
				grid[i][j] = grid[i][j]
			} else {
				if i-1 < 0 || !connected[i-1][j] {
					grid[i][j] = color
				}
				if i+1 >= m || !connected[i+1][j] {
					grid[i][j] = color
				}
				if j-1 < 0 || !connected[i][j-1] {
					grid[i][j] = color
				}
				if j+1 >= n || !connected[i][j+1] {
					grid[i][j] = color
				}
			}
		}
	}
	return grid
}

// dfs
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func colorBorder2(grid [][]int, row, col, color int) [][]int {
	m, n := len(grid), len(grid[0])
	type point struct{ x, y int }
	var borders []point
	originalColor := grid[row][col]
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(int, int)
	dfs = func(x, y int) {
		vis[x][y] = true
		isBorder := false
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			if !(0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == originalColor) {
				isBorder = true
			} else if !vis[nx][ny] {
				vis[nx][ny] = true
				dfs(nx, ny)
			}
		}
		if isBorder {
			borders = append(borders, point{x, y})
		}
	}
	dfs(row, col)

	for _, p := range borders {
		grid[p.x][p.y] = color
	}
	return grid
}
