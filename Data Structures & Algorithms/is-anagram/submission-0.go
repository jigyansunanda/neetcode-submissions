func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chars := make([]int, 26)
	for _, c := range s {
		chars[c-'a']++
	}
	for _, c := range t {
		chars[c-'a']--
	}
	for _, v := range chars {
		if v != 0 {
			return false
		}
	}
	return true
}