package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type secret_agent struct {
	person
	ltk bool
}

func (p person) pSpeak() {
	fmt.Println(p.fname, `says, "Whatsup!"`)
}

func (s secret_agent) saSpeak() {
	fmt.Println(s.fname, "has a license to kill: ", s.ltk)
}

func main() {
	p1 := person{"Richard", "Tseng"}
	sa1 := secret_agent{person{"James", "Bond"}, true}
	fmt.Println(p1.fname)
	p1.pSpeak()

	fmt.Println(sa1.fname)
	sa1.pSpeak()
	sa1.saSpeak()
}