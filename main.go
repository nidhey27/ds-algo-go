package main

import (
	"fmt"

	"github.com/nidhey27/ds-algo-go/array"
	dynamicarray "github.com/nidhey27/ds-algo-go/dynamic-array"
)

func main() {
	var input int
	fmt.Println("-------- Choose DS --------")
	fmt.Println("1. Array")
	fmt.Println("2. Dynamic Array")
	_, err := fmt.Scanln(&input)
	if err != nil {
		// Handle the error
		fmt.Println(err)
		return
	}

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

	default:
		fmt.Println("Invalid option")
	}

}
