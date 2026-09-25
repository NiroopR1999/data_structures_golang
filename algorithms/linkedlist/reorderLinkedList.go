package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {

	// Nothing to reorder if the list has 0 or 1 node.
	if head == nil || head.Next == nil {
		return
	}

	// Use slow/fast pointers to find the middle of the list.
	// Fast moves 2 steps while slow moves 1 step.
	// When fast reaches the end, slow is around the middle.
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// The second half starts after slow.
	current := slow.Next

	// Break the list into two separate lists.
	// Without this, the first and second halves would still be connected.
	slow.Next = nil

	// Reverse the second half.
	// prev will eventually become the head of the reversed second half.
	var prev *ListNode

	for current != nil {
		// Save the next node because current.Next is about to be changed.
		// Without saving it, we would lose access to the remaining list.
		temp := current.Next

		// Reverse the pointer.
		current.Next = prev

		// Move prev forward so it becomes the new head of the reversed part.
		prev = current

		// Continue processing the remaining nodes.
		current = temp
	}

	first := head
	second := prev

	// Merge the two halves alternately:
	// first node from first half,
	// first node from second half,
	// second node from first half,
	// second node from second half, ...
	//
	// We iterate using second because the second half is never longer
	// than the first half.
	for second != nil {

		// Save both next nodes before changing any pointers.
		// These are necessary because the connections will be overwritten.
		firstNext := first.Next
		secondNext := second.Next

		// Insert the second-half node after the first-half node.
		first.Next = second

		// Connect the inserted node to the original next node of first.
		// This is why firstNext had to be saved above.
		second.Next = firstNext

		// Move to the next unused nodes in both halves.
		first = firstNext
		second = secondNext
	}
}