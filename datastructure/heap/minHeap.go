package heap

import (
	"errors"
)

type MinHeap struct {
	data []int
}

func (m *MinHeap) siftUp(index int) {

	// Root has no parent, so there is nothing left to sift up.
	if index == 0 {
		return
	}

	// Find the parent so we can check whether the min-heap property is violated.
	parentIndex := (index - 1) / 2

	// If parent <= child, the min-heap property is already satisfied.
	if m.data[parentIndex] <= m.data[index] {
		return
	}

	// Parent is greater than child, so swap them.
	m.data[index], m.data[parentIndex] =
		m.data[parentIndex], m.data[index]

	// The value moved up may still violate the property with its new parent.
	m.siftUp(parentIndex)
}

func (m *MinHeap) Insert(value int) {

	// Add at the end because the heap must remain a complete binary tree.
	m.data = append(m.data, value)

	// The new value starts at the last position.
	index := len(m.data) - 1

	// Restore the min-heap property by moving the value upward.
	m.siftUp(index)
}

func (m *MinHeap) Size() int {
	return len(m.data)
}

func (m *MinHeap) Peek() (int, error) {

	// The root is at index 0, but accessing it when the heap is empty would panic.
	if m.IsEmpty() {
		return 0, errors.New("heap is empty")
	}

	return m.data[0], nil
}

func (m *MinHeap) IsEmpty() bool {
	return len(m.data) == 0
}

func (m *MinHeap) ExtractMin() (int, error) {

	// The minimum is always at the root, but we cannot access index 0 if the heap is empty.
	if m.IsEmpty() {
		return 0, errors.New("heap is empty")
	}

	// Save the minimum because the root will be replaced.
	min := m.data[0]

	if m.Size() == 1 {
		// After removing the only element, the heap becomes empty.
		m.data = []int{}
		return min, nil
	}

	// Move the last element to the root.
	// We do this because removing index 0 directly would break the complete-tree structure.
	m.data[0], m.data[len(m.data)-1] =
		m.data[len(m.data)-1], m.data[0]

	// Remove the last element, which is now the old minimum.
	// The heap structure remains complete because we remove from the end.
	m.data = m.data[:m.Size()-1]

	// The new root may violate the min-heap property, so move it downward.
	m.siftDown(0)

	return min, nil
}

func (m *MinHeap) siftDown(index int) {

	// Calculate the left child.
	leftChildIndex := (2 * index) + 1

	// If there is no left child, there cannot be a right child either.
	// Therefore this node is a leaf and there is nothing to sift down.
	if leftChildIndex >= len(m.data) {
		return
	}

	// Assume left is smaller; change to right only if right exists and is smaller.
	// A min heap must swap with the smaller child.
	minChildIndex := leftChildIndex

	// We need to check whether a right child exists before accessing it,
	// because a node can have a left child without having a right child.
	rightChildIndex := (2 * index) + 2

	// First check that the right child actually exists.
	// Then compare both children because we must swap with the SMALLER child
	// in a min heap.
	if rightChildIndex < len(m.data) &&
		m.data[rightChildIndex] < m.data[leftChildIndex] {

		minChildIndex = rightChildIndex
	}

	// If parent <= smaller child, the min-heap property is already satisfied.
	// No further movement is necessary.
	if m.data[index] <= m.data[minChildIndex] {
		return
	}

	// Parent is greater than the smaller child, so swap them.
	m.data[index], m.data[minChildIndex] =
		m.data[minChildIndex], m.data[index]

	// The value moved down may still be greater than its new children,
	// so continue sifting from its new position.
	m.siftDown(minChildIndex)
}

func (m *MinHeap) BuildMinHeap(data []int) {

	// Start with the existing array instead of inserting elements one by one.
	m.data = data

	// Leaves already satisfy the heap property because they have no children.
	// Therefore, start from the last non-leaf node.
	//
	// Last non-leaf index = n/2 - 1
	for i := (len(m.data) / 2) - 1; i >= 0; i-- {

		// Process from bottom to top so that when we sift down a node,
		// its children are already valid heaps.
		m.siftDown(i)
	}
}
