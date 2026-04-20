func hasDuplicate(nums []int) bool {
    set := map[int]bool{}
    for _,  v := range nums {
        if _, ok := set[v]; ok {
            return true
        }
        set[v] = true
    }
    return false
}
