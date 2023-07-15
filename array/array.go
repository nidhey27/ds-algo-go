package array

import "fmt"

type Array struct {
	items []int
}

// Initilize a new Array with Size X
func (a *Array) Init(size int) *Array {
	a.items = make([]int, size)

	return a
}

// Get length of an Array
func (a *Array) Length() int {
	return len(a.items)
}

// Modify element in an array
func (a *Array) ModifyElement(index, element int) {
	// If the index is out of bounds, print an error message
	if index < 0 || index >= a.Length() {
		fmt.Println("ERROR: Index", index, "out of bound.")
		return
	}

	// If the element value is 0, delete the element at the specified index
	// If the element value is positive, add or update the element at the specified index
	a.items[index] = element
}

// Print elements in an array
func (a *Array) Print() {
	fmt.Println(a.items)
}
