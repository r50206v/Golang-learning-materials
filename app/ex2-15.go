package main

import "fmt"

type gator int

func (g gator) greeting() {
	fmt.Printf("Hello, I am a gator!\n")
}

func main() {
	var g1 gator
	g1 = 5
	g1.greeting()
}