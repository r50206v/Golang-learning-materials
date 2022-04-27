package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type secret_agent struct {
	person
	ltk bool
}

func (p person) speak() string {
	// return fmt.Sprintln("uptown, func you up, uptown func you up")
	return fmt.Sprintln("hi %s, how are you?", p.fname)
}

func (sa secret_agent) speak() string {
	return fmt.Sprintln("shaken, not stirred")
}

type humanoid interface {
	speak() string
}

func vomit(h humanoid) {
	fmt.Printf("%T\n", h)
	fmt.Println(h)
	switch v := h.(type) {
		case person:
			fmt.Println(v.fname)
		case secret_agent:
			fmt.Println(v.fname)
		default:
			fmt.Println("unknown")
	}
}

func main() {
	p1 := person{"Richard", "Tseng", 27}
	sa1 := secret_agent{person{"Vivian", "ING", 25}, false}
	vomit(p1)
	vomit(sa1)
}