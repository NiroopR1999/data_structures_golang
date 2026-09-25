package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// Dummy node handles the case where left == 1.
	// Without it, we would need separate logic to update the head.
	dummy := &ListNode{Next: head}

	// leftPrev will point to the node immediately before
	// the part that needs to be reversed.
	leftPrev := dummy

	// Move leftPrev to position left-1.
	for i := 1; i < left; i++ {
		leftPrev = leftPrev.Next
	}

	// current is the first node that needs to be reversed.
	current := leftPrev.Next

	var prev *ListNode

	// Reverse exactly the nodes from left to right.
	// +1 is needed because both left and right are included.
	for i := 1; i <= right-left+1; i++ {
		// Save the next node before changing current.Next.
		// Otherwise, we would lose the rest of the list.
		next := current.Next

		// Reverse the direction of the current node.
		current.Next = prev

		// Move prev forward because current is now
		// the first node of the reversed portion.
		prev = current

		// Move current forward using the node we saved.
		current = next
	}

	// leftPrev.Next is still the original first node
	// of the reversed section.
	// After reversal, that node becomes the tail.
	// Connect it to the node after 'right'.
	leftPrev.Next.Next = current

	// Connect the node before the reversed section
	// to the new first node of the reversed section.
	leftPrev.Next = prev

	// dummy itself is not part of the answer.
	// dummy.Next is the actual head of the resulting list.
	return dummy.Next
}
