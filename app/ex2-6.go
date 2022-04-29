package main 

import "fmt"

type person struct {
	fname string
	lname string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fname, "is walking")
}

func main() {
	p1 := person{
		"Allen", 
		"Wang", 
		[]string{"a", "b", "c", "d"},
	}
	fmt.Println(p1.favFood)

	for _, v := range p1.favFood {
		fmt.Println(v)
	}
	
	s := p1.walk()
	fmt.Println(s)
}