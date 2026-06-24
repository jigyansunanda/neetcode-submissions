func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
	if m == 0 || n == 0 {
		return m+n
	}

	cache := make([][]int, m+1)
	for i := range m+1 {
		cache[i] = make([]int, n+1)
		for j := range n+1 {
			if i == 0 || j == 0 {
				cache[i][j] = i+j
			} else if word1[i-1] == word2[j-1] {
				cache[i][j] = cache[i-1][j-1]
			} else {
				cache[i][j] = 1 + min(cache[i-1][j-1], cache[i-1][j], cache[i][j-1])
			}
		}
	}

	return cache[m][n]
}
