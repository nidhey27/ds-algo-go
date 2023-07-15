package main

import (
	"fmt"

	"github.com/nidhey27/ds-algo-go/array"
	circularqueue "github.com/nidhey27/ds-algo-go/circular-queue"
	dynamicarray "github.com/nidhey27/ds-algo-go/dynamic-array"
	"github.com/nidhey27/ds-algo-go/queue"
	"github.com/nidhey27/ds-algo-go/stack"
)

func main() {
	var input int
	fmt.Println("-------- Choose DS --------")
	fmt.Println("1. Array")
	fmt.Println("2. Dynamic Array")
	fmt.Println("3. Stack")
	fmt.Println("4. Queue")
	fmt.Println("5. Circular Queue")
	// _, err := fmt.Scanln(&input)
	// if err != nil {
	// 	// Handle the error
	// 	fmt.Println(err)
	// 	return
	// }
	input = 5
	switch input {
	case 1:
		// Array
		a := array.Array{}
		a.Init(10)

		fmt.Println("Array Length:", a.Length())
		// Add
		a.ModifyElement(0, 10)
		a.ModifyElement(1, 20)
		a.ModifyElement(11, 30)

		// Update
		a.ModifyElement(1, 40)

		// Delete
		a.ModifyElement(0, 0)

		// Print
		a.Print()

	case 2:
		// code block 2
		da := dynamicarray.DynamicArray{}
		size := 0
		capacity := 5
		da.Init(size, capacity)

		// Add
		da.Add(10)
		da.Add(20)
		da.Add(30)
		da.Add(40)
		da.Add(50)

		da.Add(60)

		// Get
		fmt.Println("Element at index 10: ", da.Get(10))
		fmt.Println("Element at index 2: ", da.Get(2))

		// Update
		da.Set(4, 99)

		// Print
		fmt.Println("Size: ", da.Size())
		fmt.Println("Capacity: ", da.Capacity())
		da.Print()

	case 3:
		// code block 3
		s := stack.Stack{}
		s.NewStack(5)

		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Push(4)
		s.Push(5)
		s.Push(6)
		//
		fmt.Println("Stack Top: ", s.Peek())
		s.Pop()
		fmt.Println("Stack Top: ", s.Peek())
		s.Pop()
		fmt.Println("Stack Top: ", s.Peek())
		s.Pop()
		fmt.Println("Stack Top: ", s.Peek())
		s.Pop()
		fmt.Println("Stack Top: ", s.Peek())
		s.Pop()
		fmt.Println("Stack Top: ", s.Peek())
		s.Pop()
	case 4:
		q := queue.Queue{}
		q.NewQueue(5)

		fmt.Println("Element: ", q.Peek())
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		q.Enqueue(4)
		q.Enqueue(5)

		for !q.IsEmpty() {
			fmt.Println("Element: ", q.Peek())
			q.Dequeue()
		}
		q.Enqueue(6)
		fmt.Println("Element: ", q.Peek())

		q.Dequeue()

		q.Enqueue(7)
		fmt.Println("Element: ", q.Peek())
		q.Dequeue()

		fmt.Println("Element: ", q.Peek())
	case 5:
		q := circularqueue.CircularQueue{}
		q.NewCircularQueue(5)

		fmt.Println("Element: ", q.Peek())
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		q.Enqueue(4)
		q.Enqueue(5)
		fmt.Println("Element: ", q.Dequeue())
		q.Enqueue(6)
		fmt.Println("Element: ", q.Dequeue())

		q.Enqueue(7)

		for !q.IsEmpty() {
			fmt.Println("Element: ", q.Dequeue())
		}

	default:
		fmt.Println("Invalid option")
	}

}
