package twopointes

import "sort"

func fourSum(nums []int, target int) [][]int {
	// WHY: Sorting lets us:
	// 1. Use two pointers (left/right)
	// 2. Easily skip duplicate values
	// 3. Decide which pointer to move based on the sum
	sort.Ints(nums)

	res := [][]int{}

	// WHY: We need at least 4 numbers, so i can go only up to len(nums)-4.
	for i := 0; i < len(nums)-3; i++ {

		// WHY: If the current i has the same value as the previous i,
		// we would generate the exact same quadruplets again.
		// Example: [-2, -2, 0, 0, 2, 2]
		// We process the first -2, so skip the second -2 as i.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// WHY: After fixing nums[i], we fix a second number nums[j].
		// The remaining two numbers will be found using left/right.
		for j := i + 1; j < len(nums)-2; j++ {

			// WHY: j must be the first occurrence for this position.
			// Important: use j > i+1, NOT j > 0.
			//
			// Example: [0, 0, 0, 0]
			// i=0, j=1 is the FIRST valid j and must NOT be skipped.
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// WHY: left starts immediately after j because all 4
			// positions must be different.
			left, right := j+1, len(nums)-1

			// WHY: Once left >= right, there are no two different
			// positions remaining to form a quadruplet.
			for left < right {

				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {

					// WHY: We found a valid quadruplet.
					// Store VALUES, not indexes.
					res = append(res, []int{
						nums[i],
						nums[j],
						nums[left],
						nums[right],
					})

					// WHY: We found one answer, so move both pointers
					// to search for another possible combination.
					left++
					right--

					// WHY: If the new left value is the same as the
					// previous left value, it would produce the same
					// quadruplet again.
					//
					// Example: [1, 1, 1, 2, 3]
					// After using the first 1 as left, skip the other 1s.
					for left < right && nums[left] == nums[left-1] {
						left++
					}

					// WHY: Same idea for right.
					// Skip repeated values so we don't add duplicate
					// quadruplets.
					for left < right && nums[right] == nums[right+1] {
						right--
					}

				} else if sum > target {

					// WHY: Array is sorted.
					// Sum is too large, so decrease right to get a
					// smaller value and therefore a smaller sum.
					right--

				} else {

					// WHY: Sum is too small, so increase left to get
					// a larger value and therefore a larger sum.
					left++
				}
			}
		}
	}

	return res
}