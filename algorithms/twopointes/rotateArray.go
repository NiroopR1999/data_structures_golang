package twopointes

func rotate(nums []int, k int) {
	// Edge case: empty array.
	// We must check this before k % len(nums),
	// otherwise we would do k % 0 and panic.
	if len(nums) == 0 {
		return
	}

	// Rotating by k positions is the same as rotating
	// by k % len(nums) positions.
	//
	// Example:
	// [1,2,3,4,5], k = 7
	// 7 % 5 = 2
	// So we only need to rotate by 2.
	k = k % len(nums)

	// Reverse the entire array.
	//
	// Example:
	// [1,2,3,4,5]
	// becomes
	// [5,4,3,2,1]
	left, right := 0, len(nums)-1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]

		// Move both pointers toward the center
		// after swapping the current pair.
		left++
		right--
	}

	// Reverse the first k elements.
	//
	// Example after full reversal:
	// [5,4,3,2,1]
	//
	// k = 2
	// Reverse [5,4] -> [4,5]
	//
	// Result:
	// [4,5,3,2,1]
	left, right = 0, k-1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	// Reverse the remaining elements from index k to the end.
	//
	// [4,5 | 3,2,1]
	//         ↓
	// [4,5 | 1,2,3]
	//
	// Final:
	// [4,5,1,2,3]
	left, right = k, len(nums)-1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}