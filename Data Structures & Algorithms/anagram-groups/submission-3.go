func groupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{}
	for _, word := range strs {
		key := [26]int{}
		for _, r := range word {
			key[r-'a']++
		}
		groups[key] = append(groups[key], word)
	}
	format := [][]string{}
	for _, arr := range groups{
		format = append(format, arr)
	}
	return format
}
