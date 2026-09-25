package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// Dummy gives the head node a previous node.
	// This allows us to remove the head using the same logic
	// as removing any other node.
	dummy := &ListNode{Next: head}

	slow, fast := dummy, dummy

	count := 0

	// Move fast n+1 steps ahead.
	// The extra 1 creates the gap needed so that when
	// fast reaches nil, slow is immediately before the target node.
	for count <= n {
		count++
		fast = fast.Next
	}

	// Move both pointers together while fast is not nil.
	// Because the gap is n+1, slow will stop at the node
	// immediately before the nth node from the end.
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Skip the target node.
	// Example: 3 → 4 → 5 becomes 3 → 5.
	slow.Next = slow.Next.Next

	// dummy.Next is the actual head of the modified list.
	// This also works when the original head was removed.
	return dummy.Next
}