package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []int
	size int
}

var ErrStackIsEmpty = errors.New("stack is empty")

func (s *Stack) Push(value int) {
	// append adds the new element at the end of the slice.
	// The end of the slice represents the top of the stack.
	s.data = append(s.data, value)

	// Track the number of elements currently in the stack.
	s.size++
}

func (s *Stack) Pop() (int, error) {
	// We cannot pop anything if the stack is empty.
	if s.size == 0 {
		return 0, ErrStackIsEmpty
	}

	// The last element is the top of the stack,
	// because a stack follows LIFO: Last In, First Out.
	pop := s.data[s.size-1]

	// Remove the last element from the slice.
	// Slicing up to size-1 excludes the current top element.
	s.data = s.data[:s.size-1]

	// One element has been removed.
	s.size--

	return pop, nil
}

func (s *Stack) Peek() (int, error) {
	// There is no top element when the stack is empty.
	if s.size == 0 {
		return 0, ErrStackIsEmpty
	}

	// Return the top element without removing it.
	return s.data[s.size-1], nil
}

func (s *Stack) IsEmpty() bool {
	// The stack is empty when it contains zero elements.
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Display() {
	// Iterate through all elements currently in the stack.
	for i := range s.data {
		fmt.Print(s.data[i], " ")
	}
}