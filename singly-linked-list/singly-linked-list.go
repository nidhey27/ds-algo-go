package singlylinkedlist

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) InsertAtStart(element string) {
	node := &Node{}
	if l.IsEmpty() {
		node = &Node{data: element, next: nil}
	} else {
		node = &Node{data: element, next: l.head}
	}
	l.head = node
}

func (l *LinkedList) InsertAtEnd(element string) {
	if l.IsEmpty() {
		l.InsertAtStart(element)
		return
	}
	itr := l.head

	for itr.next != nil {
		itr = itr.next
	}

	newNode := &Node{data: element, next: nil}
	itr.next = newNode

}

func (l *LinkedList) InsertAt(index int, element string) {
	if index < 0 || index > l.Length() {
		fmt.Println("Invalid index: ", index)
		return
	}

	if index == 0 {
		l.InsertAtStart(element)
		return
	}

	itr := l.head
	count := 0
	for itr != nil {
		if index-1 == count {
			newNode := &Node{data: element, next: itr.next}
			itr.next = newNode
		}
		count += 1
		itr = itr.next

	}

}

func (l *LinkedList) DeleteAt(index int) {
	if index < 0 || index >= l.Length() {
		fmt.Println("Invalid index: ", index)
		return
	}
	// Delete first node
	if index == 0 {
		newHead := l.head.next
		l.head = newHead
		return
	}

	fmt.Println("Deleting Node...", index)
	// Delete node at given index
	itr := l.head
	for i := 0; i < index - 1; i++{
		itr = itr.next
	}
	itr.next = itr.next.next

}

func (l *LinkedList) Length() int {
	count := 0
	itr := l.head

	for itr != nil {
		count += 1
		itr = itr.next
	}

	return count
}

func (l *LinkedList) TraverseList() {
	fmt.Println("Traversing...")
	itr := l.head
	idx := 0
	for itr != nil {
		fmt.Print(itr.data, ":", idx, " -> ")
		idx += 1
		itr = itr.next

	}
	fmt.Println("END")
}
