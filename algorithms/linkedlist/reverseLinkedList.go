package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	// prev will become the new previous node for current.
	// Initially there is nothing before the first node.
	var prev *ListNode

	current := head

	for current != nil {
		// Save the next node before changing current.Next.
		// Otherwise, we would lose the rest of the linked list.
		next := current.Next

		// Reverse the direction of the current node.
		// Instead of pointing forward to next,
		// it now points backward to prev.
		current.Next = prev

		// Move prev forward.
		// Current is now the first node of the reversed portion.
		prev = current

		// Move current forward using the node we saved earlier.
		current = next
	}

	// current becomes nil after the loop.
	// prev is now the first node of the reversed list.
	return prev
}
