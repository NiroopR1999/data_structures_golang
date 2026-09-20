package linkedlist

import (
	"errors"
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

	newNode.next = s.head
	s.head = newNode
	s.size++

}

func (s *SinglyLinkedListWithoutTail) InsertAtRear(value int) {
	newNode := &Node{value: value}

	if s.size == 0 {
		s.head = newNode
		s.size++
		return
	}
	currentNode := s.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	s.size++
	currentNode.next = newNode

}

func (s *SinglyLinkedListWithoutTail) RemoveFront() (int, error) {

	if s.size == 0 {
		return 0, ErrListIsEmpty
	}
	node := s.head
	s.head = s.head.next
	s.size--
	return node.value, nil
}

func (s *SinglyLinkedListWithoutTail) RemoveRear() (int, error) {

	if s.size == 0 {
		return 0, ErrListIsEmpty
	}
	node := s.head
	if s.head.next == nil {
		s.head = nil
		s.size--
		return node.value, nil
	}
	current := s.head
	for current.next.next != nil {
		current = current.next
	}
	removedNode := current.next
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
		temp := current.next
		current.next = prev
		prev = current
		current = temp

	}
	s.head = prev
	return nil
}

func (s *SinglyLinkedListWithoutTail) Size() int {
	return s.size
}

func (s *SinglyLinkedListWithoutTail) IsEmpty() bool {
	return s.size == 0
}
