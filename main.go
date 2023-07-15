package main

import (
	"fmt"

	"github.com/nidhey27/ds-algo-go/array"
)

func main() {
	a := array.Array{}
	a.Init(10)

	fmt.Println(a.Length())
	a.ModifyElement(0, 10)
	a.ModifyElement(1, 20)
	a.ModifyElement(11, 30)

	a.ModifyElement(1, 40)
	a.ModifyElement(0, 0)
	a.Print()
	fmt.Println(a.Length())
}
