package twopointes

func twoSum(numbers []int, target int) []int {
	// left starts at the smallest value.
	// right starts at the largest value.
	//
	// Because the array is sorted:
	// - Moving left right  -> increases the sum
	// - Moving right left  -> decreases the sum
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			// LeetCode expects 1-based indices,
			// while Go uses 0-based indices.
			return []int{left + 1, right + 1}

		} else if sum > target {
			// Sum is too large.
			// Moving right left gives us a smaller number,
			// so the sum decreases.
			right--

		} else {
			// Sum is too small.
			// Moving left right gives us a larger number,
			// so the sum increases.
			left++
		}
	}

	// No pair adds up to target.
	return []int{}
}