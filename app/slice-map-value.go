package main

import (
	"fmt"
)

func main() {
	// slice
	// composite literal; map literal
	x := map[string]int{
		"Todd": 45,
		"Nina": 25,
		"Patrick": 50,
	}

	fmt.Printf("Type: %T\n", x)
	fmt.Println("Type: %T <doesnt work>", x)

	fmt.Println("Nina", x["Nina"])
	fmt.Println("Richard", x["Richard"])
	for k, _ := range x {
		fmt.Println(k, "-", x[k])
	}
}