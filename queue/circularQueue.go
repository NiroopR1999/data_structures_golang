package queue

import (
	"errors"
	"fmt"
)

var (
	ErrQueueIsFull = errors.New("queue is full")
)

type CicularQueue struct {
	data  []int
	front int
	rear  int
	cap   int
	size  int
}

// front → WHERE TO REMOVE
// rear  → WHERE TO INSERT
// size  → HOW MANY ELEMENTS EXIST

func NewCircularQueue(capacity int) *CicularQueue {
	return &CicularQueue{
		// We allocate a fixed-size array because a circular queue
		// reuses the same positions instead of growing/shrinking the slice.
		data: make([]int, capacity),

		// cap represents the maximum number of elements the queue can hold.
		cap: capacity,
	}
}

func (c *CicularQueue) Enqueue(value int) error {
	// If size == capacity, there is no free position left in the queue.
	if c.size == c.cap {
		return ErrQueueIsFull
	}

	// rear always points to the position where the next element
	// should be inserted.
	c.data[c.rear] = value

	// Move rear to the next position.
	// The modulo makes rear wrap back to 0 when it reaches capacity.
	// This is what makes the array "circular".
	c.rear = (c.rear + 1) % c.cap

	// Keep track of how many elements are currently in the queue.
	c.size++

	return nil
}

func (c *CicularQueue) Dequeue() (int, error) {
	// If size is 0, there is nothing to remove.
	if c.size == 0 {
		return 0, ErrQueueIsEmpty
	}

	// front always points to the oldest element in the queue.
	dq := c.data[c.front]

	// We don't remove the element from the slice.
	// Instead, we simply move front forward.
	// This is important because a circular queue reuses the array positions.
	c.front = (c.front + 1) % c.cap

	// One element has been removed.
	c.size--

	return dq, nil
}

func (c *CicularQueue) Peek() (int, error) {
	// There is no element to peek at if the queue is empty.
	if c.size == 0 {
		return 0, ErrQueueIsEmpty
	}

	// front points to the first element that would be dequeued.
	// Peek returns it without changing front or size.
	return c.data[c.front], nil
}

func (c *CicularQueue) IsEmpty() bool {
	// size tells us exactly how many elements are currently present.
	return c.size == 0
}

func (c *CicularQueue) Size() int {
	return c.size
}

func (c *CicularQueue) Dispaly() {
	// size tells us how many actual elements we need to print.
	// We cannot simply iterate from front to rear because rear
	// may have wrapped around to the beginning of the array.
	for i := 0; i < c.size; i++ {

		// i is the logical position inside the queue:
		// 0 = first element
		// 1 = second element
		// 2 = third element
		//
		// front + i gives the physical position in the array.
		// % cap makes that position wrap around when it reaches the end.
		//
		// Example:
		// capacity = 5
		// front = 3
		//
		// positions:
		// i = 0 -> (3 + 0) % 5 = 3
		// i = 1 -> (3 + 1) % 5 = 4
		// i = 2 -> (3 + 2) % 5 = 0
		index := (c.front + i) % c.cap

		fmt.Print(c.data[index], " ")
	}
}

