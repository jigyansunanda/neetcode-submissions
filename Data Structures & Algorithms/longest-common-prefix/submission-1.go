func longestCommonPrefix(strs []string) string {
	prefix, endIndex := strs[0], len(strs[0])-1

	if endIndex < 0 {
		return ""
	}
	for i := 1; i < len(strs); i++ {
		endIndex = min(endIndex, len(strs[i])-1)
		j := -1
		for j < endIndex {
			if prefix[j+1] == strs[i][j+1] {
				j++
			} else {
				break
			}
		}
		endIndex = j
	}
	if endIndex == -1 {
		return ""
	} else {
		return prefix[:endIndex+1]
	}
}