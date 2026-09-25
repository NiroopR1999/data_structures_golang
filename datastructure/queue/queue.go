package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	data []int
	size int
}

var ErrQueueIsEmpty = errors.New("queue is empty")

func (q *Queue) Enqueue(value int) {
	// append adds the new element to the end of the slice.
	// The end of the slice represents the rear of the queue.
	q.data = append(q.data, value)

	// Track the number of elements currently in the queue.
	q.size++
}

func (q *Queue) Dequeue() (int, error) {
	// We cannot remove an element if the queue has no elements.
	if q.size == 0 {
		return 0, ErrQueueIsEmpty
	}

	// The first element is always the front of the queue.
	dq := q.data[0]

	// Remove the first element by creating a slice that starts
	// from index 1. This makes the second element the new front.
	q.data = q.data[1:]

	// One element has been removed.
	q.size--

	return dq, nil
}

func (q *Queue) Peek() (int, error) {
	// There is nothing to peek at when the queue is empty.
	if q.size == 0 {
		return 0, ErrQueueIsEmpty
	}

	// Return the front element without removing it.
	return q.data[0], nil
}

func (q *Queue) IsEmpty() bool {
	// The queue is empty when it contains zero elements.
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Dispaly() {
	// Iterate through the queue from front to rear.
	for i := range q.data {
		fmt.Print(q.data[i], " ")
	}
}