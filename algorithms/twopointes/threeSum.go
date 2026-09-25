package twopointes

import "sort"

func threeSum(nums []int) [][]int {
	// Sorting is important because it lets us use two pointers.
	// It also puts duplicate values next to each other, making
	// duplicate triplets easy to skip.
	sort.Ints(nums)

	res := [][]int{}

	// We need at least 3 elements to form a triplet,
	// so i only needs to go up to len(nums)-2.
	for i := 0; i < len(nums)-2; i++ {

		// Skip duplicate values for the first number.
		//
		// Example:
		// [-1,-1,0,1,2]
		//  ↑  ↑
		//  i
		//
		// Both -1s would search for the same remaining pair,
		// so processing the second -1 would create duplicate results.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// left starts after i because we need 3 different positions:
		// i < left < right.
		//
		// Starting left at i would allow the same element to
		// be used twice.
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// We found a valid triplet.
				// Store the actual VALUES, not their indices.
				res = append(res, []int{
					nums[i],
					nums[left],
					nums[right],
				})

				// We found one solution, so both current values
				// have already been used.
				// Move both pointers to search for another pair.
				left++
				right--

				// Skip duplicate left values.
				//
				// Example:
				// [-1,0,0,0,1]
				//     ↑ ↑
				//
				// Using the next 0 would produce the same triplet,
				// so skip all consecutive duplicate 0s.
				//
				// left < right must be checked FIRST so we don't
				// access an invalid position after the pointers cross.
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// Skip duplicate right values for the same reason.
				//
				// Example:
				// [-2,0,1,1,2]
				//         ↑ ↑
				//
				// Reusing the duplicate 1 would create the same triplet.
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum > 0 {
				// The array is sorted.
				//
				// Sum is too large, so we need a smaller number.
				// Moving right left gives us a smaller value.
				right--

			} else {
				// Sum is too small, so we need a larger number.
				// Moving left right gives us a larger value.
				left++
			}
		}
	}

	return res
}