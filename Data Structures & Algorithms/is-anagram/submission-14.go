func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    count := map[rune]int{
    }
    for _, r := range s {
        count[r] +=1
    }
    for _, r := range t {
       if val, ok := count[r]; val == 0  || !ok {
			return false
		}
        count[r] -=1
    }
    return true
}
