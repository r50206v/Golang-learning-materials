package main

import (
	"fmt"
)

func main() {
	// slice
	// composite literal - map literal
	x := map[string]int{
		"Todd": 45, 
		"Richard": 1,
	}
	for k, v := range x {
		fmt.Println(k, "-", v)
	}
}