func twoSum(nums []int, target int) []int {
	count := map[int]int{}
	for idx, n := range nums {
		diff := target - n 
		val, ok := count[diff]
		if ok {
			return []int{val, idx}
		}
		count[n]=idx
	}
	return []int{}
}
