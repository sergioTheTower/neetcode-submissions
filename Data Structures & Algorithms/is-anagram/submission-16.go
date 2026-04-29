func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := map[rune]int{}
	for _, char := range s {
		count[char]+=1
	}
	for _, char := range t {
		if val, ok := count[char]; !ok || val == 0 {
			return false
		}
		count[char]-=1
	}
return true
}
