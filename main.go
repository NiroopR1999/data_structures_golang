package main

import (
	"data-structures-golang/queue"
	"fmt"
)

func main() {
	pq := queue.MaxPriorityQueue{}
	pq.Insert("Task A", 3)
	pq.Insert("Task B", 10)
	pq.Insert("Task C", 5)
	pq.Insert("Task D", 1)
	pq.Insert("Task E", 8)

	value, _ := pq.Peek()
	fmt.Println("Peek:", value)

	for !pq.IsEmpty() {
		item, _ := pq.ExtractMax()
		fmt.Println(item.Value, item.Priority)
	}
}
