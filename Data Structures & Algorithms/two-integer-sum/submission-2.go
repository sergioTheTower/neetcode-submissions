func twoSum(nums []int, target int) []int {
	diff := map[int]int{}
	for idx, n := range nums {
		d := target - n
		if i, ok := diff[d]; ok{
			return []int{i,idx}
		}
		diff[n] = idx
	}
    return []int{}
}
