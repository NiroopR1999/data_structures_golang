package linkedlist

func findDuplicate(nums []int) int {
	// Treat each array value as the next position to visit.
	// Example: nums[2] = 4 means from index 2, go to index 4.
	//
	// Because every value is between 1 and n, every value
	// points to another valid index.
	slow := nums[0]
	fast := nums[0]

	// Phase 1: Find a meeting point inside the cycle.
	//
	// slow moves one step at a time.
	// fast moves two steps at a time.
	//
	// Because a cycle exists, they must eventually meet.
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	// Phase 2: Find the entrance of the cycle.
	//
	// The entrance of the cycle represents the duplicate number.
	//
	// Why?
	// The duplicate value is the only value that is reached
	// from two different indices, which causes the cycle.
	slow = nums[0]

	// Move both pointers one step at a time.
	//
	// They will meet at the cycle entrance.
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	// The meeting point is the duplicate number.
	return slow
}