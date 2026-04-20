
func isAnagram(s string, t string) bool {
	setS := map[string]int{}
	for _, v := range s {
		setS[string(v)]+=1
	}
	for _, v := range t {
		setS[string(v)]-=1
	}
	for _, v := range setS{
		if v != 0 {
			return false
		}
	}
	return true
}
