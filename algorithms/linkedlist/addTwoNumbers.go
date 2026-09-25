package linkedlist


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Dummy node makes it easy to build the result list
	// without needing special handling for the first node.
	dummy := &ListNode{}
	current := dummy

	// Stores the carry when the sum of two digits is >= 10.
	carry := 0

	// Continue while either list still has digits or
	// there is a carry left to add.
	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0

		// If l1 has a digit, use it and move to the next node.
		// Otherwise, treat the missing digit as 0.
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		// Same logic for l2.
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		// Add both digits along with the carry from the previous position.
		val := v1 + v2 + carry

		// Everything except the last digit becomes the carry.
		// Example: 17 / 10 = 1.
		carry = val / 10

		// Only the last digit belongs in the current node.
		// Example: 17 % 10 = 7.
		current.Next = &ListNode{Val: val % 10}

		// Move current forward so the next digit is appended after it.
		current = current.Next
	}

	// The dummy node itself is not part of the answer,
	// so return the first actual node.
	return dummy.Next
}