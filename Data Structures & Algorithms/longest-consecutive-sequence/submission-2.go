func longestConsecutive(nums []int) int {
    set := map[int]bool{}
	for _, v := range nums {
		set[v] = true
	}
	longest := 0
	for num := range set {
		// Check if there is a left value, if not it's a start of a seq.
		length := 0
		for set[num+length] {
			length++
		}
		if length > longest {
			longest = length
		}
	}
	return longest
}
