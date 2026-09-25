package linkedlist
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists merges two sorted linked lists into one sorted linked list.
// It returns the head of the merged list.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// head node avoids special handling for the head of the merged list.
	head := &ListNode{}
	curr := head

	// Continue until one list is completely consumed.
	for list1 != nil && list2 != nil {

		// Pick the smaller node so the merged list remains sorted.
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		// Move forward in the merged list.
		curr = curr.Next
	}

	// One list may still contain nodes.
	// Since it is already sorted, we can attach it directly.
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	head=head.Next
	return head
}