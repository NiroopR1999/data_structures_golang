package queue

import "errors"

type Item struct {
	Value    any
	Priority int
}

type MaxPriorityQueue struct {
	data []Item
}

func (m *MaxPriorityQueue) siftUp(index int) {

	// Index 0 is the root. It has no parent, so there is nothing to compare.
	if index == 0 {
		return
	}

	// Find the parent so we can check whether the Max Heap property is violated.
	parentIndex := (index - 1) / 2

	// If the parent already has equal or higher Priority,
	// the Max Heap property is satisfied and we can stop.
	if m.data[parentIndex].Priority >= m.data[index].Priority {
		return
	}

	// The child has higher Priority than its parent,
	// so move the higher-Priority item upward.
	m.data[index], m.data[parentIndex] =
		m.data[parentIndex], m.data[index]

	// The item moved up may still have higher Priority than its new parent,
	// so continue checking upward.
	m.siftUp(parentIndex)
}

func (m *MaxPriorityQueue) Insert(Value any, Priority int) {

	item := Item{
		Value:    Value,
		Priority: Priority,
	}

	// Add the item at the end so the heap remains a complete binary tree.
	m.data = append(m.data, item)

	// The new item may violate the Max Heap property with its parent,
	// so move it upward until the property is restored.
	m.siftUp(len(m.data) - 1)
}

func (m *MaxPriorityQueue) Peek() (any, error) {

	// Accessing index 0 on an empty slice would cause a panic.
	if len(m.data) == 0 {
		return nil, errors.New("queue is empty")
	}

	// The root always has the highest Priority in a Max Priority Queue.
	return m.data[0].Value, nil
}

func (m *MaxPriorityQueue) ExtractMax() (Item, error) {

	// There is no item to extract, and accessing index 0 would panic.
	if len(m.data) == 0 {
		return Item{}, errors.New("queue is empty")
	}

	// Save the root because it is the highest-Priority item we need to return.
	max := m.data[0]

	if len(m.data) == 1 {

		// Removing the only item leaves the Priority queue empty.
		m.data = []Item{}

		return max, nil
	}

	// Move the last item to the root.
	// Removing the root directly would break the complete binary tree structure.
	m.data[0] = m.data[len(m.data)-1]

	// Remove the old last item.
	// Removing from the end preserves the complete binary tree structure.
	m.data = m.data[:len(m.data)-1]

	// The new root may have lower Priority than one of its children,
	// so move it downward until the Max Heap property is restored.
	m.siftDown(0)

	return max, nil
}

func (m *MaxPriorityQueue) siftDown(index int) {

	// Calculate the left child because we need at least one child
	// before there is anything to compare.
	leftChildIndex := (2 * index) + 1

	// The right child is always immediately after the left child in the array.
	rightChildIndex := leftChildIndex + 1

	// Assume the left child has the highest Priority.
	// We will change this if the right child exists and has higher Priority.
	maxChildIndex := leftChildIndex

	// If there is no left child, this node is a leaf.
	// A leaf has no children, so there is nothing to sift down.
	if leftChildIndex >= len(m.data) {
		return
	}

	// The right child may not exist because the last level of a complete
	// binary tree can have a missing right child.
	// If it exists, choose whichever child has higher Priority.
	if rightChildIndex < len(m.data) &&
		m.data[maxChildIndex].Priority < m.data[rightChildIndex].Priority {

		maxChildIndex = rightChildIndex
	}

	// If the parent already has equal or higher Priority than its highest-Priority child,
	// the Max Heap property is satisfied, so we can stop.
	if m.data[index].Priority >= m.data[maxChildIndex].Priority {
		return
	}

	// The child has higher Priority, so move it upward and the parent downward.
	m.data[maxChildIndex], m.data[index] =
		m.data[index], m.data[maxChildIndex]

	// The item moved down may still have lower Priority than its new children,
	// so continue from its new position.
	m.siftDown(maxChildIndex)
}

func (m *MaxPriorityQueue) IsEmpty() bool {
	return len(m.data) == 0
}
