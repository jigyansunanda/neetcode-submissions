func robInRange(nums []int, l, r int) int {
	n := r - l + 1
	if n == 1 {
		return nums[l]
	}
	cache := make([]int, n)
	cache[0], cache[1] = nums[l], max(nums[l], nums[l+1])
	for i := l + 2; i <= r; i++ {
		cache[i-l] = max(cache[i-1-l], cache[i-2-l]+nums[i])
	}
	return max(cache[n-1], cache[n-2])
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(robInRange(nums, 0, n-2), robInRange(nums, 1, n-1))
}