func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    count := map[rune]int{
    }
    for i, r := range s {
        count[r] +=1
        count[rune(t[i])] -=1
    }
    // for _, r := range t {
    //     count[r] -=1
    // }
    for _, c := range count {
        if c != 0  {
            return false
        }
    }
    return true
}
