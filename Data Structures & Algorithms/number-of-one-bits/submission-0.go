/*
10001011
01000101
----------

 */
func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		if (n & 1) > 0 {
			count++
		}
		n = n>>1
	}
	return count
}


