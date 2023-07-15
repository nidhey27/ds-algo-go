package circularqueue

import "fmt"

type CircularQueue struct {
	queue []int
	size  int
	front int
	rear  int
}

func (cq *CircularQueue) NewCircularQueue(size int) *CircularQueue {

	cq.queue = make([]int, size)
	cq.size = size
	cq.front = -1
	cq.rear = -1

	return cq
}

func (cq *CircularQueue) IsFull() bool {
	return (cq.rear+1)%cq.size == cq.front
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.rear == -1 && cq.front == -1
}

func (cq *CircularQueue) Enqueue(element int) {
	if cq.IsFull() {
		fmt.Println("Queue is full...!")
		return
	}
	// First insertion
	if cq.front == -1 {
		cq.front = 0
	}

	cq.rear = (cq.rear + 1) % cq.size
	cq.queue[cq.rear] = element
}

func (cq *CircularQueue) Dequeue() int {
	if cq.IsEmpty() {
		fmt.Println("Queue is empty...!")
		return -1
	}

	removedItem := cq.queue[cq.front]
	if cq.front == cq.rear {
		cq.front = -1
		cq.rear = -1
	} else {
		cq.front = (cq.front + 1) % cq.size
	}

	return removedItem
}

func (cq *CircularQueue) Peek() int {
	if cq.IsEmpty() {
		fmt.Println("Queue is empty...!")
		return -1
	}
	return cq.queue[cq.front]
}
