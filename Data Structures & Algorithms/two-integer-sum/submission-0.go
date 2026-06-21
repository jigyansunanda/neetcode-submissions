func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, v := range nums {
		j, ok := hashMap[target-v]
		if ok {
			return []int{j, i}
		}
		hashMap[v] = i
	}
	return []int{}
}