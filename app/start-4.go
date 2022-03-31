package main 

import "fmt"

type person struct {
	fname string
	lname string
	// Fname (uppercase) can make this parameter visible outside of pacakge
	// fname (lowercase) will make this parameter invisible outside of pacakge
}
type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good Morning"`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred!"`)
}

type human interface {
	speak()
}

func saySomething(h human){
	h.speak()
}

func main() {
	p1 := person{"Richard", "Tseng"}
	p1.speak()

	sa1 := secretAgent{
		person{
			"James", 
			"Bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()

	saySomething(sa1)
	saySomething(p1)
}