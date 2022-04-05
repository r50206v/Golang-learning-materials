package main

import "fmt"

type dog int

func (d dog) barkNum() {
	fmt.Println("I'm a dog", d)
}

type person struct {
	fName string
	lName string
}

func (p person) speak() {
	fmt.Println("Hello, I'm ", p.fName, p.lName)
}

func main() {
	var abc dog
	abc = 17
	fmt.Println("%T\n", abc)
	fmt.Printf("%T\n", abc)
	abc.barkNum()

	def := dog(42)
	fmt.Println("%T\n", def)
	fmt.Printf("%T\n", def)
	def.barkNum()

	// composite literal; struct literal
	p1 := person{"Todd", "McLeod"}
	p2 := person{"Nina", "Simone"}
	fmt.Println(p1)
	fmt.Println(p2)
	p1.speak()
	p2.speak()
}