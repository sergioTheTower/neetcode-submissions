func twoSum(nums []int, target int) []int {
    hash := map[int]int{}
    for idx, v := range nums {
        diff := target - v
        if val, ok := hash[diff]; ok {
            return []int{val, idx}
        }
        hash[v]= idx
    }
    return []int{}
}
