package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main() {
	p1 := person {
		"Tony",
		"Tseng",
		5,
	}
	fmt.Println(p1)
}