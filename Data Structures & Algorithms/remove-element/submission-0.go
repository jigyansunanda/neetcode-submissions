func removeElement(nums []int, val int) int {
	indexReached := -1

	for i := range nums {
		if nums[i] != val {
			indexReached++
			nums[indexReached] = nums[i]
		}
	}

	return indexReached + 1
}