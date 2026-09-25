package twopointes

func trapO(height []int) int {
	n := len(height)

	// WHY?
	// We need at least one position to calculate anything.
	// If there are no walls, no water can be trapped.
	if n == 0 {
		return 0
	}

	// WHY DO WE NEED THE TALLEST WALL ON THE LEFT AND RIGHT?
	//
	// Imagine we are standing at index i.
	//
	// height = [4, 2, 0, 3, 2, 5]
	//                  ↑
	//              current wall
	//
	// Can water stay above this current wall?
	//
	// Only if there is a wall on BOTH sides of it.
	//
	// Example:
	//
	//        4               5
	//        |               |
	//        |~~~~~~~~~~~~~~~|
	//        |   water       |
	//        |       0       |
	//
	// The walls on the left and right stop the water from flowing away.
	//
	// So, for the current position, we need to know:
	//
	// 1. How HIGH can the wall on the left be?
	// 2. How HIGH can the wall on the right be?
	//
	// We care about the TALLEST wall on each side because a shorter
	// wall on that same side doesn't limit the maximum amount of water.
	//
	// Example:
	//
	// Left side:  [4, 2, 1, 3]
	//                    ↑ current
	//
	// The tallest wall on the left is 4.
	// The walls 2 and 1 are irrelevant for the maximum boundary
	// because the wall of height 4 is the highest wall available
	// on that side.
	//
	// Similarly, we need the tallest wall on the right.
	//
	// Once we know both:
	//
	//     leftMax  = tallest wall on the left
	//     rightMax = tallest wall on the right
	//
	// the shorter of these two walls determines how high water
	// can actually rise.
	//
	// Why?
	//
	// If:
	//
	//     leftMax  = 5
	//     rightMax = 3
	//
	// water cannot rise to height 5 because it would overflow
	// over the right wall of height 3.
	//
	// Therefore:
	//
	//     water level = min(leftMax, rightMax)
	//
	// And because the current wall already occupies some height:
	//
	//     trapped water = water level - height[i]
	//
	// So the whole problem becomes:
	//
	//     water[i] = min(leftMax, rightMax) - height[i]
	//
	// We create two arrays to remember these maximum boundaries
	// for every index, so we don't have to repeatedly scan left
	// and right for every position.
	left := make([]int, n)
	right := make([]int, n)

	leftMax := height[0]
	rightMax := height[n-1]

	// WHY BUILD left[]?
	//
	// For every index i, store the tallest wall we have seen
	// from the LEFT up to that position.
	//
	// Example:
	//
	// height = [4, 2, 0, 3, 2, 5]
	//
	// left =   [4, 4, 4, 4, 4, 5]
	//
	// At index 3:
	//
	// height = 3
	// tallest wall on left = 4
	//
	// We store 4 in left[3].
	for i := 0; i < n; i++ {

		// If this wall is taller than the tallest wall
		// we've seen so far, it becomes the new leftMax.
		if height[i] > leftMax {
			leftMax = height[i]
		}

		left[i] = leftMax
	}

	// WHY BUILD right[]?
	//
	// Same idea, but from the RIGHT.
	//
	// For every index i, store the tallest wall we have seen
	// from the right up to that position.
	//
	// We iterate backwards because we're calculating information
	// about the right side.
	for i := n - 1; i >= 0; i-- {

		// If this wall is taller than the tallest wall
		// we've seen from the right, update rightMax.
		if height[i] > rightMax {
			rightMax = height[i]
		}

		right[i] = rightMax
	}

	total := 0

	for i := 0; i < n; i++ {

		// WHY TAKE THE SHORTER OF leftMax AND rightMax?
		//
		// These two walls form the boundaries that can hold water.
		//
		// Example:
		//
		// leftMax  = 5
		// rightMax = 4
		//
		// Even though the left wall can hold water up to 5,
		// the water would overflow over the right wall at 4.
		//
		// Therefore, the maximum water level is 4.
		min := left[i]

		if right[i] < min {
			min = right[i]
		}

		// WHY SUBTRACT height[i]?
		//
		// "min" tells us the maximum water LEVEL.
		//
		// But the current bar already occupies height[i].
		//
		// Example:
		//
		// water level = 5
		// current wall = 2
		//
		// The empty space available for water is:
		//
		//     5 - 2 = 3
		//
		// So:
		//
		//     trapped water = min(left[i], right[i]) - height[i]
		total += min - height[i]
	}

	return total
}

func trap(height []int) int {
	total := 0

	// WHY TWO POINTERS?
	//
	// Earlier, we stored:
	//
	// leftMax[i]  = tallest wall to the left of i
	// rightMax[i] = tallest wall to the right of i
	//
	// That required O(n) extra space.
	//
	// But notice:
	// We don't actually need the entire leftMax[] and rightMax[]
	// arrays.
	//
	// While moving from the outside toward the middle, we only
	// need the CURRENT tallest wall we've seen from each side.
	//
	// Therefore, we can keep only:
	//
	// leftMax  = tallest wall seen from the left so far
	// rightMax = tallest wall seen from the right so far
	//
	// This reduces extra space from O(n) to O(1).
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0

	// WHY left <= right?
	//
	// Each iteration processes one position and then moves either
	// left or right.
	//
	// Eventually both pointers can meet at the same index:
	//
	//     left == right
	//
	// That index is still an unprocessed position, so we allow
	// left == right and process it once.
	//
	// After processing it, the pointers cross:
	//
	//     left > right
	//
	// and the loop stops.
	for left <= right {

		// WHY DO WE PROCESS THE LEFT SIDE WHEN leftMax <= rightMax?
		//
		// Remember the water formula:
		//
		// water = min(leftMax, rightMax) - height[i]
		//
		// If:
		//
		// leftMax <= rightMax
		//
		// then we already know that the LEFT side is the limiting
		// boundary.
		//
		// Even if there is some taller wall somewhere on the right,
		// water at the current left position cannot rise above
		// leftMax because leftMax is already the shorter boundary.
		//
		// Therefore, we can safely calculate the water on the LEFT
		// without knowing the exact right-side maximum.
		if leftMax <= rightMax {

			// Update the tallest wall we've seen from the left.
			//
			// Example:
			// leftMax = 4
			// height[left] = 6
			//
			// The new leftMax becomes 6.
			leftMax = max(leftMax, height[left])

			// WHY leftMax - height[left]?
			//
			// leftMax is the limiting boundary for this position.
			//
			// Example:
			//
			// leftMax    = 5
			// height[left] = 2
			//
			// water = 5 - 2 = 3
			//
			// If height[left] == leftMax:
			//
			// water = 5 - 5 = 0
			//
			// So no separate check for negative values is needed.
			total = total + leftMax - height[left]

			// This position has been completely processed.
			left++

		} else {

			// WHY PROCESS THE RIGHT SIDE HERE?
			//
			// This means:
			//
			// rightMax < leftMax
			//
			// So the RIGHT side is currently the limiting boundary.
			//
			// Therefore, we can safely calculate the water at the
			// current right position using rightMax.
			rightMax = max(rightMax, height[right])

			// Same idea as the left side:
			//
			// water = rightMax - currentHeight
			//
			// Example:
			//
			// rightMax = 6
			// height[right] = 2
			//
			// water = 6 - 2 = 4
			total = total + rightMax - height[right]

			// This position has been completely processed.
			right--
		}
	}

	return total
}
