package main 

import "fmt"

type person struct {
	fname string
	lname string
	// Fname (uppercase) can make this parameter visible outside of pacakge
	// fname (lowercase) will make this parameter invisible outside of pacakge
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good Morning"`)
}

func main() {
	p1 := person{"Richard", "Tseng"}
	p1.speak()
}