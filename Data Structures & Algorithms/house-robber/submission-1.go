func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	cache := make([]int, n)
	cache[0], cache[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		cache[i] = max(cache[i-1], cache[i-2]+nums[i])
	}
	return max(cache[n-1], cache[n-2])
}