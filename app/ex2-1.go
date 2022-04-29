package main

import "fmt"

func main() {
	s := []int{7, 8, 12}
	fmt.Println(s)

	// index
	for i := range s {
		fmt.Println(i)
	}

	// index and value
	for i, v := range s {
		fmt.Println(i, v)
	}
}