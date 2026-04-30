func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	for _, n := range nums {
		count[n]+=1
	}
	// fmt.Println(count)
	freq := make([][]int, len(nums)+1)
	for k, v := range count {
		// fmt.Println("adding value to freq")
		freq[v] = append(freq[v], k)
	}
	// fmt.Println("Done....")
	result := []int{}
	for idx, _ := range freq {
		fmt.Println(idx)
		end := len(freq) - 1 - idx
		// fmt.Println(freq[end])
		for _, num := range freq[end] {
			result = append(result, num)
				if len(result) == k {
				return result
			}
		}
	}
return result
}
