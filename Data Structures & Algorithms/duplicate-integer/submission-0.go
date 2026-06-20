func hasDuplicate(nums []int) bool {
	hashMap := map[int]int{}
	for _, num := range nums {
		_, ok := hashMap[num]
		if ok {
			return true
		}
		hashMap[num] = num
	}
	return false
}