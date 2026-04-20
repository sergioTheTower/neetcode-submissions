func twoSum(numbers []int, target int) []int {
	leftPointer, rightPointer := 0, len(numbers)-1
	for leftPointer < rightPointer {
		leftValue := numbers[leftPointer]
		rightValue := numbers[rightPointer]
		if leftValue+rightValue == target {
			return []int{leftPointer+1, rightPointer+1}
		}
		if leftValue+rightValue > target {
			rightPointer--
			continue
		}
		leftPointer++
	}
	return []int{}
}
