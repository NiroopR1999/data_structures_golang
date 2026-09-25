package linkedlist

import "fmt"

type SinglyLinkedListWithTail struct {
	head *Node
	tail *Node
	size int
}

func (s *SinglyLinkedListWithTail) InsertAtFront(value int) {
	newNode := &Node{
		value: value,
	}

	// If the list is empty, the new node is both
	// the first node and the last node.
	if s.head == nil {
		s.tail = newNode
	}

	// The new node must point to the current head
	// because it is being inserted before the existing first node.
	newNode.next = s.head

	// Make the new node the first node of the list.
	s.head = newNode

	s.size++
}

func (s *SinglyLinkedListWithTail) InsertAtRear(value int) {
	// If the list is empty, there is no existing tail.
	// InsertAtFront handles both head and tail initialization.
	if s.head == nil {
		s.InsertAtFront(value)
		return
	}

	newNode := &Node{value: value}

	// Since we maintain a tail pointer, we already have
	// direct access to the last node.
	// No traversal from head is required.
	s.tail.next = newNode

	// The newly inserted node is now the last node.
	s.tail = newNode

	s.size++
}

func (s *SinglyLinkedListWithTail) RemoveFront() (int, error) {
	if s.size == 0 {
		return 0, ErrListIsEmpty
	}

	node := s.head

	// Move head to the second node.
	// The old head is no longer reachable from the list.
	s.head = s.head.next

	s.size--

	// If the removed node was the only node,
	// the list is now empty, so tail must also be nil.
	if s.size == 0 {
		s.tail = nil
	}

	return node.value, nil
}

func (s *SinglyLinkedListWithTail) RemoveRear() (int, error) {
	if s.size == 0 {
		return 0, ErrListIsEmpty
	}

	node := s.head

	// If there is only one node, it is both head and tail.
	// Removing it makes the entire list empty.
	if s.head.next == nil {
		s.head = nil
		s.tail = nil
		s.size--
		return node.value, nil
	}

	current := s.head

	// We need the node immediately before the tail.
	// A singly linked list has no prev pointer,
	// so we must traverse from the head to find it.
	for current.next.next != nil {
		current = current.next
	}

	removedNode := current.next

	// current becomes the new last node.
	s.tail = current

	// Remove the old tail by making the new tail
	// point to nil.
	current.next = nil

	s.size--

	return removedNode.value, nil
}

func (s *SinglyLinkedListWithTail) Reverse() error {
	if s.size == 0 {
		return ErrListIsEmpty
	}

	// The old head becomes the new tail after reversal.
	oldHead := s.head

	var prev *Node
	current := s.head

	for current != nil {
		// Save the next node before changing current.next.
		// Otherwise, we would lose access to the remaining list.
		temp := current.next

		// Reverse the current node's pointer.
		// Example:
		// 20 -> 30
		// becomes:
		// 20 <- 30
		current.next = prev

		// Move prev forward because current
		// is now the previous node for the next iteration.
		prev = current

		// Continue through the original list using
		// the pointer we saved before reversing it.
		current = temp
	}

	// prev is now the first node of the reversed list.
	s.head = prev

	// The original first node is now the last node.
	s.tail = oldHead

	return nil
}

func (s *SinglyLinkedListWithTail) Size() int {
	return s.size
}

func (s *SinglyLinkedListWithTail) IsEmpty() bool {
	return s.size == 0
}

func (s *SinglyLinkedListWithTail) Display() {
	current := s.head

	for current != nil {
		fmt.Print(current.value, "->")

		// Move to the next node.
		// Without this, current would never change
		// and the loop would never terminate.
		current = current.next
	}

	fmt.Println(nil)
}

func (s *SinglyLinkedListWithTail) Insert(value, index int) {
	// Valid insertion positions are 0 through size.
	// size is valid because inserting at size means
	// inserting after the current last node.
	if index < 0 || index > s.size {
		return
	}

	// Index 0 means the new node becomes the head.
	if index == 0 {
		s.InsertAtFront(value)
		return
	}

	// Index == size means inserting after the current tail.
	// InsertAtRear also updates the tail correctly.
	if index == s.size {
		s.InsertAtRear(value)
		return
	}

	current := s.head

	newNode := &Node{
		value: value,
	}

	// Stop at the node immediately before the insertion position.
	//
	// Example:
	// 10 -> 20 -> 30
	//
	// Insert 15 at index 1:
	//
	// current = 10
	//
	// We need current.next (20) so that we can place
	// the new node between 10 and 20.
	for i := 1; i < index; i++ {
		current = current.next
	}

	// First make the new node point to the node
	// that currently occupies the insertion position.
	newNode.next = current.next

	// Then make the previous node point to the new node.
	//
	// Before:
	// current -> next
	//
	// After:
	// current -> newNode -> next
	current.next = newNode

	s.size++
}

func (s *SinglyLinkedListWithTail) RemoveByIndex(index int) {
	// Valid removal positions are 0 through size-1.
	if index < 0 || index >= s.size {
		return
	}

	// Removing index 0 means removing the head.
	// RemoveFront handles the head and tail edge cases.
	if index == 0 {
		s.RemoveFront()
		return
	}

	current := s.head

	// Stop at the node immediately before the node
	// we want to remove.
	for i := 1; i < index; i++ {
		current = current.next
	}

	// current.next is the node we want to remove.
	removedNode := current.next

	// If we are removing the tail,
	// current becomes the new tail.
	if removedNode == s.tail {
		s.tail = current
	}

	// Skip over removedNode.
	//
	// Before:
	// current -> removedNode -> next
	//
	// After:
	// current ----------------> next
	//
	// This changes the actual linked-list structure.
	current.next = removedNode.next

	s.size--
}

func (s *SinglyLinkedListWithTail) RemoveByValue(value int) {
	if s.size == 0 {
		return
	}

	// Handle the head separately because there is no
	// previous node whose next pointer we can modify.
	if s.head.value == value {
		s.RemoveFront()
		return
	}

	current := s.head

	// We inspect current.next because we need current
	// to be the node immediately before the node we remove.
	for current.next != nil {
		if current.next.value == value {
			break
		}

		current = current.next
	}

	// We reached the end without finding the value.
	if current.next == nil {
		return
	}

	// current.next is the node that will be removed.
	removedNode := current.next

	// If the node being removed is the tail,
	// current becomes the new tail.
	if removedNode == s.tail {
		s.tail = current
	}

	// Skip over the matching node.
	//
	// Before:
	// current -> removedNode -> next
	//
	// After:
	// current ----------------> next
	current.next = removedNode.next

	s.size--
}