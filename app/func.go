package main

import (
	"fmt"
)

type person struct{
	fName string
	lName string
}

func main() {
	// composite literal; struct literal
	p1 := person{"Richard", "Tseng"}
	p2 := person{"Vivian", "Tseng"}

	fmt.Println(p1)
	fmt.Println(p2)

	m := meaning()
	fmt.Println(m)
}

// func (receiver) identifier(params) (returns) {<code>}
func meaning() int {
	return 42
}