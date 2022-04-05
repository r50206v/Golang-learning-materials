package main

import "fmt"

type person struct {
	fName string
	lName string
}

// func (receiver) identifier(parameters) (returns) { <code> }
func (m person) speak() {
	fmt.Println("Hello, my name is", m.fName)
}

func main() {
	// composite literal, struct literal
	p1 := person{"Todd", "Tseng"}
	p2 := person{"Richard", "MAMA"}
	
	p1.speak()
	p2.speak()
}