package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type SinglyLinkedListWithoutTail struct {
	head *Node
	size int
}

var ErrListIsEmpty = errors.New("list is empty")

func (s *SinglyLinkedListWithoutTail) InsertAtFront(value int) {

	newNode := &Node{
		value: value,
	}

	// The new node becomes the first node,
	// so it must point to the current head.
	newNode.next = s.head

	// Update head because the new node is now the first node.
	s.head = newNode

	s.size++
}

func (s *SinglyLinkedListWithoutTail) InsertAtRear(value int) {
	newNode := &Node{value: value}

	// If the list is empty, the new node itself becomes the head.
	if s.size == 0 {
		s.head = newNode
		s.size++
		return
	}

	currentNode := s.head

	// Traverse until we reach the last node.
	// The last node is identified by next == nil.
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	// Connect the current last node to the new node.
	currentNode.next = newNode

	s.size++
}

func (s *SinglyLinkedListWithoutTail) RemoveFront() (int, error) {

	if s.size == 0 {
		return 0, ErrListIsEmpty
	}

	node := s.head

	// Move head to the second node.
	// This removes the old head from the list.
	s.head = s.head.next

	s.size--

	return node.value, nil
}

func (s *SinglyLinkedListWithoutTail) RemoveRear() (int, error) {

	if s.size == 0 {
		return 0, ErrListIsEmpty
	}

	node := s.head

	// Special case: if there is only one node,
	// removing the rear means removing the head as well.
	if s.head.next == nil {
		s.head = nil
		s.size--
		return node.value, nil
	}

	current := s.head

	// Stop at the node immediately before the last node.
	// We need this node because it must point to nil
	// after the last node is removed.
	for current.next.next != nil {
		current = current.next
	}

	removedNode := current.next

	// Disconnect the last node from the list.
	current.next = nil

	s.size--

	return removedNode.value, nil
}

func (s *SinglyLinkedListWithoutTail) Reverse() error {
	if s.size == 0 {
		return ErrListIsEmpty
	}

	var prev *Node
	current := s.head

	for current != nil {

		// Save the next node before changing current.next.
		// Without this, we would lose access to the rest of the list.
		temp := current.next

		// Reverse the direction of the current node's pointer.
		current.next = prev

		// Move prev forward because current is now
		// the previous node for the next iteration.
		prev = current

		// Move current forward using the saved pointer.
		current = temp
	}

	// prev is now the new first node of the reversed list.
	s.head = prev

	return nil
}

func (s *SinglyLinkedListWithoutTail) Size() int {
	return s.size
}

func (s *SinglyLinkedListWithoutTail) IsEmpty() bool {
	return s.size == 0
}

func (s *SinglyLinkedListWithoutTail) Display() {
	current := s.head

	for current != nil {
		fmt.Print(current.value, "->")

		// Move to the next node.
		// Without this, current would always point to the same node
		// and the loop would never terminate.
		current = current.next
	}

	fmt.Println(nil)
}

func (s *SinglyLinkedListWithoutTail) Insert(value, index int) {

	// Valid insertion positions are 0 through size.
	// size is valid because inserting at size means inserting at the rear.
	if index < 0 || index > s.size {
		return
	}

	// Inserting at index 0 means inserting at the front.
	if index == 0 {
		s.InsertAtFront(value)
		return
	}

	current := s.head

	newNode := &Node{
		value: value,
	}

	// Stop at the node immediately before the insertion position.
	// Example: inserting at index 2 means we stop at index 1.
	for i := 1; i < index; i++ {
		current = current.next
	}

	// First connect the new node to the node currently at the index.
	newNode.next = current.next

	// Then connect the previous node to the new node.
	// This inserts the new node into the chain.
	current.next = newNode

	s.size++
}

func (s *SinglyLinkedListWithoutTail) RemoveByIndex(index int) {
	if index < 0 || index >= s.size {
		return
	}

	// Removing index 0 means removing the head.
	if index == 0 {
		s.head = s.head.next
		s.size--
		return
	}

	current := s.head

	// Stop at the node immediately before the node we want to remove.
	for i := 1; i < index; i++ {
		current = current.next
	}

	// Skip over the node at index.
	// This changes the actual linked-list structure.
	//
	// current -> nodeToRemove -> nextNode
	//
	// becomes:
	//
	// current -------------> nextNode
	current.next = current.next.next

	s.size--
}

func (s *SinglyLinkedListWithoutTail) RemoveByValue(value int) {
	if s.size == 0 {
		return
	}

	current := s.head

	// Handle the head separately because there is no previous
	// node whose next pointer we can modify.
	if current.value == value {
		s.head = s.head.next
		s.size--
		return
	}

	// Search for the node whose value matches.
	// We inspect current.next because we need the previous node
	// to modify its next pointer.
	for current.next != nil {
		if current.next.value == value {
			break
		}

		current = current.next
	}

	// The value wasn't found.
	if current.next == nil {
		return
	}

	// Skip the matching node.
	//
	// current -> matchingNode -> nextNode
	//
	// becomes:
	//
	// current ----------------> nextNode
	current.next = current.next.next

	s.size--
}