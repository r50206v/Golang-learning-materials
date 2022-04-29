package main 

import "fmt"

type person struct {
	fname string
	lname string
	favFood []string
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
}