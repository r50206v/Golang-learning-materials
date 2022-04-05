package main

import (
	"fmt"
)

func main() {
	// slice
	// composite literal
	x := []int{7, 9, 42}
	fmt.Println(x)
	// [7, 9, 42]

	// slice literal
	y := make([]int, 0, 100)
	y = append(y, 8)
	y = append(y, 12)
	y = append(y, 43)
	fmt.Println(y, "length=", len(y))
	// [7, 9, 43] length=3
}