package dynamicarray

import "fmt"

type DynamicArray struct {
	items    []int
	capacity int // Total number of elements the array can hold
	size     int // Current number of elements in the array
}

// Initilize a new Dynamic Array with Size X
func (da *DynamicArray) Init(size, capacity int) *DynamicArray {
	da.size = size
	da.capacity = capacity
	da.items = make([]int, da.capacity)
	return da
}

func (da *DynamicArray) Add(element int) {
	if da.size == da.capacity {
		da.resize()
	}
	da.items[da.size] = element
	da.size++
}

func (da *DynamicArray) resize() {
	fmt.Println("Resizing array..")
	fmt.Println("Current Size:", da.size, "Current Capacity: ", da.capacity)
	da.capacity *= 2
	newItems := make([]int, da.capacity)
	copy(newItems, da.items)
	da.items = newItems
	fmt.Println("New Size:", da.size, "New Capacity: ", da.capacity)
	da.Print()
}

// Get element at
func (da *DynamicArray) Get(index int) int {
	if index < 0 || index >= da.size {
		fmt.Println("Index ", index, " out of bounds.")
		return -1
	}

	return da.items[index]

}

// Get element at
func (da *DynamicArray) Set(index, element int) {
	if index < 0 || index >= da.size {
		fmt.Println("Index ", index, " out of bounds.")
		return
	}
	da.items[index] = element
}

// Get Size of an Array
func (da *DynamicArray) Size() int {
	return da.size
}

// Get Capacity of an Array
func (da *DynamicArray) Capacity() int {
	return da.capacity
}

// Print elements in an array
func (da *DynamicArray) Print() {
	fmt.Println(da.items)
}
