package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	age	int
}

type secret_agent struct {
	person 
	ltk bool
}

func (p person) speak() {
	fmt.Println(p.fname, `says, "Whatsup!"`)
}

func (s secret_agent) speak() {
	fmt.Println(s.fname, "has a license to kill: ", s.ltk)
}

func main() {
	p1 := person{"Nina", "Simone", 25}
	sa1 := secret_agent{person{"Ian", "Fleming", 42}, false}
	fmt.Println(p1)
	fmt.Println(sa1)
	p1.speak()
	sa1.speak()
	sa1.person.speak()
}