func isAnagram(s string, t string) bool {
if len(s) != len(t) {
    return false
}
countS := map[string]int{}
countT := map[string]int{}

for idx, str := range s {
countS[string(str)]++
countT[string(t[idx])]++
}

 for k, v := range countS {
        if countT[k] != v {
            return false
        }
    }
return true
}
