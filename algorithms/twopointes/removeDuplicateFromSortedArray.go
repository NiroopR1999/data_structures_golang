package twopointes
func removeDuplicates(nums []int) int {
	// write points to the next position where we should
	// place a new unique element.
	//
	// We start at 1 because nums[0] is always unique.
	write := 1

	// read scans the array looking for new unique elements.
	read := 1

	// Edge case:
	// If nums is empty, write would be 1, which is incorrect.
	// So handle the empty array before starting.
	if len(nums) == 0 {
		return 0
	}

	for read < len(nums) {

		// The array is sorted, so duplicates are always adjacent.
		//
		// Therefore, we only need to compare the current element
		// with the previous element.
		//
		// Example:
		// [1,1,2,2,3]
		//     ↑ ↑
		//   read-1 read
		//
		// 1 == 1 → duplicate → skip
		//
		// When they are different, we found a new unique element.
		if nums[read] != nums[read-1] {

			// Copy the new unique value to the next write position.
			nums[write] = nums[read]

			// Move write because we have stored one more
			// unique element.
			write++
		}

		// read always moves because every element must be checked.
		read++
	}

	// write represents the number of unique elements,
	// not the last index.
	//
	// Example:
	// [1,2,3,2,3]
	// write = 3
	//
	// So the valid portion is nums[:3] → [1,2,3].
	return write
}