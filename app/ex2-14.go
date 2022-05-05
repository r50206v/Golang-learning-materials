package main

import "fmt"

type gator int

func main() {
	var g1 gator
	g1 = 42

	var x int 

	x = int(g1)
	fmt.Printf("%T\n", x)
	fmt.Println(x)
}