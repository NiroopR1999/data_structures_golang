package twopointes

func merge(nums1 []int, m int, nums2 []int, n int) {
	// i points to the last actual element in nums1.
	// Example: nums1 = [1,2,3,0,0,0], m = 3 -> i = 2
	i := m - 1

	// j points to the last element in nums2.
	// Example: nums2 = [2,5,6], n = 3 -> j = 2
	j := n - 1

	// k points to the last available position in nums1.
	// We fill from right to left because the empty positions
	// are at the end, so we avoid overwriting unprocessed values.
	k := m + n - 1

	// We only need to process nums2.
	//
	// Edge case: if nums2 is empty (n = 0), j = -1,
	// so the loop doesn't execute and nums1 remains unchanged.
	//
	// Why don't we need "i >= 0" in the loop condition?
	// Because if nums1 becomes empty first, we can simply copy
	// the remaining nums2 elements.
	for j >= 0 {

		// i >= 0 handles the edge case where all original
		// nums1 elements have already been used.
		//
		// Example:
		// nums1 = [0,0,0], m = 0
		// nums2 = [1,2,3]
		// Here i = -1, so we MUST NOT access nums1[i].
		//
		// If nums1[i] is larger, put it at k because we are
		// filling the largest remaining element first.
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			// nums2[j] is larger OR both values are equal.
			//
			// Example:
			// nums1[i] = 2, nums2[j] = 2
			// We can safely take nums2[j].
			nums1[k] = nums2[j]
			j--
		}

		// Move k left because the current largest element
		// has been placed.
		k--
	}

	// If nums1 still has elements remaining, do nothing.
	// They are already in the correct positions.
	//
	// Example:
	// nums1 = [4,5,6,0,0,0]
	// nums2 = [1,2,3]
	//
	// After nums2 is exhausted:
	// nums1 = [1,2,3,4,5,6]
	//
	// Therefore, we don't need another loop to copy nums1.
}