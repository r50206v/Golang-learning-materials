package main 

import (
	"fmt"
)

func main() {
	// slice
	// composite literal
	x := []int{7, 9, 42}
	for i, _ := range x {
		fmt.Println(i, "-", x[i])
	}

	// slice literal
	y := make([]int, 0, 100)
	y = append(y, 777)

	for i, v := range y {
		fmt.Println(i, "=", v)
	}
}