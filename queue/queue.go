package queue

import "fmt"

type Queue struct {
	queue []int
	size  int
	front int
	rear  int
}

func (q *Queue) NewQueue(size int) *Queue {

	q.queue = make([]int, size)
	q.size = size
	q.front = -1
	q.rear = -1

	return q
}

func (q *Queue) IsFull() bool {
	return q.rear == q.size-1
}

func (q *Queue) IsEmpty() bool {
	return q.rear == -1
}

func (q *Queue) Enqueue(elemet int) {
	// First insertion
	if q.front == -1 {
		q.front = 0
	}
	if q.IsFull() {
		fmt.Println("Queue is full...!")
		return
	}
	q.rear = q.rear + 1
	q.queue[q.rear] = elemet
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty...!")
		return -1
	}
	removedElemet := q.queue[q.front]

	for i := 0; i < q.rear; i++ {
		q.queue[i] = q.queue[i+1]
	}
	q.rear = q.rear - 1
	return removedElemet
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty...!")
		return -1
	}
	return q.queue[q.front]
}
