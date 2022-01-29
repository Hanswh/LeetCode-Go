package leetcode

type pair struct {
	x int
	y int
}

var dirs = []pair{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

// bfs
func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = -1
		}
	}

	var queue []pair
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				res[i][j] = 0
				queue = append(queue, pair{i, j})
			}
		}
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, dir := range dirs {
			if x, y := p.x+dir.x, p.y+dir.y; x >= 0 && x < m && y >= 0 && y < n && res[x][y] == -1 {
				res[x][y] = res[p.x][p.y] + 1
				queue = append(queue, pair{x, y})
			}
		}
	}

	return res
}
