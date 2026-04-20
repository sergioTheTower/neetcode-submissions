func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	result := 0
	l, r := 0, len(height)-1
	lMax, rMax := height[l], height[r]
	for l < r {
		if lMax < rMax {
			l += 1
			lMax = max(lMax, height[l])
			result += lMax - height[l]
			continue
		}
		r -= 1
		rMax = max(rMax, height[r])
		result += rMax - height[r]

	}
	return result
}
