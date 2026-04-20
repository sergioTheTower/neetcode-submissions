func twoSum(nums []int, target int) []int {
    track := map[int]int{}

    for idx, num := range nums {
        diff := target - num
        if val, ok := track[diff]; ok {
            return []int{val, idx}
        }
        track[num] = idx
    }
    return []int{}  
}
