package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	// Both start at the head so we can use two pointers moving at different speeds.
	slow, fast := head, head

	// Fast must be valid and have a next node because it moves two steps at a time.
	for fast != nil && fast.Next != nil {

		// Slow moves one step at a time.
		slow = slow.Next

		// Fast moves two steps at a time, allowing it to eventually catch slow if a cycle exists.
		fast = fast.Next.Next

		// If both point to the same node, fast has caught slow inside a cycle.
		if slow == fast {
			return true
		}
	}

	// If fast reaches the end, the list has no cycle.
	return false
}