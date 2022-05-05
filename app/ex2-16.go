package main

import "fmt"

type gator int

type flamingo bool

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}

type swampCreature interface {
	greeting()
}

func bayou(s swampCreature) {
	s.greeting()
}

func main() {
	var g1 gator
	bayou(g1)

	var f1 flamingo
	bayou(f1)
}
