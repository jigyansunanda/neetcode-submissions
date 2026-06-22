func groupAnagrams(strs []string) [][]string {
	hashMap := map[[26]int][]string{}
	for _, str := range strs {
		var hashKey [26]int
		for _, c := range str {
			hashKey[c-'a']++
		}
		hashMap[hashKey] = append(hashMap[hashKey], str)
	}
	anagramGroups := [][]string{}
	for _, group := range hashMap {
		anagramGroups = append(anagramGroups, group)
	}
	return anagramGroups
}