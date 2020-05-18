package main

import (
	"cs320/shape"
	"fmt"
)

func main() {
	a := make([]shape.Shape, 5)
	for i := 0; i < 3; i += 1 {
		a[i] = new(shape.Circle)
		a[i].SetLength(5)
	}
	for i := 3; i < 5; i += 1 {
		a[i] = new(shape.Square)
		a[i].SetLength(5)
	}
	//
	for i := 0; i < 5; i += 1 {
		fmt.Println(a[i].GetArea())
	}

	fmt.Println(shape.Pi)
	fmt.Println(shape.not)
}
