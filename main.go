package main

import (
	"fmt"

	"github.com/nidhey27/ds-algo-go/array"
)

func main() {
	var input int
	fmt.Println("-------- Choose DS --------")
	fmt.Println("1. Array")
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

	case 3:
		// code block 3

	default:
		fmt.Println("Invalid option")
	}

}
