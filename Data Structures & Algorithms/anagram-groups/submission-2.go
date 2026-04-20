func groupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{}
	for _, word := range strs {
		wordKey := [26]int{}
		for _, char := range word {
			wordKey[char -'a']++
		}
		groups[wordKey] = append(groups[wordKey], word)
	}
	format := [][]string{}
	for _, arr := range groups {
		format = append(format, arr)
	}
	return format
}
