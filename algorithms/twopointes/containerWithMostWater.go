package twopointes

func maxArea(heights []int) int {
	maxArea := 0

	left, right := 0, len(heights)-1

	// WHY TWO POINTERS?
	//
	// Brute force would check every pair:
	//
	// (0,1), (0,2), (0,3)...
	// (1,2), (1,3)...
	//
	// That is O(n²).
	//
	// Instead, we start with the widest possible container:
	//
	// left = first wall
	// right = last wall
	//
	// This gives us the maximum possible WIDTH.
	//
	// Every time we move a pointer, the width becomes smaller.
	// Therefore, we need to make sure we only throw away positions
	// that cannot possibly produce a better answer.
	for left < right {

		// The amount of water is limited by the SHORTER wall.
		//
		// Example:
		//
		// left height  = 8
		// right height = 5
		//
		//       8
		//       |
		//       |
		//       |     5
		//       |     |
		//       |     |
		//
		// Water can only reach height 5.
		//
		// So:
		// area = shorter height × width
		minIndex := left

		if heights[right] < heights[minIndex] {
			minIndex = right
		}

		area := heights[minIndex] * (right - left)

		if area > maxArea {
			maxArea = area
		}

		// WHY DO WE MOVE THE SHORTER HEIGHT?
		//
		// Suppose:
		//
		// left  = 3
		// right = 8
		//
		// Current area:
		//
		// 3 × (right-left)
		//
		// The height is limited by 3.
		//
		// ------------------------------------------------
		// What happens if we move the TALLER side (8)?
		// ------------------------------------------------
		//
		// Example:
		//
		// left = 3
		// right = 8
		//
		// If we move right:
		//
		//     width becomes smaller
		//     ↓
		//     shorter height is STILL 3
		//
		// So the new area is:
		//
		// 3 × smaller_width
		//
		// This can NEVER be greater than our current:
		//
		// 3 × larger_width
		//
		// Therefore, moving the taller side can NEVER help.
		//
		// ------------------------------------------------
		// What happens if we move the SHORTER side (3)?
		// ------------------------------------------------
		//
		// We know the width will decrease.
		//
		// But there is at least a POSSIBILITY that the new
		// wall is taller than 3.
		//
		// Example:
		//
		// Current:
		// height = min(3, 8) = 3
		//
		// Move left:
		//
		// new left height = 7
		// right height    = 8
		//
		// Now:
		// height = min(7, 8) = 7
		//
		// We lost some width, but gained a much larger height.
		// Therefore, a larger area is POSSIBLE.
		//
		// So:
		//
		// SHORTER SIDE → MUST MOVE
		// TALLER SIDE   → NEVER HELPS
		if left == minIndex {
			left++
		} else {
			right--
		}
	}

	return maxArea
}