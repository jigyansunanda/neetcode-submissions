var dirs = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func longestIncreasingPath(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])

	// intialize cache array
	cache := make([][]int, rows)
	for i := range cache {
		cache[i] = make([]int, cols)
		for j := range cols {
			cache[i][j] = -1
		}
	}

	maxPathLength := 1

	for i := range rows {
		for j := range cols {
			if cache[i][j] == -1 {
				currentPathLength := dfs(matrix, cache, i, j, rows, cols)
				maxPathLength = max(maxPathLength, currentPathLength)
			}
		}
	}

	return maxPathLength
}

func dfs(matrix [][]int, cache [][]int, x int, y int, rows int, cols int) int {
	if cache[x][y] != -1 {
		return cache[x][y]
	}
	cache[x][y] = 1
	for i := range 4 {
		r, c := x+dirs[i][0], y+dirs[i][1]
		if r >= 0 && c >= 0 && r < rows && c < cols {
			if matrix[r][c] > matrix[x][y] {
				cache[x][y] = max(cache[x][y], 1+dfs(matrix, cache, r, c, rows, cols))
			}
		}
	}
	return cache[x][y]
}