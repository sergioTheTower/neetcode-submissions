func groupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{}
	for _, word := range strs {
		key := [26]int{}
		for _, char := range word {
			key[char-'a']++
		}
		groups[key] = append(groups[key], word)
	}
	result := [][]string{}
	for _, value := range groups {
		result = append(result, value)
	}
return result
}
