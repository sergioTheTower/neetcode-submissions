func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	for _, n := range nums {
		count[n]+=1
	}
	
	freq := make([][]int, len(nums)+1)
	for key, v := range count {
		freq[v] = append(freq[v], key)
	}
	result := []int{}
	for idx, _ := range freq {
		lastIdx := len(freq)-idx -1
		for _, num := range freq[lastIdx] {
		result = append(result, num)
		if len(result) == k {
			return result
		}
		}
	}
	return result
}
