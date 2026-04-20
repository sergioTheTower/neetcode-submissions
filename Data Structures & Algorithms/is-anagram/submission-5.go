// func isAnagram(s string, t string) bool {
// if len(s) != len(t) {
//     return false
// }
// countS := map[string]int{}
// countT := map[string]int{}

// for idx, str := range s {
// countS[string(str)]++
// countT[string(t[idx])]++
// }

//  for k, v := range countS {
//         if countT[k] != v {
//             return false
//         }
//     }
// return true
// }


func isAnagram(s string, t string) bool {
	// Anagrams must have the same length. This is a crucial early exit check.
	if len(s) != len(t) {
		return false
	}

	// 1. Create a single map to track character counts (using rune for full Unicode support).
	counts := make(map[rune]int)

	// 2. Increment counts for string s and decrement counts for string t.
	// We use two separate loops to safely handle potential multi-byte characters (runes)
	// that could cause misalignment if we relied on index iteration.
	for _, char := range s {
		counts[char]++
	}

	for _, char := range t {
		counts[char]--
		
		// Optimization: If the count drops below zero at any point, 
		// it means 't' has more of a certain character than 's', so they cannot be anagrams.
		if counts[char] < 0 {
			return false
		}
	}

	// 3. Implicit Check: Since we verified the lengths are equal, 
	// if no count ever dropped below zero, all counts must be exactly zero.
	// Therefore, no final loop is needed, and we can return true.
	return true
}