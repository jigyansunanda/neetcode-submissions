func tribonacci(n int) int {
	cache := make([]int, n+3)
	cache[0] = 0
	cache[1] = 1
	cache[2] = 1
	for i:=3; i<=n; i++ {
		cache[i] = cache[i-1] + cache[i-2] + cache[i-3]
	}
	return cache[n]
}
