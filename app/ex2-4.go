package main 

import "fmt"

type person struct {
	fname string
	lname string
}

func main() {
	p1 := person{"Allen", "Wang"}	
	fmt.Println(p1.fname)
}